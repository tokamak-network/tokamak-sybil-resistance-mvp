// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract NewSybil {
    // ---------- Type Declarations ----------
    struct PairPacked {
        uint64  windowStart; // 0 => not open
        bool    loFunded;    // funded by lower idx
        bool    hiFunded;    // funded by higher idx
        uint128 stakeAmt;    // snapshot stake (wei)
    }

    // ---------- State Variables ----------
    // Reentrancy
    uint256 private constant _NOT_ENTERED = 1;
    uint256 private constant _ENTERED     = 2;
    uint256 private _status = _NOT_ENTERED;

    // Admin / params
    address public owner;
    uint256 public stakeS;    // default stake (wei) for NEW pairs
    uint64  public windowT;   // steal window (seconds)
    uint32  public batchSize; // edges per submit (proof binds to this)

    // Idx registry
    mapping(address => uint32) public accountIdx; // 1-based; 0 = unset
    uint32 public nextIdx = 1;

    // Pair state (nested mapping + two bools)
    mapping(uint32 => mapping(uint32 => PairPacked)) public pairs;

    // Canonical link registry (dedup forever)
    mapping(uint32 => mapping(uint32 => bool)) public isLinkedIdx;

    // Append-only unforged edge log
    mapping(uint32 => uint64) public unforged;
    uint32 public nextEdgeId   = 1;  // next id to write (append)
    uint32 public lastForgedId = 0;  // last processed id
    uint64 public batchId      = 0;  // increments each submit

    // Commitments
    bytes32 public graphRoot;
    bytes32 public scoreRoot;

    // ---------- Events ----------
    event ParamsUpdated(uint256 stakeS, uint64 windowT);
    event BatchSizeUpdated(uint32 batchSize);
    event AccountCreated(address indexed owner, uint32 indexed idx);
    event Vouched(address indexed attester, address indexed subject, uint32 attIdx, uint32 subIdx, uint256 stake);
    event WindowOpened(address indexed lo, address indexed hi, uint32 ilo, uint32 ihi, uint64 start, uint64 end);
    event Stolen(address indexed thief, address indexed victim, uint32 thiefIdx, uint32 victimIdx, uint256 payout);
    event ClosedNoLink(address indexed caller, address indexed counterparty);
    event Linked(address indexed lo, address indexed hi, uint32 ilo, uint32 ihi, uint64 windowStart, uint64 windowEnd);
    event BatchSubmitted(
        uint64 indexed batchId,
        uint32 count,
        uint32 batchSizeUsed,
        bytes32 storageHash,
        bytes32 oldGraphRoot,
        bytes32 oldScoreRoot,
        bytes32 newGraphRoot,
        bytes32 newScoreRoot,
        bytes  edgesPacked
    );

    // ---------- Errors ----------
    error AlreadyHi();
    error AlreadyLo();
    error AlreadyLinked();
    error BadValue();
    error Early();
    error EmptyBatch();
    error MissingIdx();
    error NoWindow();
    error NotBothFunded();
    error NotHiOnly();
    error NotLoOnly();
    error NotOwner();
    error NotParty();
    error PastWindow();
    error Self();
    error StakeZero();
    error VerifyFail();
    error EthXferFail();

    // ---------- Modifiers ----------
    modifier nonReentrant() {
        if (_status != _NOT_ENTERED) revert();
        _status = _ENTERED;
        _;
        _status = _NOT_ENTERED;
    }

    modifier onlyOwner() { 
        if (msg.sender != owner) revert NotOwner();
        _;
    }

    // ---------- Constructor ----------
    constructor(uint256 _s, uint64 _t, uint32 _batchSize) {
        if (_s == 0 || _t == 0 || _batchSize == 0) revert BadValue();
        owner     = msg.sender;
        stakeS    = _s;
        windowT   = _t;
        batchSize = _batchSize;

        emit ParamsUpdated(_s, _t);
        emit BatchSizeUpdated(_batchSize);
    }

    // ---------- External Functions ----------
    function setParams(uint256 _s, uint64 _t) external onlyOwner {
        if (_s == 0 || _t == 0) revert BadValue();
        stakeS  = _s;
        windowT = _t;
        emit ParamsUpdated(_s, _t);
    }

    function setBatchSize(uint32 _n) external onlyOwner {
        if (_n == 0) revert BadValue();
        batchSize = _n;
        emit BatchSizeUpdated(_n);
    }

    function totalAccounts() external view returns (uint32) { 
        return nextIdx - 1; 
    }

    function requiredStake(address x, address y) external view returns (uint256) {
        uint32 xIdx = accountIdx[x];
        uint32 yIdx = accountIdx[y];
        if (xIdx == 0 || yIdx == 0 || xIdx == yIdx) return stakeS;
        (uint32 ilo, uint32 ihi) = xIdx < yIdx ? (xIdx, yIdx) : (yIdx, xIdx);
        PairPacked storage p = pairs[ilo][ihi];
        return p.stakeAmt == 0 ? stakeS : uint256(p.stakeAmt);
    }

    function pendingEdges() external view returns (uint32) {
        uint32 endSnapshot = nextEdgeId - 1;
        if (endSnapshot < lastForgedId) return 0;
        return endSnapshot - lastForgedId;
    }

    /// Lock required stake to vouch `subject`. Auto-creates idx for both.
    function vouch(address subject) external payable nonReentrant {
        // ensure indices
        uint32 attIdx = _ensureIdx(msg.sender);
        uint32 subIdx = _ensureIdx(subject);

        // canonicalize by idx
        (uint32 ilo, uint32 ihi) = attIdx < subIdx ? (attIdx, subIdx) : (subIdx, attIdx);
        if (isLinkedIdx[ilo][ihi]) revert AlreadyLinked();

        PairPacked storage p = pairs[ilo][ihi];

        // stake snapshot & check
        uint256 req = p.stakeAmt == 0 ? stakeS : uint256(p.stakeAmt);
        if (msg.value != req) revert BadValue();
        if (p.stakeAmt == 0) {
            if (req == 0) revert StakeZero();
            if (req > type(uint128).max) revert BadValue();
            p.stakeAmt = uint128(req);
        }

        // set appropriate funded flag based on caller's idx
        bool callerIsLo = (attIdx == ilo);
        if (callerIsLo) {
            if (p.loFunded) revert AlreadyLo();
            p.loFunded = true;
        } else {
            if (p.hiFunded) revert AlreadyHi();
            p.hiFunded = true;
        }

        emit Vouched(msg.sender, subject, attIdx, subIdx, req);

        // open window once both funded
        if (p.windowStart == 0 && p.loFunded && p.hiFunded) {
            p.windowStart = uint64(block.timestamp);
            // derive lo/hi addresses from idxs for UX
            address loAddr = (attIdx == ilo) ? msg.sender : subject;
            address hiAddr = (attIdx == ilo) ? subject   : msg.sender;
            emit WindowOpened(loAddr, hiAddr, ilo, ihi, p.windowStart, p.windowStart + windowT);
        }
    }

    /// Cancel before mutual; refunds your snapshotted stake.
    function cancelVouch(address counterparty) external nonReentrant {
        (uint32 ilo, uint32 ihi, uint32 myIdx, ) = _idxOrder(msg.sender, counterparty);
        PairPacked storage p = pairs[ilo][ihi];

        if (p.windowStart != 0) revert NoWindow(); // already mutual
        uint256 s = uint256(p.stakeAmt);
        if (s == 0) revert StakeZero();

        bool callerIsLo = (myIdx == ilo);
        if (callerIsLo) {
            // lo can cancel only if hi hasn't funded
            if (!(p.loFunded && !p.hiFunded)) revert NotLoOnly();
            p.loFunded = false;
            _send(payable(msg.sender), s);
        } else {
            if (!(p.hiFunded && !p.loFunded)) revert NotHiOnly();
            p.hiFunded = false;
            _send(payable(msg.sender), s);
        }

        // reset snapshot if neither funded
        if (!p.loFunded && !p.hiFunded) {
            p.stakeAmt = 0;
        }

        emit ClosedNoLink(msg.sender, counterparty);
    }

    /// During window, steal and end; no link.
    function steal(address counterparty) external nonReentrant {
        (uint32 ilo, uint32 ihi, uint32 myIdx, uint32 otherIdx) = _idxOrder(msg.sender, counterparty);
        PairPacked storage p = pairs[ilo][ihi];

        if (p.windowStart == 0) revert NoWindow();
        if (block.timestamp > p.windowStart + windowT) revert PastWindow();
        if (!(p.loFunded && p.hiFunded)) revert NotBothFunded();

        uint256 s = uint256(p.stakeAmt);
        if (s == 0) revert StakeZero();

        // clear pair state
        delete pairs[ilo][ihi];

        // payout thief (2*s)
        _send(payable(msg.sender), s * 2);

        emit Stolen(msg.sender, counterparty, myIdx, otherIdx, s * 2);
    }

    /// During window, close without stealing; refund both; no link.
    function closeWithoutSteal(address counterparty) external nonReentrant {
        (uint32 ilo, uint32 ihi, , ) = _idxOrder(msg.sender, counterparty);
        PairPacked storage p = pairs[ilo][ihi];

        if (p.windowStart == 0) revert NoWindow();
        if (block.timestamp > p.windowStart + windowT) revert PastWindow();
        if (!(p.loFunded && p.hiFunded)) revert NotBothFunded();

        uint256 s = uint256(p.stakeAmt);
        if (s == 0) revert StakeZero();

        delete pairs[ilo][ihi];
        _send(payable(msg.sender), s);
        _send(payable(counterparty), s);
        
        emit ClosedNoLink(msg.sender, counterparty);
    }

    /// After window: create link, refund both, append unforged edge (new storage slot).
    function finalize(address a, address b) external nonReentrant {
        (uint32 ilo, uint32 ihi, , ) = _idxOrder(a, b);
        PairPacked storage p = pairs[ilo][ihi];

        if (p.windowStart == 0) revert NoWindow();
        if (block.timestamp <= p.windowStart + windowT) revert Early();
        if (!(p.loFunded && p.hiFunded)) revert NotBothFunded();

        uint256 s = uint256(p.stakeAmt);
        if (s == 0) revert StakeZero();
        if (isLinkedIdx[ilo][ihi]) revert AlreadyLinked();

        // set canonical link
        isLinkedIdx[ilo][ihi] = true;

        uint64 ws = p.windowStart;
        uint64 we = ws + windowT;

        // refund both parties
        delete pairs[ilo][ihi];
        _send(payable(a), s);
        _send(payable(b), s);

        emit Linked(a, b, ilo, ihi, ws, we);

        // Append unforged edge (0->non0 SSTORE ~20k)
        uint32 id = nextEdgeId++;
        unforged[id] = (uint64(ilo) << 32) | uint64(ihi);
    }

    /**
     * Snapshot current end, process up to `batchSize` edges, recompute batch hash from storage,
     * verify proof, emit DA event, delete processed entries, and advance the pointer.
     */
    function submitBatch(bytes32 newGraph, bytes32 newScore, bytes calldata proof) external nonReentrant {
        uint32 start       = lastForgedId + 1;
        uint32 endSnapshot = nextEdgeId - 1; // stable end-of-queue snapshot
        if (endSnapshot < start) revert EmptyBatch();

        uint32 avail = endSnapshot - lastForgedId;
        uint32 n    = avail < batchSize ? avail : batchSize;
        if (n == 0) revert EmptyBatch();

        // Seed hash with (batchId, batchSize) for domain separation
        bytes32 h = keccak256(abi.encodePacked("B", batchId, batchSize));

        // Prepare DA payload
        bytes memory edgesPacked = new bytes(uint256(n) * 8); 
        uint256 off = 0;

        // Recompute hash & assemble DA from storage slice [start .. start+n-1]
        for (uint32 i = 0; i < n; ++i) {
            uint32 id = start + i;
            uint64 w = unforged[id]; // must exist
            uint32 ilo = uint32(w >> 32);
            uint32 ihi = uint32(w);

            // fold hash (canonical ilo<ihi stored at finalize)
            bytes32 ek = keccak256(abi.encodePacked(ilo, ihi));
            h = keccak256(abi.encodePacked(h, ek));

            // write big-endian 8 bytes into edgesPacked
            edgesPacked[off + 0] = bytes1(uint8(ilo >> 24));
            edgesPacked[off + 1] = bytes1(uint8(ilo >> 16));
            edgesPacked[off + 2] = bytes1(uint8(ilo >> 8));
            edgesPacked[off + 3] = bytes1(uint8(ilo));
            edgesPacked[off + 4] = bytes1(uint8(ihi >> 24));
            edgesPacked[off + 5] = bytes1(uint8(ihi >> 16));
            edgesPacked[off + 6] = bytes1(uint8(ihi >> 8));
            edgesPacked[off + 7] = bytes1(uint8(ihi));
            off += 8;
        }

        bytes32 oldG = graphRoot;
        bytes32 oldS = scoreRoot;

        // >>> Replace with real zk verifier bound to:
        // (oldG, oldS, newGraph, newScore, batchId, batchSize, n, h)
        if (!_verifyProof(proof, oldG, oldS, newGraph, newScore, batchId, batchSize, n, h)) {
            revert VerifyFail();
        }

        // Commit new roots
        graphRoot = newGraph;
        scoreRoot = newScore;

        // Emit single DA event
        emit BatchSubmitted(
            batchId,
            n,
            batchSize,
            h,
            oldG,
            oldS,
            newGraph,
            newScore,
            edgesPacked
        );

        // Delete processed entries (non0->0 ~5k each; refunds capped at 20% of tx gas)
        for (uint32 i = 0; i < n; ++i) {
            delete unforged[start + i];
        }

        // Advance pointer & bump batchId
        lastForgedId = lastForgedId + n;
        unchecked { batchId += 1; }
    }

    // ---------- Public Functions ----------
    function hasLink(address a, address b) public view returns (bool) {
        uint32 aIdx = accountIdx[a];
        uint32 bIdx = accountIdx[b];
        if (aIdx == 0 || bIdx == 0 || aIdx == bIdx) return false;
        (uint32 ilo, uint32 ihi) = aIdx < bIdx ? (aIdx, bIdx) : (bIdx, aIdx);
        return isLinkedIdx[ilo][ihi];
    }

    function isFinalizeReady(address x, address y) public view returns (bool) {
        (uint32 ilo, uint32 ihi, , ) = _idxOrder(x, y);
        PairPacked storage p = pairs[ilo][ihi];
        if (p.windowStart == 0) return false;
        if (!(p.loFunded && p.hiFunded)) return false;
        return block.timestamp > (p.windowStart + windowT);
    }

    // ---------- Internal Functions ----------
    function _ensureIdx(address a) internal returns (uint32 idx) {
        idx = accountIdx[a];
        if (idx != 0) return idx;
        require(nextIdx != type(uint32).max, "IDX_EXHAUSTED");
        idx = nextIdx++;
        accountIdx[a] = idx;
        emit AccountCreated(a, idx);
    }

    function _idxOrder(address a, address b) internal view returns (uint32 ilo, uint32 ihi, uint32 aIdx, uint32 bIdx) {
        aIdx = accountIdx[a];
        bIdx = accountIdx[b];
        if (aIdx == 0 || bIdx == 0) revert MissingIdx();
        if (aIdx < bIdx) { ilo = aIdx; ihi = bIdx; }
        else if (aIdx > bIdx) { ilo = bIdx; ihi = aIdx; }
        else revert Self();
    }

    function _verifyProof(
        bytes calldata /*proof*/,
        bytes32 /*oldG*/,
        bytes32 /*oldS*/,
        bytes32 /*newG*/,
        bytes32 /*newS*/,
        uint64  /*batchId_*/,
        uint32  /*batchSize_*/,
        uint32  /*count_*/,
        bytes32 /*hash_*/
    ) internal pure returns (bool) {
        return true;
    }

    function _send(address payable to, uint256 amt) internal {
        (bool ok, ) = to.call{value: amt}("");
        if (!ok) revert EthXferFail();
    }
}
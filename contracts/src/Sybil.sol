// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "./interfaces/ISybil.sol";
import "./interfaces/IVerifier.sol";
import "./types/SybilHelpers.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";

contract Sybil is Initializable, AccessControlUpgradeable, ISybil, SybilHelpers {
    struct Verifier {
        IVerifier verifierInterface;
        uint256 maxTx; // maximum rollup transactions in a batch: L1-tx transactions
        uint256 nLevel; // number of levels of the circuit
    }

    struct ScoreSnapshot {
        uint32 score;
        uint32 batchNum;
    }

    struct AccountInfo {
        uint192 balance; 
        uint24 idx;      
    }

    struct Transaction {
        uint8 identifier;
        uint24 from;
        uint24 to;
        uint128 amount;
    }
    uint256 constant _TXN_TOTALBYTES = 23; // Total bytes per transaction
    uint256 constant _MAX_TXNS = 5; // Max transactions per batch
    uint128 constant _LIMIT_AMOUNT = (1 << 127); // Max loadAmount per call
    uint256 constant _RFIELD = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
    uint256 public _MIN_BALANCE = 1;
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    uint24 public lastIdx;
    uint256 public lastAddedTxn;
    uint256 public lastForgedTxn;
    uint256 public batchSize = 5;
    uint256 public explodeAmount = (1 << 50);
    uint256 public scoringRequiredBalance = (1 << 16);
    uint32 public lastForgedBatch;
    // uint32 public currentFillingBatch;

    mapping(address => AccountInfo) public accountInfo;
    mapping(uint32 => uint256) public accountRootMap;
    mapping(uint32 => uint256) public vouchRootMap;
    mapping(uint32 => uint256) public scoreRootMap;
    mapping(uint32 => uint256) public exitRootMap;
    mapping(uint256 => Transaction) public unprocessedBatchesMap;
    mapping(uint32 => bytes32) public txsDataHashMap;
    // mapping(address => uint256) public balances;
    mapping(address => mapping(address => bool)) public vouches;
    mapping(address => ScoreSnapshot) public scoreSnapshots;

    // Verifier
    Verifier public verifier;

    event TxEvent(
        uint256 indexed lastAddedTxn,
        uint8 indexed identifier,
        uint24 from,
        uint24 to,
        uint256 amount
    );
    event ForgeBatch(uint32 indexed lastForgedBatch, uint256 lastForgedTxn, uint256 batchSize);
    // event WithdrawEvent(uint48 indexed idx, uint32 indexed numExitRoot);
    event ExplodeAmountUpdated(uint256 explodeAmount);
    event ScoringRequiredBalanceUpdated(uint256 newBalance);

    /**
     * @dev Initializes the contract with the specified parameters.
     * This function can only be called once during the deployment of the contract.
     *
     * @param _verifier The address of the verifier contract to be used for rollup verification.
     * @param maxTx The maximum number of transactions allowed in a single batch.
     * @param nLevel The number of levels in the verification circuit.
     * @param _poseidon2Elements The address of the Poseidon hash function elements for 2 elements.
     * @param _poseidon3Elements The address of the Poseidon hash function elements for 3 elements.
     *
     * @notice The deployer of the contract will be granted the `ADMIN_ROLE`.
     */
    function initialize(
        address _verifier,
        uint256 maxTx,
        uint256 nLevel,
        address _poseidon2Elements,
        address _poseidon3Elements,
        address _adminRole
    ) public initializer {

        __AccessControl_init();
        _grantRole(ADMIN_ROLE, _adminRole);

        _initializeVerifiers(_verifier, maxTx, nLevel);

        _initializeHelpers(_poseidon2Elements, _poseidon3Elements);
    }

    function deposit() external payable {
        AccountInfo memory info = accountInfo[msg.sender];
        if (msg.value >= _LIMIT_AMOUNT) {
            revert LimitAmountExceeded();
        }
        if (msg.value < _MIN_BALANCE) {
            revert InsufficientETH();
        }
        if (info.balance == 0) {
            lastIdx++;
            accountInfo[msg.sender].idx = lastIdx;
            _addTx(0, lastIdx, uint24(0), uint128(msg.value));
        } else {
            _addTx(1, info.idx, uint24(0), uint128(msg.value));
        }
        accountInfo[msg.sender].balance = info.balance + uint192(msg.value);
    }

    function withdraw(uint256 amount) external {
        AccountInfo memory info = accountInfo[msg.sender];
        if (amount >= _LIMIT_AMOUNT) {
            revert LimitAmountExceeded();
        }
        if (amount + _MIN_BALANCE > info.balance) {
            revert InsufficientBalance();
        }
        
        unchecked {
            accountInfo[msg.sender].balance = info.balance - uint192(amount);
        }
        (bool success, ) = msg.sender.call{value: amount}("");
        if (!success) {
            revert EthTransferFailed();
        }
        _addTx(2, info.idx, uint24(0), uint128(amount));
    }

    /**
     * @dev Allows a user to vouch for another account.
     *
     * @param toEthAddr The index of the account that is being vouched.
     */
    function vouch(address toEthAddr) external {
        AccountInfo memory senderInfo = accountInfo[msg.sender];
        AccountInfo memory receiverInfo = accountInfo[toEthAddr];
        if (senderInfo.balance == 0) {
            revert SenderHasZeroBalance();
        }
        if (toEthAddr == msg.sender) {
            revert SelfVouch();
        }
        if (receiverInfo.balance == 0) {
            revert ReceiverHasZeroBalance();
        }
        vouches[msg.sender][toEthAddr] = true;
        _addTx(3, senderInfo.idx, receiverInfo.idx, 0);
    }

    /**
     * @dev Allows a user to remove their vouch for another account.
     *
     * @param toEthAddr The index of the account that is being unvouched.
     */

    function unvouch(address toEthAddr) external {
        if (!vouches[msg.sender][toEthAddr]) {
            revert NotVouched(msg.sender, toEthAddr);
        }

        vouches[msg.sender][toEthAddr] = false;
        _addTx(4, accountInfo[msg.sender].idx, accountInfo[toEthAddr].idx, 0);
    }

    /**
     * @dev Allows a user to explode multiple accounts.
     *
     * This function enables a user to explode multiple account by providing an array of address
     *
     * @param toEthAddrs The array of address of the account that is being exploded.
     *
     * Requirement:
     * - All address in `toEthAddrs` must be vouched.
     */
    function explodeMultiple(address[] calldata toEthAddrs) external {
        for (uint256 i = 0; i < toEthAddrs.length; ++i) {
            address toEthAddr = toEthAddrs[i];
            if (!vouches[toEthAddr][msg.sender]) {
                revert NotVouched(msg.sender, toEthAddr);
            }
        }

        AccountInfo memory senderInfo = accountInfo[msg.sender];
        for (uint256 i = 0; i < toEthAddrs.length; ++i) {
            address toEthAddr = toEthAddrs[i];
            AccountInfo memory receiverInfo = accountInfo[toEthAddr];
            uint192 penalty = uint192(Math.min(
                explodeAmount,
                receiverInfo.balance - _MIN_BALANCE
            ));
            unchecked {
                accountInfo[toEthAddr].balance = receiverInfo.balance - penalty;
            }
            accountInfo[msg.sender].balance = senderInfo.balance + penalty;
            vouches[toEthAddr][msg.sender] = false;
            vouches[msg.sender][toEthAddr] = false;
            _addTx(5, senderInfo.idx, receiverInfo.idx, 0);
        }
    }

    /**
     * @dev Processes a batch of transactions and verifies the associated proof.
     *
     * @param newAccountRoot The new account root to be set for the batch.
     * @param newVouchRoot The new vouch root to be set for the batch.
     * @param newScoreRoot The new score root to be set for the batch.
     * @param proofA The first part of the proof used for verification.
     * @param proofB The second part of the proof used for verification.
     * @param proofC The third part of the proof used for verification.
     *
     * @notice The function will revert if the provided proof is invalid.
     *
     * @dev Emits a {ForgeBatch} event indicating the new batch has been forged.
     */
    function forgeBatch(
        uint256 newAccountRoot,
        uint256 newVouchRoot,
        uint256 newScoreRoot,
        uint256[2] calldata proofA,
        uint256[2][2] calldata proofB,
        uint256[2] calldata proofC
    ) external {
        if (lastAddedTxn < lastForgedTxn + batchSize) {
            revert BatchNotFull();
        }
        uint256 input = _constructCircuitInput(
            newAccountRoot,
            newVouchRoot,
            newScoreRoot
        );

        // Verify the proof
        if (
            !verifier.verifierInterface.verifyProof(
                proofA,
                proofB,
                proofC,
                [input]
            )
        ) {
            revert InvalidProof();
        }

        _clearBatchFromQueue();
        lastForgedBatch++;

        accountRootMap[lastForgedBatch] = newAccountRoot;
        vouchRootMap[lastForgedBatch] = newVouchRoot;
        scoreRootMap[lastForgedBatch] = newScoreRoot;
        
        // emit ForgeBatch(lastForgedBatch, lastForgedTxn, batchSize);
        emit ForgeBatch(uint32(1), 5, 5);
    }

    function proveScoreMerkleProof(
        uint32 numScoreRoot,
        uint24 idx,
        uint32 score,
        uint256[] calldata siblings
    ) external {
        uint256[2] memory arrayState;
        arrayState[0] = score;
        arrayState[0] = uint256(uint160(msg.sender));

        uint256 stateHash = _insPoseidonUnit2.poseidon(arrayState);
        uint256 scoreRoot = scoreRootMap[numScoreRoot];

        if (!_smtVerifier(scoreRoot, siblings, idx, stateHash)) {
            revert SmtProofInvalid();
        }

        scoreSnapshots[msg.sender].batchNum = numScoreRoot;
        scoreSnapshots[msg.sender].score = score;
    }

    /**
     * @dev Updates the amount used for the explode operation.
     *
     * @param _explodeAmount The new amount to be set.
     *
     * @notice This function can only be called by an account with the `ADMIN_ROLE`.
     */
    function updateExplodeAmount(
        uint256 _explodeAmount
    ) external onlyRole(ADMIN_ROLE) {
        explodeAmount = _explodeAmount;
        emit ExplodeAmountUpdated(explodeAmount);
    }

    function updateScoringRequiredBalance(
        uint256 _scoringRequiredBalance
    ) external onlyRole(ADMIN_ROLE) {
        scoringRequiredBalance = _scoringRequiredBalance;
        emit ScoringRequiredBalanceUpdated(scoringRequiredBalance);
    }

    /**
     * @dev Retrieves the length of the transaction queue.
     *
     * @return The number of batches in the transaction queue.
     */
    function getQueueLength() external view returns (uint256) {
        return lastAddedTxn - lastForgedTxn;
    }

    /**
     * @dev Adds a transaction to the current filling batch.
     *
     * @param identifier It is used to identify the type of transaction.
     * @param from The Ethereum address who initiated the transaction.
     * @param to The receipient address associated with the transaction.
     * @param amount The amount of Ether.
     *
     * @dev Emits a {TxEvent} event.
     */
    function _addTx(
        uint8 identifier,
        uint24 from,
        uint24 to,
        uint128 amount
    ) internal {
        unprocessedBatchesMap[lastAddedTxn] = Transaction(
            identifier,
            from,
            to,
            amount
        );
        lastAddedTxn++;

        emit TxEvent(lastAddedTxn, identifier, from, to, amount);
    }

    /**
     * @dev Clears the processed batch from the transaction queue.
     */
    function _clearBatchFromQueue() internal {
        for (uint256 i = 0; i < _MAX_TXNS; ++i) {
            delete unprocessedBatchesMap[lastForgedBatch + i];
        }
        lastForgedTxn = lastForgedTxn + batchSize;
    }

    /**
     * @dev Initializes the rollup verifier with the specified parameters.
     *
     * @param _verifier The address of the verifier contract to be used.
     * @param _maxTx The maximum number of transactions allowed in a batch.
     * @param _nLevel The number of levels in the verification circuit.
     *
     * @dev Reverts with `InvalidVerifierAddress` if the provided verifier address is zero.
     */
    function _initializeVerifiers(
        address _verifier,
        uint256 _maxTx,
        uint256 _nLevel
    ) internal {
        if (_verifier == address(0)) {
            revert InvalidVerifierAddress();
        }

        verifier = Verifier({
            verifierInterface: IVerifier(_verifier),
            maxTx: _maxTx,
            nLevel: _nLevel
        });
    }

    /**
     * @dev Constructs the input for the verification circuit.
     *
     * @param newAccountRoot The new account root to be included in the input.
     * @param newVouchRoot The new vouch root to be included in the input.
     * @param newScoreRoot The new score root to be included in the input.
     *
     * @return The hashed input for the verification circuit, reduced modulo `_RFIELD`.
     */
    function _constructCircuitInput(
        uint256 newAccountRoot,
        uint256 newVouchRoot,
        uint256 newScoreRoot
    ) internal view returns (uint256) {
        uint256 oldAccountRoot = accountRootMap[lastForgedBatch];
        uint256 oldVouchRoot = vouchRootMap[lastForgedBatch];
        uint256 oldScoreRoot = scoreRootMap[lastForgedBatch];
        Transaction[] memory transactions = new Transaction[](batchSize);
        for (uint256 i = 0; i < _MAX_TXNS; ++i) {
            transactions[i] = unprocessedBatchesMap[lastForgedTxn + i];
        }

        bytes memory txnData = abi.encode(transactions);

        bytes memory inputBytes = abi.encodePacked(
            oldAccountRoot,
            oldVouchRoot,
            oldScoreRoot,
            newAccountRoot,
            newVouchRoot,
            newScoreRoot,
            txnData
        );
        return uint256(sha256(inputBytes)) % _RFIELD;
    }
}

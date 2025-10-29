// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

interface INewSybil {
    // Events
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

    // Errors
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

    // External Functions
    function setParams(uint256 _s, uint64 _t) external;
    function setBatchSize(uint32 _n) external;
    function totalAccounts() external view returns (uint32);
    function requiredStake(address x, address y) external view returns (uint256);
    function pendingEdges() external view returns (uint32);
    function vouch(address subject) external payable;
    function cancelVouch(address counterparty) external;
    function steal(address counterparty) external;
    function closeWithoutSteal(address counterparty) external;
    function finalize(address a, address b) external;
    function submitBatch(bytes32 newGraph, bytes32 newScore, bytes calldata proof) external;
    function hasLink(address a, address b) external view returns (bool);
    function isFinalizeReady(address x, address y) external view returns (bool);
}
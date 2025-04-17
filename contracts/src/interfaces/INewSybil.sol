// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

interface INewSybil {
    error InvalidVerifierAddress();
    error LoadAmountExceedsLimit();
    error LoadAmountDoesNotMatch();
    error AmountExceedsLimit();
    error WithdrawAlreadyDone();
    error SmtProofInvalid();
    error EthTransferFailed();
    error InvalidProof();
    error InvalidFromIdx();
    error InvalidToIdx();
    error LimitAmountExceeded();
    error InsufficientETH();
    error InsufficientBalance();
    error SenderHasZeroBalance();
    error ReceiverHasZeroBalance();
    error NotVouched(address from, address to);

    // Initialization function
    function initialize(
        address verifier,
        uint256 maxTx,
        uint256 nLevel,
        address _poseidon2Elements,
        address _poseidon3Elements,
        address _adminRole
    ) external;

    // Batch forging function
    function forgeBatch(
        uint256 newStRoot,
        uint256 newVouchRoot,
        uint256 newScoreRoot,
        uint256[2] calldata proofA,
        uint256[2][2] calldata proofB,
        uint256[2] calldata proofC
    ) external;

    // Getter functions
    function getStateRoot(uint32 batchNum) external view returns (uint256);
    function getLastForgedBatch() external view returns (uint32);

    // L1 Transaction Queue functions
    function getL1TransactionQueue(
        uint32 queueIndex
    ) external view returns (bytes memory);
    function getQueueLength() external view returns (uint32);

    // Deposit Function
    function deposit() external payable;

    // Withdraw Function
    function withdraw(uint256 amount) external;

    // Explode function
    function explodeMultiple(address[] calldata toEthAddrs) external;

    // Vouch function
    function vouch(address toEthAddr) external;

    // Unvouch function
    function unvouch(address toEthAddr) external;

    // Updates the score
    function proveScoreMerkleProof(
        uint32 numScoreRoot, 
		uint24 idx,
		uint32 score, 
		uint256[] memory siblings
    ) external;

    // setter functions
    function updateExplodeAmount(uint256 _explodeAmount) external;
    function updateMinBalance(uint256 _minBalance) external;

}

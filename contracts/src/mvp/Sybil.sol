// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "../interfaces/IMVPSybil.sol";
import "../interfaces/IVerifierRollup.sol";
import "../types/mvp/SybilHelpers.sol";

contract Sybil is Initializable, AccessControlUpgradeable, IMVPSybil, MVPSybilHelpers {

    struct VerifierRollup {
        VerifierRollupInterface verifierInterface;
        uint256 maxTx; // maximum rollup transactions in a batch: L1-tx transactions
        uint256 nLevel; // number of levels of the circuit
    }

    uint48 constant _RESERVED_IDX = 255;
    uint48 constant _EXIT_IDX = 1;
    uint48 constant _EXPLODE_IDX = 2;
    uint256 constant _TXN_TOTALBYTES = 128; // Total bytes per transaction
    uint256 constant _MAX_TXNS = 1000; // Max transactions per batch
    uint256 constant _LIMIT_LOADAMOUNT = (1 << 128); // Max loadAmount per call
    uint256 constant _RFIELD = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    uint48 public lastIdx;
    uint32 public lastForgedBatch;
    uint32 public currentFillingBatch;
    uint256 public minBalance;
    uint256 public explodeAmount;

    mapping(uint32 => uint256) public accountRootMap;
    mapping(uint32 => uint256) public vouchRootMap;
    mapping(uint32 => uint256) public scoreRootMap;
    mapping(uint32 => uint256) public exitRootMap;
    mapping(uint32 => bytes) public unprocessedBatchesMap;

    // Mapping of exit nullifiers, only allowing each withdrawal to be made once
    mapping(uint32 => mapping(uint48 => bool)) public exitNullifierMap;

    // Verifier
    VerifierRollup public rollupVerifier;

    event L1UserTxEvent(
        uint32 indexed queueIndex,
        uint8 indexed position,
        bytes l1UserTx
    );

    event ForgeBatch(uint32 indexed batchNum, uint16 l1UserTxsLen);
    event WithdrawEvent(
        uint48 indexed idx,
        uint32 indexed numExitRoot
    );
    event ExplodeAmountUpdated(uint256 explodeAmount);
    event MinBalanceUpdated(uint256 minBalance);

    /**
     * @dev Initializes the contract with the specified parameters.
     * This function can only be called once during the deployment of the contract.
     *
     * @param verifier The address of the verifier contract to be used for rollup verification.
     * @param maxTx The maximum number of transactions allowed in a single batch.
     * @param nLevel The number of levels in the verification circuit.
     * @param _poseidon2Elements The address of the Poseidon hash function elements for 2 elements.        
     * @param _poseidon3Elements The address of the Poseidon hash function elements for 3 elements.                   
     * @param _poseidon4Elements The address of the Poseidon hash function elements for 4 elements.
     *
     * @notice The deployer of the contract will be granted the `ADMIN_ROLE`.
    */
    function initialize(
        address verifier,
        uint256 maxTx,
        uint256 nLevel,
        address _poseidon2Elements,
        address _poseidon3Elements,
        address _poseidon4Elements,
        address _adminRole
    ) public initializer {
        lastIdx = _RESERVED_IDX;
        currentFillingBatch = 1;

        __AccessControl_init();
        _grantRole(ADMIN_ROLE, _adminRole);

        _initializeVerifiers(
            verifier,
            maxTx,
            nLevel
        );

        _initializeHelpers(
            _poseidon2Elements,
            _poseidon3Elements,
            _poseidon4Elements
        );
    }

    /**
     * @dev Allows a user to create an account deposit.
     * 
     * This function accepts a specified amount of Ether, which is converted from a fixed-point 
     * representation to a standard uint256 value.
     * 
     * @param loadAmountF The amount of Ether to deposit, represented as a fixed-point number
     * 
     * Requirements:
     * 
     * - The `loadAmountF` must be less than the maximum load amount defined by `_LIMIT_LOADAMOUNT`.
     * - The amount of Ether sent with the transaction must match the converted `loadAmount`.
    */
    function createAccountDeposit(uint40 loadAmountF) external payable override {
        uint256 loadAmount = _float2Fix(loadAmountF);
        if(loadAmount >= _LIMIT_LOADAMOUNT) {
            revert LoadAmountExceedsLimit();
        }

        if(loadAmount != msg.value) {
            revert LoadAmountDoesNotMatch();
        }
        
        _addTx(msg.sender, 0, loadAmountF, 0, 0);
    }

    /**
     * @dev Allows a user to deposit Ether into their account.
     *
     * This function accepts a specified amount of Ether, which is converted from a fixed-point 
     * representation to a standard uint256 value.
     *
     * @param fromIdx The index of the account to which the deposit is being made.
     * @param loadAmountF The amount of Ether to deposit, represented as a fixed-point number 
     * 
     * Requirements:
     * 
     * - The `loadAmountF` must be less than the maximum load amount defined by `_LIMIT_LOADAMOUNT`.
     * - The amount of Ether sent with the transaction must match the converted `loadAmount`.
    */
    function deposit(uint48 fromIdx, uint40 loadAmountF) external payable override {
        uint256 loadAmount = _float2Fix(loadAmountF);

        if(loadAmount >= _LIMIT_LOADAMOUNT) {
            revert LoadAmountExceedsLimit();
        }

        if(loadAmount != msg.value) {
            revert LoadAmountDoesNotMatch();
        }

        _validateFromIdx(fromIdx);

        _addTx(msg.sender, fromIdx, loadAmountF, 0, 0);
    }

    /**
     * @dev Allows a user to vouch for another account.
     *
     * @param fromIdx The index of the account that is vouching.
     * @param toIdx The index of the account being vouched for.
     * 
     * Requirement:
     * - Both `fromIdx` and `toIdx` must be valid indices.
    */
    function vouch(uint48 fromIdx, uint48 toIdx) external {

        _validateFromIdx(fromIdx);
        _validateToIdx(toIdx);

        _addTx(msg.sender, fromIdx, 0, 1, toIdx);
    }
    /**
     * @dev Allows a user to remove their vouch for another account.
     *
     * @param fromIdx The index of the account that is unvouching.
     * @param toIdx The index of the account being unvouched for.
     * 
     * Requirement:
     * - Both `fromIdx` and `toIdx` must be valid indices.
    */
    function unvouch(uint48 fromIdx, uint48 toIdx) external {

        _validateFromIdx(fromIdx);
        _validateToIdx(toIdx);

        _addTx(msg.sender, fromIdx, 0, 0, toIdx);
    }

    /**
     * @dev Allows a user to exit their account by withdrawing a specified amount of Ether.
     *
     * @param fromIdx The index of the account that is exiting.
     * @param amountF The amount of Ether to withdraw, represented as a fixed-point number 
     *
     * Requirement:
     * - The `amountF` must be less than the maximum load amount defined by `_LIMIT_LOADAMOUNT`.
    */
    function exit(uint48 fromIdx, uint40 amountF) external override {
        uint256 amount = _float2Fix(amountF);

        if(amount >= _LIMIT_LOADAMOUNT) {
            revert AmountExceedsLimit();
        }

        _validateFromIdx(fromIdx);

        _addTx(msg.sender, fromIdx, 0, amountF, _EXIT_IDX);
    }

    /**
     * @dev Allows a user to explode multiple accounts from a specified index.
     *
     * This function enables a user to explode (or transfer) their account to multiple 
     * specified indices. The transaction is added to the queue for processing.
     *
     * @param fromIdx The index of the account that is explodeMultiple.
     * @param toIdxs An array of indices representing the accounts being exploded for.
     * 
     * Requirement:
     * - All indices in `toIdxs` must be valid.
    */
    function explodeMultiple(uint48 fromIdx, uint48[] memory toIdxs) external override {

        _validateFromIdx(fromIdx);

        uint256 length = toIdxs.length;
        for (uint256 i = 0; i < length; ++i) {
            _validateToIdx(toIdxs[i]);
        }

        for (uint256 i = 0; i < length; ++i) {
            _addTx(msg.sender, fromIdx, 0, 2, toIdxs[i]);
        }
    }

    /**
     * @dev Processes a batch of transactions and verifies the associated proof.
     *
     * @param newLastIdx The new last index to be set for the batch.
     * @param newAccountRoot The new account root to be set for the batch.
     * @param newVouchRoot The new vouch root to be set for the batch.
     * @param newScoreRoot The new score root to be set for the batch.
     * @param newExitRoot The new exit root to be set for the batch.
     * @param proofA The first part of the proof used for verification.
     * @param proofB The second part of the proof used for verification.
     * @param proofC The third part of the proof used for verification.
     *
     * @notice The function will revert if the provided proof is invalid.
     *
     * @dev Emits a {ForgeBatch} event indicating the new batch has been forged.
    */
    function forgeBatch(
        uint48 newLastIdx,
        uint256 newAccountRoot,
        uint256 newVouchRoot,
        uint256 newScoreRoot,
        uint256 newExitRoot,
        uint256[2] calldata proofA,
        uint256[2][2] calldata proofB,
        uint256[2] calldata proofC
    ) external override {
        uint256 input = _constructCircuitInput(
          newLastIdx,
          newAccountRoot,
          newVouchRoot,
          newScoreRoot,
          newExitRoot
      );

        // Verify the proof using the specific rollup verifier
        if (
            !rollupVerifier.verifierInterface.verifyProof(
                proofA,
                proofB,
                proofC,
                [input]
            )
        ) {
            revert InvalidProof();
        }

        lastForgedBatch++;
        lastIdx = newLastIdx;
        accountRootMap[lastForgedBatch] = newAccountRoot;
        vouchRootMap[lastForgedBatch] = newVouchRoot;
        scoreRootMap[lastForgedBatch] = newScoreRoot;
        exitRootMap[lastForgedBatch] = newExitRoot;

        uint16 l1UserTxsLen = _clearBatchFromQueue();

        emit ForgeBatch(lastForgedBatch, l1UserTxsLen);
    }

    /**
     * @dev Allows a user to withdraw funds based on a Merkle proof.
     *
     * @param amount The amount of funds to withdraw
     * @param numExitRoot The index of the exit root to be used for verification.
     * @param siblings An array of sibling hashes used in the Merkle proof.
     * @param idx The index of the account from which the funds are being withdrawn.
     *
     * @notice The function will revert if the withdrawal has already been processed or if 
     *         the proof is invalid.
     *
     * @dev Emits a {WithdrawEvent} indicating the withdrawal has been processed.
    */
    function withdrawMerkleProof(
        uint192 amount,
        uint32 numExitRoot,
        uint256[] calldata siblings,
        uint48 idx
    ) external {
        uint256[4] memory arrayState = _buildTreeState(amount, msg.sender);
        uint256 stateHash = _hash4Elements(arrayState);

        uint256 exitRoot = exitRootMap[numExitRoot];

        if (exitNullifierMap[numExitRoot][idx]) {
            revert WithdrawAlreadyDone();
        }

        if (!_smtVerifier(exitRoot, siblings, idx, stateHash)) {
            revert SmtProofInvalid();
        }

        exitNullifierMap[numExitRoot][idx] = true;

        _withdrawFunds(amount);
        emit WithdrawEvent(idx, numExitRoot);
    }

    /**
     * @dev Updates the amount used for the explode operation.
     *
     * @param _explodeAmount The new amount to be set.
     *
     * @notice This function can only be called by an account with the `ADMIN_ROLE`.
    */
    function updateExplodeAmount(uint256 _explodeAmount) external override onlyRole(ADMIN_ROLE) {
        explodeAmount = _explodeAmount;
        emit ExplodeAmountUpdated(explodeAmount);
    }

    /**
     * @dev Updates the minimum balance required for accounts.
     *
     * @param _minBalance The new minimum balance to be set.
     *
     * @notice This function can only be called by an account with the `ADMIN_ROLE`.
    */
    function updateMinBalance(uint256 _minBalance) external override onlyRole(ADMIN_ROLE){
        minBalance = _minBalance;
        emit MinBalanceUpdated(explodeAmount);
    }

    /**
     * @dev Retrieves the state root for a specific batch number.
     *
     * @param batchNum The batch number associated with the state root.
     * @return The account root of the specified batch number.
    */
    function getStateRoot(uint32 batchNum) external view override returns (uint256) {
        return accountRootMap[batchNum];
    }

    /**
     * @dev Retrieves the last forged batch number.
     *
     * @return The last forged batch number.
    */
    function getLastForgedBatch() external view override returns (uint32) {
        return lastForgedBatch;
    }

    /**
     * @dev Retrieves the transaction queue for a specific index.
     * 
     * @param queueIndex The index of the transaction queue.
     * @return A bytes array containing the unprocessed transactions for the specified index.
    */
    function getL1TransactionQueue(uint32 queueIndex) external view override returns (bytes memory) {
        return unprocessedBatchesMap[queueIndex];
    }

    /**
     * @dev Retrieves the length of the transaction queue.
     *
     * @return The number of batches in the transaction queue.
    */
    function getQueueLength() external view override returns (uint32) {
        return currentFillingBatch - lastForgedBatch;
    }

    /**
     * @dev Adds a transaction to the current filling batch.
     *
     * @param ethAddress The Ethereum address associated with the transaction.
     * @param fromIdx The index of the account from which the transaction originates.
     * @param loadAmountF The amount of Ether to load, represented as a fixed-point number 
     * @param amountF The amount of Ether to transfer, represented as a fixed-point number 
     * @param toIdx The index of the account to which the transaction is directed.
     *
     * @dev Emits a {L1User TxEvent} event.
    */
    function _addTx(
        address ethAddress,
        uint48 fromIdx,
        uint40 loadAmountF,
        uint40 amountF,
        uint48 toIdx
    ) public override {
        bytes memory l1Tx = abi.encodePacked(
            ethAddress,
            fromIdx,
            loadAmountF,
            amountF,
            toIdx
        );

        uint256 currentPosition = unprocessedBatchesMap[currentFillingBatch].length /
            _TXN_TOTALBYTES;

        unprocessedBatchesMap[currentFillingBatch] = bytes.concat(
            unprocessedBatchesMap[currentFillingBatch],
            l1Tx
        );

        emit L1UserTxEvent(currentFillingBatch, uint8(currentPosition), l1Tx);

        if (currentPosition + 1 >= _MAX_TXNS) {
            currentFillingBatch++;
        }
    }

    /**
     * @dev Clears the processed batch from the transaction queue.
     * 
     * @return The number of transactions that were in the cleared batch.
    */
    function _clearBatchFromQueue() internal returns (uint16) {
        uint16 l1UserTxsLen = uint16(
            unprocessedBatchesMap[lastForgedBatch].length / _TXN_TOTALBYTES
        );
        delete unprocessedBatchesMap[lastForgedBatch];
        if (lastForgedBatch + 1 == currentFillingBatch) {
            currentFillingBatch++;
        }
        return l1UserTxsLen;
    }

    /**
     * @dev Transfers Ether to the specified address.
     *
     * @param value The amount of Ether to transfer, specified in wei.
     *
     * @dev Reverts with `EthTransferFailed` if the transfer is unsuccessful.
    */
    function _safeTransfer(uint256 value) internal {
        (bool success, ) = msg.sender.call{value: value}(new bytes(0));
        if (!success) {
            revert EthTransferFailed();
        }
    }

    /**
     * @dev Withdraws a specified amount of funds from the contract.
     *
     * @param amount The amount of Ether to withdraw, specified in wei.
    */
    function _withdrawFunds(uint192 amount) internal {
        _safeTransfer(amount);
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

        rollupVerifier = VerifierRollup({
            verifierInterface: VerifierRollupInterface(_verifier),
            maxTx: _maxTx,
            nLevel: _nLevel
        });
    }

    /**
     * @dev Constructs the input for the verification circuit.
     *
     * @param newLastIdx The new last index to be included in the input.
     * @param newAccountRoot The new account root to be included in the input.
     * @param newVouchRoot The new vouch root to be included in the input.
     * @param newScoreRoot The new score root to be included in the input.
     * @param newExitRoot The new exit root to be included in the input.
     * 
     * @return The hashed input for the verification circuit, reduced modulo `_RFIELD`.
    */
    function _constructCircuitInput(
        uint48 newLastIdx,
        uint256 newAccountRoot,
        uint256 newVouchRoot,
        uint256 newScoreRoot,
        uint256 newExitRoot
    ) internal view returns (uint256) {
        uint256 oldAccountRoot = accountRootMap[lastForgedBatch];
        uint256 oldVouchRoot = vouchRootMap[lastForgedBatch];
        uint256 oldScoreRoot = scoreRootMap[lastForgedBatch];
        uint48 oldLastIdx = lastIdx;
        bytes memory txnData = unprocessedBatchesMap[lastForgedBatch+1];

        bytes memory inputBytes = abi.encodePacked(
            oldLastIdx,
            oldAccountRoot,
            oldVouchRoot,
            oldScoreRoot,
            newLastIdx,
            newAccountRoot,
            newVouchRoot,
            newScoreRoot,
            newExitRoot,
            txnData
        );
        return uint256(sha256(inputBytes)) % _RFIELD;
    }    

    /**
     * @dev Builds the state for the Merkle tree.
     *
     * @param amount The amount to be included in the state.
     * @param user The address of the user associated with the state.
     * 
     * @return A uint256 array representing the state for the Merkle tree.
    */
    function _buildTreeState(uint192 amount, address user) internal pure returns (uint256[4] memory) {
        uint256[4] memory state;
        state[0] = amount;
        state[1] = uint256(uint160(user)); // Convert address to uint256
        state[2] = 0;
        state[3] = 0;
        return state;
    }

    /**
     * @dev Validates the `fromIdx` parameter to ensure it is within acceptable bounds.
     *
     * @param fromIdx The index to validate.
     *
     * @dev Reverts with `InvalidFromIdx` if validation fails.
    */
    function _validateFromIdx(uint48 fromIdx) internal view {
        if ((fromIdx <= _RESERVED_IDX) || (fromIdx > lastIdx)) {
            revert InvalidFromIdx();
        }
    }

    /**
     * @dev Validates the `toIdx` parameter to ensure it is within acceptable bounds.
     *
     * @param toIdx The index to validate.
     *
     * @dev Reverts with `InvalidToIdx` if validation fails.
    
    
    function _validateToIdx(uint48 toIdx) internal view {
        if ((toIdx <= _RESERVED_IDX) || (toIdx > lastIdx)) {
            revert InvalidToIdx();
        }
    }

    /**
     * @dev Converts a fixed-point representation to a standard uint256 value.
     *
     * @param floatVal The fixed-point number to convert
     * 
     * @return The converted value as a uint256.
    */
    function _float2Fix(uint40 floatVal) internal pure returns(uint256) {
        uint256 m = floatVal & 0x7FFFFFFFF;
        uint256 e = floatVal >> 35;

        uint256 exp = 10**e;
        uint256 fix = m * exp;

        return fix;
    }

}

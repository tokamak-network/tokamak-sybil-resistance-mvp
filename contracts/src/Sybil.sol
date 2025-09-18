// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "./interfaces/ISybil.sol";
import "./interfaces/IVerifier.sol";
import "./types/SybilHelpers.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";

/**
 * @title Sybil Resistance Contract
 * @author Tokamak Network
 * @notice This contract implements a sybil resistance mechanism using Zero-Knowledge proofs and vouching systems
 * @dev This contract uses upgradeable patterns and access control for administrative functions
 */
contract Sybil is
    Initializable,
    AccessControlUpgradeable,
    ISybil,
    SybilHelpers
{

    /// @notice Structure to store user's score snapshot at a specific batch
    struct ScoreSnapshot {
        uint32 score; /// @dev User's score value
        uint32 batchNum; /// @dev Batch number when the score was recorded
    }

    /// @notice Structure to store account information
    struct AccountInfo {
        uint192 balance; /// @dev Account balance in wei
        uint24 idx; /// @dev Unique account index assigned during first deposit
    }

    /// @notice Structure representing a transaction in the rollup
    struct Transaction {
        uint8 identifier; /// @dev Transaction type identifier (0: createAccount, 1: deposit, 2: withdraw, 3: vouch, 4: unvouch, 5: explode)
        uint24 from; /// @dev Sender account index
        uint24 to; /// @dev Receiver account index
        uint128 amount; /// @dev Transaction amount
    }

    /// @dev Maximum transactions allowed per batch
    uint256 constant _MAX_TXNS = 5;
    /// @dev Maximum amount that can be deposited or withdrawn in a single transaction
    uint128 constant _LIMIT_AMOUNT = (1 << 127);
    /// @dev BN254 field modulus for circuit calculations
    uint256 constant _RFIELD =
        21888242871839275222246405745257275088548364400416034343698204186575808495617;
    
    /// @notice Minimum balance that must remain in an account after deposit
    uint256 public _MIN_BALANCE = 1;
    /// @notice Admin role identifier for access control
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    /// @notice Last assigned account index
    uint24 public lastIdx;
    /// @notice Index of the last added transaction
    uint256 public lastAddedTxn;
    /// @notice Index of the last forged transaction
    uint256 public lastForgedTxn;
    /// @notice Number of transactions required to form a complete batch
    uint256 public batchSize = 5;
    /// @notice Amount deducted from exploded accounts as penalty
    uint256 public explodeAmount = (1 << 50);
    /// @notice Minimum balance required to participate in scoring
    uint256 public scoringRequiredBalance = (1 << 16);
    /// @notice Last forged batch number
    uint32 public lastForgedBatch;

    /// @notice Mapping from user address to their account information
    mapping(address => AccountInfo) public accountInfo;
    /// @notice Mapping from batch number to account merkle root
    mapping(uint32 => uint256) public accountRootMap;
    /// @notice Mapping from batch number to vouch merkle root
    mapping(uint32 => uint256) public vouchRootMap;
    /// @notice Mapping from batch number to score merkle root
    mapping(uint32 => uint256) public scoreRootMap;
    /// @notice Mapping from transaction index to transaction data
    mapping(uint256 => Transaction) public unprocessedBatchesMap;
    /// @notice Mapping to track vouch relationships between accounts
    mapping(address => mapping(address => bool)) public vouches;
    /// @notice Mapping from user address to their score snapshot
    mapping(address => ScoreSnapshot) public scoreSnapshots;

    /// @notice Verifier contract address
    IVerifier public verifier;

    /// @notice Emitted when a new transaction is added to the queue
    /// @param lastAddedTxn Index of the transaction that was added
    /// @param identifier Type of transaction (0-5)
    /// @param from Sender account index
    /// @param to Receiver account index
    /// @param amount Transaction amount
    event TxEvent(
        uint256 indexed lastAddedTxn,
        uint8 indexed identifier,
        uint24 from,
        uint24 to,
        uint256 amount
    );

    /// @notice Emitted when a batch is successfully forged
    /// @param lastForgedBatch Batch Number that was forged
    /// @param lastForgedTxn Index of the last transaction in the forged batch
    /// @param batchSize Number of transactions in the batch
    /// @param txnData Encoded transaction data
    event ForgeBatch(
        uint32 indexed lastForgedBatch,
        uint256 lastForgedTxn,
        uint256 batchSize,
        bytes txnData
    );

    /// @notice Emitted when explode amount is updated
    /// @param explodeAmount New explode amount value
    event ExplodeAmountUpdated(uint256 explodeAmount);

    /// @notice Emitted when scoring required balance is updated
    /// @param newBalance New required balance for scoring
    event ScoringRequiredBalanceUpdated(uint256 newBalance);

    /**
     * @notice Initializes the contract with the specified parameters
     * @dev This function can only be called once during the deployment of the contract
     * @param _verifier The address of the verifier contract to be used for rollup verification
     * @param _poseidon1Elements The address of the Poseidon hash function contract for 1 element
     * @param _poseidon2Elements The address of the Poseidon hash function contract for 2 elements
     * @param _poseidon3Elements The address of the Poseidon hash function contract for 3 elements
     * @param _adminRole The address that will be granted admin privileges
     */
    function initialize(
        address _verifier,
        address _poseidon1Elements,
        address _poseidon2Elements,
        address _poseidon3Elements,
        address _adminRole
    ) public initializer {
        __AccessControl_init();
        _grantRole(ADMIN_ROLE, _adminRole);

        if (_verifier == address(0)) {
            revert InvalidVerifierAddress();
        }
        verifier = IVerifier(_verifier);

        _initializeHelpers(_poseidon1Elements, _poseidon2Elements, _poseidon3Elements);
    }

    /**
     * @notice Allows users to deposit ETH into their account
     * @dev Creates a new account if this is the user's first deposit, otherwise adds to existing balance
     * @dev Reverts if deposit amount exceeds limit or is below minimum balance
     * @dev Emits a TxEvent with identifier 0 (new account) or 1 (existing account)
     */
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

    /**
     * @notice Allows users to withdraw ETH from their account
     * @dev Ensures minimum balance is maintained after withdrawal
     * @param amount The amount of ETH to withdraw in wei
     * @dev Reverts if withdrawal amount exceeds limit or would leave insufficient balance
     * @dev Emits a TxEvent with identifier 2 (withdraw)
     */
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
     * @notice Allows a user to vouch for another account
     * @dev Creates a trust relationship between two accounts
     * @param toEthAddr The address of the account being vouched for
     * @dev Reverts if already vouched, sender has zero balance, self-vouch attempt, or receiver has zero balance
     * @dev Emits a TxEvent with identifier 3 (vouch)
     */
    function vouch(address toEthAddr) external {
        AccountInfo memory senderInfo = accountInfo[msg.sender];
        AccountInfo memory receiverInfo = accountInfo[toEthAddr];
        
        if (vouches[msg.sender][toEthAddr]) {
            revert AlreadyVouched(msg.sender, toEthAddr);
        }
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
     * @notice Allows a user to remove their vouch for another account
     * @dev Removes the trust relationship between two accounts
     * @param toEthAddr The address of the account being unvouched
     * @dev Reverts if no existing vouch relationship
     * @dev Emits a TxEvent with identifier 4 (unvouch)
     */

    function unvouch(address toEthAddr) external {
        if (!vouches[msg.sender][toEthAddr]) {
            revert NotVouched(msg.sender, toEthAddr);
        }

        vouches[msg.sender][toEthAddr] = false;
        _addTx(4, accountInfo[msg.sender].idx, accountInfo[toEthAddr].idx, 0);
    }

    /**
     * @notice Allows a user to explode multiple accounts that have vouched for them
     * @dev Transfers a penalty amount from each exploded account to the caller
     * @param toEthAddrs Array of addresses to explode
     * @dev Reverts if any address in the array has not vouched for the caller
     * @dev Removes mutual vouch relationships and transfers penalty amounts
     * @dev Emits TxEvent with identifier 5 (explode) for each exploded account
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
            uint192 penalty = uint192(
                Math.min(explodeAmount, receiverInfo.balance - _MIN_BALANCE)
            );
            unchecked {
                accountInfo[toEthAddr].balance = receiverInfo.balance - penalty;
            }
            accountInfo[msg.sender].balance = senderInfo.balance + penalty;
            vouches[toEthAddr][msg.sender] = false;
            vouches[msg.sender][toEthAddr] = false;
            _addTx(5, senderInfo.idx, receiverInfo.idx, uint128(penalty));
        }
    }

    /**
     * @notice Processes a batch of transactions and verifies the associated ZK proof
     * @dev Verifies the state transition using zero-knowledge proofs
     * @param newAccountRoot The new account merkle root after processing the batch
     * @param newVouchRoot The new vouch merkle root after processing the batch
     * @param newScoreRoot The new score merkle root after processing the batch
     * @param proofA First component of the ZK proof
     * @param proofB Second component of the ZK proof
     * @param proofC Third component of the ZK proof
     * @dev Reverts if batch is not full or proof verification fails
     * @dev Updates the merkle roots and emits ForgeBatch event
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

        Transaction[] memory transactions = new Transaction[](batchSize);
        for (uint256 i = 0; i < _MAX_TXNS; ++i) {
            transactions[i] = unprocessedBatchesMap[lastForgedTxn + i];
        }
        bytes memory txnData = abi.encode(transactions);

        uint256 input = _constructCircuitInput(
            newAccountRoot,
            newVouchRoot,
            newScoreRoot,
            txnData
        );

        // Verify the proof
        if (
            !verifier.verifyProof(
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

        emit ForgeBatch(lastForgedBatch, lastForgedTxn, batchSize, txnData);
    }

    /**
     * @notice Proves a user's score using a Merkle proof against a specific batch's score root
     * @dev Verifies the user's score using sparse Merkle tree verification
     * @param numScoreRoot The batch number containing the score root to verify against
     * @param targetIdx The user's account index in the tree
     * @param score The claimed score value
     * @param siblings Array of sibling hashes for the Merkle proof
     * @dev Reverts if the Merkle proof verification fails
     * @dev Updates the user's score snapshot upon successful verification
     */
    function proveScoreMerkleProof(
        uint32 numScoreRoot,
        uint24 targetIdx,
        uint32 score,
        uint256[] calldata siblings
    ) external {
        if (accountInfo[msg.sender].idx != targetIdx) {
            revert IncorrectAccountIndex();
        }
        uint256[1] memory arrayState;
        arrayState[0] = score;
        uint256 stateHash = _insPoseidonUnit1.poseidon(arrayState);
        uint256 scoreRoot = scoreRootMap[numScoreRoot];

        if (!_smtVerifier(scoreRoot, siblings, targetIdx, stateHash)) {
            revert SmtProofInvalid();
        }

        scoreSnapshots[msg.sender].batchNum = numScoreRoot;
        scoreSnapshots[msg.sender].score = score;
    }

    function demoSmTVerifier(
        uint256 scoreRoot,
        uint256[] calldata siblings,
        uint256 targetIdx,
        uint256 stateHash
    ) external view returns (bool) {
        return _smtVerifier(scoreRoot, siblings, targetIdx, stateHash);
    }


    function proveScoreMerkleProofDebug(
        uint32 numScoreRoot,
        uint24 targetIdx,
        uint32 score,
        uint256[] calldata siblings
    ) external {
        uint256[1] memory arrayState;
        arrayState[0] = score;
        uint256 stateHash = _insPoseidonUnit1.poseidon(arrayState);
        uint256 scoreRoot = scoreRootMap[numScoreRoot];
        bool result = _smtVerifierDebug(scoreRoot, siblings, targetIdx, stateHash);
        if (result) {
        scoreSnapshots[msg.sender].batchNum = numScoreRoot;
        scoreSnapshots[msg.sender].score = score;
        }

    }

    function updateScore(address user, uint32 score) external {
        scoreSnapshots[user].score = score;
    }

    /**
     * @notice Updates the amount used for the explode operation
     * @dev Only callable by accounts with ADMIN_ROLE
     * @param _explodeAmount The new explode penalty amount
     * @dev Emits ExplodeAmountUpdated event
     */
    function updateExplodeAmount(
        uint256 _explodeAmount
    ) external onlyRole(ADMIN_ROLE) {
        explodeAmount = _explodeAmount;
        emit ExplodeAmountUpdated(explodeAmount);
    }

    /**
     * @notice Updates the minimum balance required for scoring participation
     * @dev Only callable by accounts with ADMIN_ROLE
     * @param _scoringRequiredBalance The new minimum balance requirement
     * @dev Emits ScoringRequiredBalanceUpdated event
     */
    function updateScoringRequiredBalance(
        uint256 _scoringRequiredBalance
    ) external onlyRole(ADMIN_ROLE) {
        scoringRequiredBalance = _scoringRequiredBalance;
        emit ScoringRequiredBalanceUpdated(scoringRequiredBalance);
    }

    /**
     * @notice Retrieves the number of transactions waiting to be processed
     * @return The length of the transaction queue (pending transactions)
     */
    function getQueueLength() external view returns (uint256) {
        return lastAddedTxn - lastForgedTxn;
    }

    /**
     * @notice Retrieves a user's current score
     * @param user The address of the user
     * @return score The user's current score value
     */
    function getScore(address user) external view returns (uint32 score) {
        return scoreSnapshots[user].score;
    }

    /**
     * @dev Adds a transaction to the current filling batch
     * @param identifier Transaction type identifier (0-5)
     * @param from The sender's account index
     * @param to The recipient's account index  
     * @param amount The transaction amount
     * @dev Emits a TxEvent with transaction details
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
     * @dev Clears the processed batch from the transaction queue
     * @dev Deletes the transaction data and updates the lastForgedTxn pointer
     */
    function _clearBatchFromQueue() internal {
        for (uint256 i = 0; i < _MAX_TXNS; ++i) {
            delete unprocessedBatchesMap[lastForgedBatch + i];
        }
        lastForgedTxn = lastForgedTxn + batchSize;
    }



    /**
     * @dev Constructs the input for the verification circuit
     * @param newAccountRoot The new account root after state transition
     * @param newVouchRoot The new vouch root after state transition
     * @param newScoreRoot The new score root after state transition
     * @param txnData The encoded transaction data for the batch
     * @return The hashed input for circuit verification, reduced modulo _RFIELD
     */
    function _constructCircuitInput(
        uint256 newAccountRoot,
        uint256 newVouchRoot,
        uint256 newScoreRoot,
        bytes memory txnData
    ) internal view returns (uint256) {
        uint256 oldAccountRoot = accountRootMap[lastForgedBatch];
        uint256 oldVouchRoot = vouchRootMap[lastForgedBatch];
        uint256 oldScoreRoot = scoreRootMap[lastForgedBatch];

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

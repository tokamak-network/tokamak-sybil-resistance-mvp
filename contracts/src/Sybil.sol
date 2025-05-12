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

    struct Transaction {
        uint8 identifier;
        address from;
        address to;
        uint256 amount;
    }

    uint256 constant _TXN_TOTALBYTES = 73; // Total bytes per transaction
    uint256 constant _MAX_TXNS = 256; // Max transactions per batch
    uint256 constant _LIMIT_AMOUNT = (1 << 128); // Max loadAmount per call
    uint256 constant _RFIELD = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
    uint256 public _MIN_BALANCE = 1;
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    uint256 public explodeAmount = (1 << 50);
    uint256 public scoringRequiredBalance = (1 << 16);
    uint32 public lastForgedBatch;
    uint32 public currentFillingBatch;

    mapping(uint32 => uint256) public accountRootMap;
    mapping(uint32 => uint256) public vouchRootMap;
    mapping(uint32 => uint256) public scoreRootMap;
    mapping(uint32 => uint256) public exitRootMap;
    mapping(uint32 => Transaction[]) public unprocessedBatchesMap;
    mapping(uint32 => bytes32) public txsDataHashMap;
    mapping(address => uint256) public balances;
    mapping(address => mapping(address => bool)) public vouches;
    mapping(address => ScoreSnapshot) public scoreSnapshots;

    // Verifier
    Verifier public verifier;

    event L1UserTxEvent(
        uint32 indexed queueIndex,
        uint8 indexed position,
        bytes l1UserTx
    );
    event ForgeBatch(uint32 indexed batchNum, uint16 l1UserTxsLen);
    event WithdrawEvent(uint48 indexed idx, uint32 indexed numExitRoot);
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
        currentFillingBatch = 2;

        __AccessControl_init();
        _grantRole(ADMIN_ROLE, _adminRole);

        _initializeVerifiers(_verifier, maxTx, nLevel);

        _initializeHelpers(_poseidon2Elements, _poseidon3Elements);
    }

    function deposit() external payable {
        uint256 userBalance = balances[msg.sender];
        if (msg.value >= _LIMIT_AMOUNT) {
            revert LimitAmountExceeded();
        }
        if (msg.value < _MIN_BALANCE) {
            revert InsufficientETH();
        }
        if (userBalance == 0) {
            _addTx(0, msg.sender, address(0), msg.value);
        } else {
            _addTx(1, msg.sender, address(0), msg.value);
        }
        balances[msg.sender] = userBalance + msg.value;
    }

    function withdraw(uint256 amount) external {
        uint256 userBalance = balances[msg.sender];
        if (amount >= _LIMIT_AMOUNT) {
            revert LimitAmountExceeded();
        }
        if (amount + _MIN_BALANCE > userBalance) {
            revert InsufficientBalance();
        }
        
        unchecked {
            balances[msg.sender] = userBalance - amount;
        }
        (bool success, ) = msg.sender.call{value: amount}("");
        if (!success) {
            revert EthTransferFailed();
        }
        _addTx(2, msg.sender, address(0), amount);
    }

    /**
     * @dev Allows a user to vouch for another account.
     *
     * @param toEthAddr The index of the account that is being vouched.
     */
    function vouch(address toEthAddr) external {
        if (balances[msg.sender] == 0) {
            revert SenderHasZeroBalance();
        }
        if (toEthAddr == msg.sender) {
            revert SelfVouch();
        }
        if (balances[toEthAddr] == 0) {
            revert ReceiverHasZeroBalance();
        }
        vouches[msg.sender][toEthAddr] = true;
        _addTx(3, msg.sender, toEthAddr, 0);
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
        _addTx(4, msg.sender, toEthAddr, 0);
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

        for (uint256 i = 0; i < toEthAddrs.length; ++i) {
            address toEthAddr = toEthAddrs[i];
            uint256 userBalance = balances[toEthAddr];
            uint256 penalty = Math.min(
                explodeAmount,
                userBalance - _MIN_BALANCE
            );
            unchecked {
                balances[toEthAddr] = userBalance - penalty;
            }
            balances[msg.sender] = balances[msg.sender] + penalty;
            vouches[toEthAddr][msg.sender] = false;
            vouches[msg.sender][toEthAddr] = false;
            _addTx(5, msg.sender, toEthAddr, 0);
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

        lastForgedBatch++;
        accountRootMap[lastForgedBatch] = newAccountRoot;
        vouchRootMap[lastForgedBatch] = newVouchRoot;
        scoreRootMap[lastForgedBatch] = newScoreRoot;

        uint16 l1UserTxsLen = _clearBatchFromQueue();

        emit ForgeBatch(lastForgedBatch, l1UserTxsLen);
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
    function getQueueLength() external view returns (uint32) {
        return currentFillingBatch - lastForgedBatch;
    }

    /**
     * @dev Adds a transaction to the current filling batch.
     *
     * @param identifier It is used to identify the type of transaction.
     * @param from The Ethereum address who initiated the transaction.
     * @param to The receipient address associated with the transaction.
     * @param amount The amount of Ether.
     *
     * @dev Emits a {L1User TxEvent} event.
     */
    function _addTx(
        uint8 identifier,
        address from,
        address to,
        uint256 amount
    ) internal {
        Transaction memory transaction = Transaction(
            identifier,
            from,
            to,
            amount
        );
        unprocessedBatchesMap[currentFillingBatch].push(transaction);

        uint256 currentPosition = unprocessedBatchesMap[currentFillingBatch].length - 1;

        emit L1UserTxEvent(
            currentFillingBatch,
            uint8(currentPosition),
            abi.encodePacked(identifier, from, to, amount)
        );

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
        Transaction[] memory transactions = unprocessedBatchesMap[
            lastForgedBatch + 1
        ];

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

// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "./interfaces/INewSybil.sol";
import "./interfaces/IVerifier.sol";
import "./types/SybilHelpers.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";

contract NewSybil is Initializable, AccessControlUpgradeable, INewSybil, MVPSybilHelpers {

    struct Verifier {
        IVerifier verifierInterface;
        uint256 maxTx; // maximum rollup transactions in a batch: L1-tx transactions
        uint256 nLevel; // number of levels of the circuit
    }

    struct ScoreSnapshot {
		uint32 score;
		uint32 batchNum;
    }

    uint256 constant _TXN_TOTALBYTES = 73; // Total bytes per transaction
    uint256 constant _MAX_TXNS = 256; // Max transactions per batch
    uint256 constant _LIMIT_AMOUNT = (1 << 128); // Max loadAmount per call
    uint256 constant _RFIELD = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    uint256 public explodeAmount = (1 << 50);
    uint256 public minBalance = (1 << 1);
    uint32 public lastForgedBatch;
    uint32 public currentFillingBatch;

    mapping(uint32 => uint256) public accountRootMap;
    mapping(uint32 => uint256) public vouchRootMap;
    mapping(uint32 => uint256) public scoreRootMap;
    mapping(uint32 => uint256) public exitRootMap;
    mapping(uint32 => bytes) public unprocessedBatchesMap;
    mapping(uint32 => bytes32) public txsDataHashMap;
    mapping(address => uint256) public balances;
    mapping(address => mapping(address => bool)) public vouches;
    mapping(address => ScoreSnapshot) public scoreSnapshots;
    
    // Mapping of exit nullifiers, only allowing each withdrawal to be made once
    mapping(uint32 => mapping(uint48 => bool)) public exitNullifierMap;

    // Verifier
    Verifier public verifier;

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

        _initializeVerifiers(
            _verifier,
            maxTx,
            nLevel
        );

        _initializeHelpers(
            _poseidon2Elements,
            _poseidon3Elements
        );
    }

    function deposit() external payable override {
        if (msg.value >= _LIMIT_AMOUNT) {
            revert LimitAmountExceeded();
        }
        if (msg.value < minBalance) {
            revert InsufficientETH();
        }
        if(balances[msg.sender] == 0) {
            _addTx(0, msg.sender, address(0), msg.value);  
        } else {
            _addTx(1, msg.sender, address(0), msg.value);
        }
        balances[msg.sender] += msg.value;
    }

    function withdraw(uint256 amount) external {
        if (amount >= _LIMIT_AMOUNT) {
            revert LimitAmountExceeded();
        }
        // require(amount < _LIMIT_AMOUNT, LimitAmountExceeded());
        if (amount + minBalance > balances[msg.sender]) {
            revert InsufficientBalance();
        }
        balances[msg.sender] -= amount;
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
        // require(balances[msg.sender] != 0, SenderHasZeroBalance());
        if (balances[toEthAddr] == 0) {
            revert ReceiverHasZeroBalance();
        }
        // require(balances[toEthAddr] != 0, ReceiverHasZeroBalance());
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
                uint256 penalty = Math.min(explodeAmount, balances[toEthAddr] - minBalance);
                balances[toEthAddr] -= penalty;
                balances[msg.sender] += penalty;
                vouches[toEthAddr][msg.sender] = false;
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
    ) external override {
        uint256 input = _constructCircuitInput(
          newAccountRoot,
          newVouchRoot,
          newScoreRoot
      );

        // Verify the proof using the specific rollup verifier
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
        uint256[2] memory arrayState = _buildTreeState(
            score,
            msg.sender
        );
        uint256 stateHash = _hash2Elements(arrayState);
        uint256 scoreRoot = scoreRootMap[numScoreRoot];
        
        if(!_smtVerifier(scoreRoot, siblings, idx, stateHash)) {
            revert SmtProofInvalid();
        }
        // require(_smtVerifier(scoreRoot, siblings, idx, stateHash), SmtProofInvalid());
        
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
        emit MinBalanceUpdated(minBalance);
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
     * @param identifer It is used to identify the type of transaction.
     * @param from The Ethereum address who initiated the transaction.
     * @param to The receipient address associated with the transaction.
     * @param amount The amount of Ether.
     *
     * @dev Emits a {L1User TxEvent} event.
    */
    function _addTx(
        uint256 identifer,
        address from,
        address to,
        uint256 amount
    ) internal  {
        bytes memory l1Tx = abi.encodePacked(
            identifer,
            from,
            to,
            amount
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
    // function _safeTransfer(uint256 value) internal {
    //     (bool success, ) = msg.sender.call{value: value}(new bytes(0));
    //     if (!success) {
    //         revert EthTransferFailed();
    //     }
    // }

    /**
    //  * @dev Withdraws a specified amount of funds from the contract.
    //  *
    //  * @param amount The amount of Ether to withdraw, specified in wei.
    // */
    // function _withdrawFunds(uint192 amount) internal {
    //     _safeTransfer(amount);
    // }

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
        bytes memory txnData = unprocessedBatchesMap[lastForgedBatch+1];

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

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "./interfaces/ISybil.sol";

/**
 * @title CummulativeScore
 * @author Aryan Soni
 * @notice This contract demonstrates how to interface with the DepositManager's accStaked function
 * @dev This contract provides read access to staking information from the DepositManager
 */
/// @notice Interface for the DepositManager contract
interface IDepositManager {
    function accStaked(address layer2, address account) external view returns (uint256 wtonAmount);
    function accStakedLayer2(address layer2) external view returns (uint256 wtonAmount);
    function accStakedAccount(address account) external view returns (uint256 wtonAmount);
    function pendingUnstaked(address layer2, address account) external view returns (uint256 wtonAmount);
    function accUnstaked(address layer2, address account) external view returns (uint256 wtonAmount);
}

contract CummulativeScore is AccessControl {

    /// @notice Admin role identifier for access control
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    
    /// @notice Address of the DepositManager contract
    address public depositManager;
    
    /// @notice Address of the Sybil contract
    address public sybilContract;
    
    /// @notice Event emitted when DepositManager address is updated
    event DepositManagerUpdated(address indexed oldManager, address indexed newManager);
    
    /// @notice Event emitted when Sybil contract address is updated
    event SybilContractUpdated(address indexed oldSybil, address indexed newSybil);
    

    /**
     * @notice Constructor to initialize the contract with the DepositManager and Sybil contract addresses
     * @param _depositManager Address of the DepositManager contract
     * @param _sybilContract Address of the Sybil contract
     * @param _adminRole Address that will be granted admin privileges
     */
    constructor(
        address _depositManager,
        address _sybilContract,
        address _adminRole
    ) {
        require(_depositManager != address(0), "Staking: DepositManager cannot be zero address");
        require(_sybilContract != address(0), "Staking: Sybil contract cannot be zero address");
        require(_adminRole != address(0), "Staking: Admin role cannot be zero address");
        
        _grantRole(ADMIN_ROLE, _adminRole);
        _grantRole(DEFAULT_ADMIN_ROLE, _adminRole);
        
        depositManager = _depositManager;
        sybilContract = _sybilContract;
        
        emit DepositManagerUpdated(address(0), _depositManager);
        emit SybilContractUpdated(address(0), _sybilContract);
    }

    /**
     * @notice Updates the DepositManager contract address
     * @param _newDepositManager New DepositManager contract address
     * @dev Only callable by accounts with ADMIN_ROLE
     */
    function updateDepositManager(address _newDepositManager) external onlyRole(ADMIN_ROLE) {
        require(_newDepositManager != address(0), "Staking: DepositManager cannot be zero address");
        
        address oldManager = depositManager;
        depositManager = _newDepositManager;
        
        emit DepositManagerUpdated(oldManager, _newDepositManager);
    }

    /**
     * @notice Updates the Sybil contract address
     * @param _newSybilContract New Sybil contract address
     * @dev Only callable by accounts with ADMIN_ROLE
     */
    function updateSybilContract(address _newSybilContract) external onlyRole(ADMIN_ROLE) {
        require(_newSybilContract != address(0), "Staking: Sybil contract cannot be zero address");
        
        address oldSybil = sybilContract;
        sybilContract = _newSybilContract;
        
        emit SybilContractUpdated(oldSybil, _newSybilContract);
    }

    /**
     * @notice Gets the staked amount for a specific layer2 and account
     * @param layer2 Address of the layer2 contract
     * @param account Address of the account to query
     * @return wtonAmount The amount of WTON staked by the account in the layer2 (in whole tokens)
     */
    function getAccStaked(address layer2, address account) external view returns (uint256 wtonAmount) {
        require(layer2 != address(0), "Staking: Layer2 cannot be zero address");
        require(account != address(0), "Staking: Account cannot be zero address");
        
        return IDepositManager(depositManager).accStaked(layer2, account) / 10**27;
    }

    /**
     * @notice Batch query for multiple accounts in a layer2
     * @param layer2 Address of the layer2 contract
     * @param accounts Array of account addresses to query
     * @return stakedAmounts Array of staked amounts corresponding to each account
     */
    function batchGetAccStaked(address layer2, address[] calldata accounts) 
        external 
        view 
        returns (uint256[] memory stakedAmounts) 
    {
        require(layer2 != address(0), "Staking: Layer2 cannot be zero address");
        require(accounts.length > 0, "Staking: Accounts array cannot be empty");
        
        stakedAmounts = new uint256[](accounts.length);
        IDepositManager manager = IDepositManager(depositManager);
        
        for (uint256 i = 0; i < accounts.length; i++) {
            require(accounts[i] != address(0), "Staking: Account cannot be zero address");
            stakedAmounts[i] = manager.accStaked(layer2, accounts[i]) / 10**27;
        }
    }

    /**
     * @notice Checks if an account has staked tokens in a layer2
     * @param layer2 Address of the layer2 contract
     * @param account Address of the account to check
     * @return hasStaked True if the account has staked tokens, false otherwise
     */
    function hasStakedTokens(address layer2, address account) external view returns (bool hasStaked) {
        require(layer2 != address(0), "Staking: Layer2 cannot be zero address");
        require(account != address(0), "Staking: Account cannot be zero address");
        
        uint256 stakedAmount = IDepositManager(depositManager).accStaked(layer2, account) / 10**27;
        return stakedAmount > 0;
    }

    /**
     * @notice Gets the cumulative score (staking score + sybil score) for an account
     * @param layer2 Address of the layer2 contract
     * @param account Address of the account to query
     * @return cumulativeScore The sum of staked WTON amount and Sybil score
     * @dev 1 WTON = 1 score point, so if staked = 100 WTON and sybil score = 50, cumulative = 150
     */
    function getCumulativeScore(address layer2, address account) external view returns (uint256 cumulativeScore) {
        require(layer2 != address(0), "Staking: Layer2 cannot be zero address");
        require(account != address(0), "Staking: Account cannot be zero address");
        
        // Get staking score (1 WTON = 1 score point)
        uint256 stakingScore = IDepositManager(depositManager).accStaked(layer2, account) / 10**27;
        
        // Get sybil score
        uint32 sybilScore = ISybil(sybilContract).getScore(account);
        
        // Return cumulative score
        cumulativeScore = stakingScore + uint256(sybilScore);
    }

    /**
     * @notice Gets detailed cumulative score breakdown for an account
     * @param layer2 Address of the layer2 contract
     * @param account Address of the account to query
     * @return stakingScore The staking score from DepositManager (WTON amount)
     * @return sybilScore The score from Sybil contract
     * @return cumulativeScore The sum of both scores
     */
    function getCumulativeScoreBreakdown(address layer2, address account) 
        external 
        view 
        returns (
            uint256 stakingScore,
            uint32 sybilScore,
            uint256 cumulativeScore
        ) 
    {
        require(layer2 != address(0), "Staking: Layer2 cannot be zero address");
        require(account != address(0), "Staking: Account cannot be zero address");
        
        // Get staking score (1 WTON = 1 score point)
        stakingScore = IDepositManager(depositManager).accStaked(layer2, account) / 10**27;
        
        // Get sybil score
        sybilScore = ISybil(sybilContract).getScore(account);
        
        // Calculate cumulative score
        cumulativeScore = stakingScore + uint256(sybilScore);
    }

    /**
     * @notice Batch query cumulative scores for multiple accounts
     * @param layer2 Address of the layer2 contract
     * @param accounts Array of account addresses to query
     * @return cumulativeScores Array of cumulative scores corresponding to each account
     */
    function batchGetCumulativeScores(address layer2, address[] calldata accounts) 
        external 
        view 
        returns (uint256[] memory cumulativeScores) 
    {
        require(layer2 != address(0), "Staking: Layer2 cannot be zero address");
        require(accounts.length > 0, "Staking: Accounts array cannot be empty");
        
        cumulativeScores = new uint256[](accounts.length);
        IDepositManager manager = IDepositManager(depositManager);
        ISybil sybil = ISybil(sybilContract);
        
        for (uint256 i = 0; i < accounts.length; i++) {
            require(accounts[i] != address(0), "Staking: Account cannot be zero address");
            
            uint256 stakingScore = manager.accStaked(layer2, accounts[i]) / 10**27;
            uint32 sybilScore = sybil.getScore(accounts[i]);
            cumulativeScores[i] = stakingScore + uint256(sybilScore);
        }
    }

    /**
     * @notice Gets the DepositManager contract address
     * @return The address of the DepositManager contract
     */
    function getDepositManager() external view returns (address) {
        return depositManager;
    }

    /**
     * @notice Gets the Sybil contract address
     * @return The address of the Sybil contract
     */
    function getSybilContract() external view returns (address) {
        return sybilContract;
    }
} 
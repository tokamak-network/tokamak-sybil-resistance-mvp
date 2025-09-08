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
    
    /// @notice Token name
    string public constant name = "Cumulative Score Token";
    
    /// @notice Token symbol
    string public constant symbol = "CST";
    
    /// @notice Token decimals (0 since scores are whole numbers)
    uint8 public constant decimals = 0;
    
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
        require(_depositManager != address(0), "CummulativeScore: DepositManager cannot be zero address");
        require(_sybilContract != address(0), "CummulativeScore: Sybil contract cannot be zero address");
        require(_adminRole != address(0), "CummulativeScore: Admin role cannot be zero address");
        
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
        require(_newDepositManager != address(0), "CummulativeScore: DepositManager cannot be zero address");
        
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
        require(_newSybilContract != address(0), "CummulativeScore: Sybil contract cannot be zero address");
        
        address oldSybil = sybilContract;
        sybilContract = _newSybilContract;
        
        emit SybilContractUpdated(oldSybil, _newSybilContract);
    }

    /**
     * @notice Gets the total cumulative balance for an account across all layer2s (ERC-20 standard)
     * @param account Address of the account to query
     * @return balance The total cumulative score (staking + sybil) for the account
     */
    function balanceOf(address account) external view returns (uint256 balance) {
        require(account != address(0), "CummulativeScore: Account cannot be zero address");
        
        // Get staking score (1 WTON = 1 score point)
        uint256 stakingScore = IDepositManager(depositManager).accStakedAccount(account) / 10**27;
        
        // Get sybil score
        uint32 sybilScore = ISybil(sybilContract).getScore(account);
        
        // Return total balance
        balance = stakingScore + uint256(sybilScore);
    }

    /**
     * @notice Gets the staking balance for a specific layer2 and account
     * @param layer2 Address of the layer2 contract
     * @param account Address of the account to query
     * @return balance The amount of WTON staked by the account in the layer2 (in whole tokens)
     */
    function stakingLayer2BalanceOf(address layer2, address account) external view returns (uint256 balance) {
        require(layer2 != address(0), "CummulativeScore: Layer2 cannot be zero address");
        require(account != address(0), "CummulativeScore: Account cannot be zero address");
        
        return IDepositManager(depositManager).accStaked(layer2, account) / 10**27;
    }

    function totalSupply() external pure returns (uint256 supply) {
        return type(uint256).max;
    }

    /**
     * @notice Batch query for multiple accounts in a layer2
     * @param layer2 Address of the layer2 contract
     * @param accounts Array of account addresses to query
     * @return balances Array of staking balances corresponding to each account
     */
    function stakingLayer2BalanceOfBatch(address layer2, address[] calldata accounts) 
        external 
        view 
        returns (uint256[] memory balances) 
    {
        require(layer2 != address(0), "CummulativeScore: Layer2 cannot be zero address");
        require(accounts.length > 0, "CummulativeScore: Accounts array cannot be empty");
        
        balances = new uint256[](accounts.length);
        IDepositManager manager = IDepositManager(depositManager);
        
        for (uint256 i = 0; i < accounts.length; i++) {
            require(accounts[i] != address(0), "CummulativeScore: Account cannot be zero address");
            balances[i] = manager.accStaked(layer2, accounts[i]) / 10**27;
        }
    }

    /**
     * @notice Checks if an account has any staking balance (ERC-20 style)
     * @param layer2 Address of the layer2 contract
     * @param account Address of the account to check
     * @return hasBalance True if the account has staked tokens, false otherwise
     */
    function hasBalance(address layer2, address account) external view returns (bool) {
        require(layer2 != address(0), "CummulativeScore: Layer2 cannot be zero address");
        require(account != address(0), "CummulativeScore: Account cannot be zero address");
        
        uint256 stakedAmount = IDepositManager(depositManager).accStaked(layer2, account) / 10**27;
        return stakedAmount > 0;
    }

    /**
     * @notice Gets the total balance (staking of layer2 + sybil score) for an account (ERC-20 style)
     * @param layer2 Address of the layer2 contract
     * @param account Address of the account to query
     * @return totalBalance The sum of staked WTON amount and Sybil score
     * @dev 1 WTON = 1 score point, so if staked = 100 WTON and sybil score = 50, total = 150
     */
    function totalLayer2BalanceOf(address layer2, address account) external view returns (uint256 totalBalance) {
        require(layer2 != address(0), "CummulativeScore: Layer2 cannot be zero address");
        require(account != address(0), "CummulativeScore: Account cannot be zero address");
        
        // Get staking score (1 WTON = 1 score point)
        uint256 stakingScore = IDepositManager(depositManager).accStaked(layer2, account) / 10**27;
        
        // Get sybil score
        uint32 sybilScore = ISybil(sybilContract).getScore(account);
        
        // Return total balance
        totalBalance = stakingScore + uint256(sybilScore);
    }

    /**
     * @notice Gets detailed balance breakdown for an account (ERC-20 style)
     * @param account Address of the account to query
     * @return stakingBalance The staking balance from DepositManager (WTON amount)
     * @return sybilBalance The score from Sybil contract
     * @return totalBalance The sum of both balances
     */
    function balanceBreakdown(address account) 
        external 
        view 
        returns (
            uint256 stakingBalance,
            uint32 sybilBalance,
            uint256 totalBalance
        ) 
    {
        require(account != address(0), "CummulativeScore: Account cannot be zero address");
        
        // Get staking balance (1 WTON = 1 score point)
        stakingBalance = IDepositManager(depositManager).accStakedAccount(account) / 10**27;
        
        // Get sybil balance
        sybilBalance = ISybil(sybilContract).getScore(account);
        
        // Calculate total balance
        totalBalance = stakingBalance + uint256(sybilBalance);
    }

    /**
     * @notice Batch query total balances for multiple accounts (ERC-20 style)
     * @param layer2 Address of the layer2 contract
     * @param accounts Array of account addresses to query
     * @return totalBalances Array of total balances corresponding to each account
     */
    function totalBalanceOfBatch(address layer2, address[] calldata accounts) 
        external 
        view 
        returns (uint256[] memory totalBalances) 
    {
        require(layer2 != address(0), "CummulativeScore: Layer2 cannot be zero address");
        require(accounts.length > 0, "CummulativeScore: Accounts array cannot be empty");
        
        totalBalances = new uint256[](accounts.length);
        IDepositManager manager = IDepositManager(depositManager);
        ISybil sybil = ISybil(sybilContract);
        
        for (uint256 i = 0; i < accounts.length; i++) {
            require(accounts[i] != address(0), "CummulativeScore: Account cannot be zero address");
            
            uint256 stakingScore = manager.accStaked(layer2, accounts[i]) / 10**27;
            uint32 sybilScore = sybil.getScore(accounts[i]);
            totalBalances[i] = stakingScore + uint256(sybilScore);
        }
    }
} 
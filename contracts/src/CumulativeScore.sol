// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "./interfaces/ISybil.sol";

/**
 * @title CumulativeScore
 * @author Aryan Soni
 * @notice This contract demonstrates how to interface with the ISeigManager's stakeOf function
 * @dev This contract provides read access to staking information from the SeigManager
 */
/// @notice Interface for the SeigManager contract
interface ISeigManager {
    function stakeOf(address account) external view returns (uint256);
}

contract CumulativeScore is AccessControl {

    /// @notice Admin role identifier for access control
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    
    /// @notice Token name
    string public constant name = "Cumulative Score Token";
    
    /// @notice Token symbol
    string public constant symbol = "CST";
    
    /// @notice Token decimals (0 since scores are whole numbers)
    uint8 public constant decimals = 0;
    
    /// @notice Address of the SeigManager contract
    address public seigManager;
    
    /// @notice Address of the Sybil contract
    address public sybilContract;
    
    /// @notice Event emitted when SeigManager address is updated
    event SeigManagerUpdated(address indexed oldManager, address indexed newManager);
    
    /// @notice Event emitted when Sybil contract address is updated
    event SybilContractUpdated(address indexed oldSybil, address indexed newSybil);
    

    /**
     * @notice Constructor to initialize the contract with the SeigManager and Sybil contract addresses
     * @param _seigManager Address of the SeigManager contract
     * @param _sybilContract Address of the Sybil contract
     * @param _adminRole Address that will be granted admin privileges
     */
    constructor(
        address _seigManager,
        address _sybilContract,
        address _adminRole
    ) {
        require(_seigManager != address(0), "CumulativeScore: SeigManager cannot be zero address");
        require(_sybilContract != address(0), "CumulativeScore: Sybil contract cannot be zero address");
        require(_adminRole != address(0), "CumulativeScore: Admin role cannot be zero address");
        
        _grantRole(ADMIN_ROLE, _adminRole);
        _grantRole(DEFAULT_ADMIN_ROLE, _adminRole);
        
        seigManager = _seigManager;
        sybilContract = _sybilContract;
        
        emit SeigManagerUpdated(address(0), _seigManager);
        emit SybilContractUpdated(address(0), _sybilContract);
    }

    /**
     * @notice Updates the SeigManager contract address
     * @param _newSeigManager New SeigManager contract address
     * @dev Only callable by accounts with ADMIN_ROLE
     */
    function updateSeigManager(address _newSeigManager) external onlyRole(ADMIN_ROLE) {
        require(_newSeigManager != address(0), "CumulativeScore: SeigManager cannot be zero address");
        
        address oldManager = seigManager;
        seigManager = _newSeigManager;
        
        emit SeigManagerUpdated(oldManager, _newSeigManager);
    }

    /**
     * @notice Updates the Sybil contract address
     * @param _newSybilContract New Sybil contract address
     * @dev Only callable by accounts with ADMIN_ROLE
     */
    function updateSybilContract(address _newSybilContract) external onlyRole(ADMIN_ROLE) {
        require(_newSybilContract != address(0), "CumulativeScore: Sybil contract cannot be zero address");
        
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
        require(account != address(0), "CumulativeScore: Account cannot be zero address");
        
        // Get staking score (1 WTON = 1 score point)
        uint256 stakingScore = ISeigManager(seigManager).stakeOf(account) / 10**27;
        
        // Get sybil score
        uint32 sybilScore = ISybil(sybilContract).getScore(account);
        
        // Return total balance
        balance = stakingScore + uint256(sybilScore);
    }


    function totalSupply() external pure returns (uint256 supply) {
        return type(uint256).max;
    }

    /**
     * @notice Gets detailed balance breakdown for an account (ERC-20 style)
     * @param account Address of the account to query
     * @return stakingBalance The staking balance from SeigManager (WTON amount)
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
        require(account != address(0), "CumulativeScore: Account cannot be zero address");
        
        // Get staking balance (1 WTON = 1 score point)
        stakingBalance = ISeigManager(seigManager).stakeOf(account) / 10**27;
        
        // Get sybil balance
        sybilBalance = ISybil(sybilContract).getScore(account);
        
        // Calculate total balance
        totalBalance = stakingBalance + uint256(sybilBalance);
    }

    /**
     * @notice Batch query total balances for multiple accounts
     * @param accounts Array of account addresses to query
     * @return totalBalances Array of total balances corresponding to each account
     */
    function totalBalanceOfBatch(address[] calldata accounts) 
        external 
        view 
        returns (uint256[] memory totalBalances) 
    {
        require(accounts.length > 0, "CumulativeScore: Accounts array cannot be empty");
        
        totalBalances = new uint256[](accounts.length);
        ISeigManager manager = ISeigManager(seigManager);
        ISybil sybil = ISybil(sybilContract);
        
        for (uint256 i = 0; i < accounts.length; i++) {
            require(accounts[i] != address(0), "CumulativeScore: Account cannot be zero address");
            
            uint256 stakingScore = manager.stakeOf(accounts[i]) / 10**27;
            uint32 sybilScore = sybil.getScore(accounts[i]);
            totalBalances[i] = stakingScore + uint256(sybilScore);
        }
    }
} 
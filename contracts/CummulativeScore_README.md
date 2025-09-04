# CummulativeScore Contract

This contract demonstrates how to interface with the [Tokamak Network's DepositManager](https://github.com/tokamak-network/ton-staking-v2) `accStaked` function and combines it with Sybil resistance scores to provide cumulative scoring.

## Overview

The `CummulativeScore` contract provides a clean interface to calculate cumulative scores by combining:
- **Staking scores** from the DepositManager contract (1 WTON = 1 score point)
- **Sybil resistance scores** from the Sybil contract

Key features include:
- Query individual account staking amounts (converted to whole tokens)
- Calculate cumulative scores (staking + sybil)
- Batch query multiple accounts
- Detailed score breakdowns
- Access control for administrative functions

## Features

### Core Functions

1. **getAccStaked(address layer2, address account)** - Gets the staked amount for a specific layer2 and account (in whole tokens)
2. **getCumulativeScore(address layer2, address account)** - Gets the total score (staking + sybil) for an account
3. **getCumulativeScoreBreakdown(address layer2, address account)** - Gets detailed breakdown of staking score, sybil score, and total
4. **hasStakedTokens(address layer2, address account)** - Check if an account has any staked tokens
5. **batchGetAccStaked(address layer2, address[] accounts)** - Batch query staking amounts for multiple accounts
6. **batchGetCumulativeScores(address layer2, address[] accounts)** - Batch query cumulative scores for multiple accounts

### Administration

- **updateDepositManager(address _newDepositManager)** - Update the DepositManager contract address (Admin only)
- **updateSybilContract(address _newSybilContract)** - Update the Sybil contract address (Admin only)
- **getDepositManager()** - Get current DepositManager address
- **getSybilContract()** - Get current Sybil contract address

### Token Conversion

The contract automatically converts the 27-decimal format returned by DepositManager to whole tokens:
- **Raw DepositManager value**: `22000000000000000000000000000` (27 decimals)
- **Converted to score**: `22` points (divided by 10^27)

## Deployment

### Prerequisites

1. You need the address of the deployed DepositManager contract from the ton-staking-v2 repository
2. You need the address of the deployed Sybil contract
3. Set up your environment variables:
   ```bash
    DEPOSIT_MANAGER=0xYourDepositManagerAddress
    SYBIL_CONTRACT_ADDRESS=0xYourSybilContractAddress
   ```

### Deploy using Forge

```bash
# Deploy CummulativeScore contract
forge script script/DeployCummulativeScore.s.sol:DeployStakingInterface --rpc-url $RPC_URL --broadcast --verify
```

### Verify Contract

```bash
forge verify-contract <CONTRACT_ADDRESS> \
  src/CummulativeScore.sol:CummulativeScore \
  --etherscan-api-key <API_KEY> \
  --rpc-url <RPC_URL> \
  --compiler-version 0.8.24 \
  --constructor-args $(cast abi-encode "constructor(address,address,address)" <DEPOSIT_MANAGER> <SYBIL_CONTRACT> <ADMIN_ADDRESS>)
```

## Architecture

```
┌─────────────────┐    queries   ┌─────────────────┐
│ CummulativeScore│ ──────────> │ DepositManager  │
│                 │             │ (staking data)  │
│                 │    queries   ├─────────────────┤
│                 │ ──────────> │ Sybil Contract  │
│                 │             │ (sybil scores)  │
└─────────────────┘             └─────────────────┘
         │                               │
         │                               │
    Users/DApps                   Combined scoring:
    query this for               Staking + Sybil = Total
    cumulative scores
```

## Scoring System

- **Staking Score**: 1 WTON token = 1 score point
- **Sybil Score**: Retrieved from Sybil contract (uint32)
- **Cumulative Score**: Staking Score + Sybil Score
- **Token Conversion**: Raw DepositManager values (27 decimals) are divided by 10^27

### Example:
- User stakes: `22000000000000000000000000000` (raw value)
- Converted staking score: `22` points
- Sybil score: `50` points
- **Total cumulative score: `72` points**


## Security Considerations

1. **Access Control**: Admin functions are protected by OpenZeppelin's AccessControl
   - `ADMIN_ROLE`: Can update DepositManager and Sybil contract addresses
   - `DEFAULT_ADMIN_ROLE`: Can manage role assignments
2. **Zero Address Checks**: All functions validate input addresses
3. **Read-Only**: This contract only reads from external contracts, doesn't modify their state
4. **Integer Division**: Uses integer division for token conversion (fractional tokens are truncated)

## Integration Requirements

This contract requires integration with:

1. **DepositManager** from [ton-staking-v2 repository](https://github.com/tokamak-network/ton-staking-v2)
2. **Sybil Contract** that implements the ISybil interface. 
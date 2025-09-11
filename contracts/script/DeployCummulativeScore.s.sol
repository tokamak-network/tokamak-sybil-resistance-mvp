// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/Script.sol";
import { CumulativeScore } from "../src/CumulativeScore.sol";

contract DeployStakingInterface is Script {

    function run() external {
        // Get required addresses from environment variables
        address seigManager = vm.envAddress("Seig_Manager");
        address sybilContract = vm.envAddress("SYBIL_CONTRACT_ADDRESS");
        address adminRole = msg.sender;

        vm.startBroadcast();
        
        // Deploy the Staking contract
        CumulativeScore stakingContract = new CumulativeScore(
            seigManager,
            sybilContract,
            adminRole
        );

        vm.stopBroadcast();

        console2.log("Staking contract is deployed at:", address(stakingContract));
        console2.log("SeigManager address:", seigManager);
        console2.log("Sybil contract address:", sybilContract);
        console2.log("Admin role granted to:", adminRole);
    }
} 
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import {Sybil} from "../src/Sybil.sol";

contract TestSyncEvents is Script {

    function run() external {
        // Existing deployed Sybil contract address
        address sybilContractAddress = vm.envAddress("SYBIL_CONTRACT_ADDRESS");

        // Using the Deployed Sybil contract
        vm.startBroadcast(msg.sender);
        Sybil sybilContract = Sybil(sybilContractAddress);

        sybilContract.deposit{value: 10 wei}();

        console2.log("Sybil contract address", address(sybilContract));

        vm.stopBroadcast();
    }
}

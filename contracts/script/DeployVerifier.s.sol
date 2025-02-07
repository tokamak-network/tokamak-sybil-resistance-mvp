// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import {
    Sybil
} from "../src/mvp/Sybil.sol";
import {
    Verifier
} from "../src/Verifier.sol";

contract MyScript is Script {
    function run() external {
        // Deploy the Verifier contract
        vm.startBroadcast();
        Verifier verifier = new Verifier();
        vm.stopBroadcast();

        console2.log("VerifierRollupStub deployed at:", address(verifier));
    }
}
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import {
    Sybil
} from "../src/mvp/Sybil.sol";
import {
    VerifierRollupStub
} from "../src/VerifierRollupStub.sol";

contract MyScript is Script {
    function run() external {
        // Deploy the VerifierRollupStub contract
        vm.startBroadcast();
        VerifierRollupStub verifier = new VerifierRollupStub();
        vm.stopBroadcast();

        console2.log("VerifierRollupStub deployed at:", address(verifier));
    }
}
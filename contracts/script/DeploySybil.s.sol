// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {DevOpsTools} from "lib/foundry-devops/src/DevOpsTools.sol";
import "forge-std/Script.sol";
import {Sybil} from "../src/mvp/Sybil.sol";
import {VerifierRollupStub} from "../src/stub/VerifierRollupStub.sol";

contract FunctionScript is Script {
    error VerifierRollupStubNotDeployed();

    function run() external {
        address verifier = DevOpsTools.get_most_recent_deployment(
            "VerifierRollupStub",
            block.chainid
        );
        uint256 maxTx = vm.envUint("MAXTX");
        uint256 nLevel = vm.envUint("NLEVEL");

        // Specify Poseidon contract addresses
        address poseidon2Elements = vm.envAddress("POSEIDON2ELEMENTS");
        address poseidon3Elements = vm.envAddress("POSEIDON3ELEMENTS");
        address poseidon4Elements = vm.envAddress("POSEIDON4ELEMENTS");

        vm.startBroadcast();
        // Deploy the Sybil contract
        Sybil sybilContract = new Sybil();

        // Calling initialize at the time of deployment
        sybilContract.initialize(
            verifier,
            maxTx,
            nLevel,
            poseidon2Elements,
            poseidon3Elements,
            poseidon4Elements
        );

        vm.stopBroadcast();

        console2.log("Sybil contract is deployed at:", address(sybilContract));
    }
}

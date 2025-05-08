// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import { NewSybil } from "../src/NewSybil.sol";
import {Verifier} from "../src/Verifier.sol";

contract FunctionScript is Script {

    function run() external {
        address verifier = vm.envAddress("VERIFIER");
        uint256 maxTx = vm.envUint("MAXTX");
        uint256 nLevel = vm.envUint("NLEVEL");

        // Specify Poseidon contract addresses
        address poseidon2Elements = vm.envAddress("POSEIDON2ELEMENTS");
        address poseidon3Elements = vm.envAddress("POSEIDON3ELEMENTS");
        address adminRole = msg.sender;

        vm.startBroadcast();
        // Deploy the Sybil contract
        NewSybil sybilContract = new NewSybil();

        // Calling initialize at the time of deployment
        sybilContract.initialize(
            verifier,
            maxTx,
            nLevel,
            poseidon2Elements,
            poseidon3Elements,
            adminRole
        );

        vm.stopBroadcast();

        console2.log("Sybil contract is deployed at:", address(sybilContract));
    }
}
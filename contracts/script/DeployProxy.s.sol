// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import "forge-std/Script.sol";
import {Sybil} from "../src/Sybil.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract DeploySybil is Script {
    function run() external {
        address verifier = vm.envAddress("VERIFIER");
        uint256 maxTx = vm.envUint("MAXTX");
        uint256 nLevel = vm.envUint("NLEVEL");

        // Specify Poseidon contract addresses
        address poseidon2Elements = vm.envAddress("POSEIDON2ELEMENTS");
        address poseidon3Elements = vm.envAddress("POSEIDON3ELEMENTS");
        
        address adminRole = msg.sender;
        address initialOwner = msg.sender;
        vm.startBroadcast();

        Sybil sybil = new Sybil();

        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(sybil),
            initialOwner,
            abi.encodeWithSelector(
                Sybil(sybil).initialize.selector,
                verifier,
                maxTx,  
                nLevel,
                poseidon2Elements, 
                poseidon3Elements,
                adminRole 
            )
        );
            
        vm.stopBroadcast();
        console.log("Sybil address:", address(sybil));
        console.log("Proxy address:", address(proxy));
    }
}

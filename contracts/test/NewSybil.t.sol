// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import "forge-std/Test.sol";
import "../src/NewSybil.sol";
import "../src/interfaces/INewSybil.sol";
import "./utils/Constants.sol";
import "./types/NewTransactionTypes.sol";
import "../src/Verifier.sol";

contract MockPoseidon2 is PoseidonUnit2 {
    function poseidon(
        uint256[2] memory input
    ) external pure override returns (uint256) {}
}

contract MockPoseidon3 is PoseidonUnit3 {
    function poseidon(
        uint256[3] memory input
    ) external pure override returns (uint256) {}
}

contract MvpTest is Test, NewTransactionTypeHelper {
    NewSybil public sybil;
    bytes32[] public hashes;

    function setup() public {
        PoseidonUnit2 mockPoseidon2 = new MockPoseidon2();
        PoseidonUnit3 mockPoseidon3 = new MockPoseidon3();
        emit log_address(address(mockPoseidon2));
        emit log_address(address(mockPoseidon3));

        Verifier verifierStub = new Verifier();

        address verifiers = address(verifierStub);
        address adminRole = address(this);
        uint256 maxTx = uint256(256);
        uint256 nLevels = uint256(1);

        sybil = new NewSybil();

        sybil.initialize(
            verifiers, 
            maxTx, 
            nLevels, 
            address(mockPoseidon2), 
            address(mockPoseidon3), 
            adminRole
        );
    }
}
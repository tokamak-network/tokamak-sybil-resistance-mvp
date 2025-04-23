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

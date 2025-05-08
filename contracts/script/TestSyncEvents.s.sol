// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import {NewSybil} from "../src/NewSybil.sol";

contract TestSyncEvents is Script {
    struct TxParams {
        uint48 fromIdx;
        uint40 loadAmountF;
        uint40 amountF;
        uint48 toIdx;
    }

    function run() external {
        // Existing deployed Sybil contract address
        address sybilContractAddress = vm.envAddress("SYBIL_CONTRACT_ADDRESS");

        vm.startBroadcast();
        // Using the Deployed Sybil contract
        NewSybil sybilContract = NewSybil(sybilContractAddress);

        console2.log("Sybil contract is deployed at:", address(sybilContract));

        TxParams memory params1 = validCreateAccountDeposit();
        uint256 loadAmount1 = _float2Fix(params1.loadAmountF);

        sybilContract.deposit{value: loadAmount1}();

        vm.stopBroadcast();
    }

    function validCreateAccountDeposit() public pure returns (TxParams memory) {
        return TxParams({fromIdx: 0, loadAmountF: 2, amountF: 0, toIdx: 0});
    }

    function _float2Fix(uint40 floatVal) internal pure returns (uint256) {
        uint256 m = floatVal & 0x7FFFFFFFF;
        uint256 e = floatVal >> 35;

        uint256 exp = 10 ** e;
        uint256 fix = m * exp;

        return fix;
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import {Sybil} from "../src/Sybil.sol";

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
        uint256 maxTx = vm.envUint("MAXTX");
        uint256 nLevel = vm.envUint("NLEVEL");

        vm.startBroadcast();
        // Using the Deployed Sybil contract
        Sybil sybilContract = Sybil(sybilContractAddress);

        console2.log("Sybil contract is deployed at:", address(sybilContract));

        TxParams memory params1 = validCreateAccountDeposit();
        uint256 loadAmount1 = _float2Fix(params1.loadAmountF);

        sybilContract.createAccountDeposit{value: loadAmount1}(
            params1.loadAmountF
        );

        // Test deposit
        TxParams memory params2 = validDeposit();
        uint256 loadAmount2 = _float2Fix(params2.loadAmountF);
        sybilContract.deposit{value: loadAmount2}(
            params2.fromIdx,
            params2.loadAmountF
        );

        // Test vouch
        TxParams memory params3 = validVouch();
        sybilContract.vouch(params3.fromIdx, params3.toIdx);

        // Test unvouch
        TxParams memory params4 = validUnvouch();
        sybilContract.unvouch(params4.fromIdx, params4.toIdx);

        // Test exit
        TxParams memory params5 = validExit();
        sybilContract.exit(params5.fromIdx, params5.amountF);

        vm.stopBroadcast();
    }

    function validCreateAccountDeposit() public pure returns (TxParams memory) {
        return TxParams({fromIdx: 0, loadAmountF: 2, amountF: 0, toIdx: 0});
    }

    function validDeposit() public pure returns (TxParams memory) {
        return TxParams({fromIdx: 1, loadAmountF: 3, amountF: 0, toIdx: 0});
    }

    function validVouch() public pure returns (TxParams memory) {
        return TxParams({fromIdx: 1, loadAmountF: 0, amountF: 0, toIdx: 2});
    }

    function validUnvouch() public pure returns (TxParams memory) {
        return TxParams({fromIdx: 1, loadAmountF: 0, amountF: 0, toIdx: 2});
    }

    function validExit() public pure returns (TxParams memory) {
        return TxParams({fromIdx: 1, loadAmountF: 0, amountF: 1, toIdx: 1});
    }

    function _float2Fix(uint40 floatVal) internal pure returns (uint256) {
        uint256 m = floatVal & 0x7FFFFFFFF;
        uint256 e = floatVal >> 35;

        uint256 exp = 10 ** e;
        uint256 fix = m * exp;

        return fix;
    }
}

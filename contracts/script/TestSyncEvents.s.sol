// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import "forge-std/Script.sol";
import {Sybil} from "../src/mvp/Sybil.sol";

contract TestSyncEvents is Script {
    struct TxParams {
        uint48 fromIdx;
        uint40 loadAmountF;
        uint40 amountF;
        uint48 toIdx;
    }

    function run() external {
        // Existing deployed Sybil contract address
        address sybilContractAddress = vm.envAddress("SYBILCONTRACTADDRESS");
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

        vm.stopBroadcast();
    }

    function validCreateAccountDeposit() public pure returns (TxParams memory) {
        return TxParams({fromIdx: 0, loadAmountF: 1, amountF: 0, toIdx: 0});
    }

    function _float2Fix(uint40 floatVal) internal pure returns (uint256) {
        uint256 m = floatVal & 0x7FFFFFFFF;
        uint256 e = floatVal >> 35;

        uint256 exp = 10 ** e;
        uint256 fix = m * exp;

        return fix;
    }
}

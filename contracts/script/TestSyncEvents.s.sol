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

        vm.startBroadcast();
        // Using the Deployed Sybil contract
        Sybil sybilContract = Sybil(sybilContractAddress);

        console2.log("Sybil contract is deployed at:", address(sybilContract));

        TxParams memory params1 = validCreateAccountDeposit();
        uint256 loadAmount1 = _float2Fix(params1.loadAmountF);

        sybilContract.createAccountDeposit{value: loadAmount1}(
            params1.loadAmountF
        );

        // forgeBatch(sybilContract);

        // Test deposit
        TxParams memory params2 = validDeposit();
        uint256 loadAmount2 = _float2Fix(params2.loadAmountF);

        try
            sybilContract.deposit{value: loadAmount2}(
                params2.fromIdx,
                params2.loadAmountF
            )
        {
            console2.log("Deposit successfull");
        } catch Error(string memory reason) {
            console2.log("Deposit failed with reason:", reason);
        } catch {
            console2.log("Deposit failed with unknown reason");
        }

        // Test vouch
        TxParams memory params3 = validVouch();
        try sybilContract.vouch(params3.fromIdx, params3.toIdx) {
            console2.log("Vouch successfull");
        } catch Error(string memory reason) {
            console2.log("Vouch failed with reason:", reason);
        } catch {
            console2.log("Vouch failed with unknown reason");
        }

        // Test unvouch
        TxParams memory params4 = validUnvouch();
        try sybilContract.unvouch(params4.fromIdx, params4.toIdx) {
            console2.log("Unvouch successful");
        } catch Error(string memory reason) {
            console2.log("Unvouch failed with reason:", reason);
        } catch {
            console2.log("Unvouch failed with unknown reason");
        }

        // Test exit
        TxParams memory params5 = validExit();
        try sybilContract.exit(params5.fromIdx, params5.amountF) {
            console2.log("Exit successfull");
        } catch Error(string memory reason) {
            console2.log("Exit failed with reason:", reason);
        } catch {
            console2.log("Exit failed with unknown reason");
        }

        vm.stopBroadcast();
    }

    // function forgeBatch(Sybil sybilContract) internal {
    //     // Get current state roots
    //     uint32 lastForgedBatch = sybilContract.lastForgedBatch();
    //     uint48 currentLastIdx = sybilContract.lastIdx();
    //     console2.log("Last forged batch:", lastForgedBatch);

    //     bytes memory txData = sybilContract.getL1TransactionQueue(
    //         lastForgedBatch + 1
    //     );

    //     uint256 newAccountRoot = uint256(
    //         keccak256(abi.encodePacked(txData, "accountRoot"))
    //     );
    //     uint256 newVouchRoot = uint256(
    //         keccak256(abi.encodePacked(txData, "vouchRoot"))
    //     );
    //     uint256 newScoreRoot = uint256(
    //         keccak256(abi.encodePacked(txData, "scoreRoot"))
    //     );
    //     uint256 newExitRoot = uint256(
    //         keccak256(abi.encodePacked(txData, "exitRoot"))
    //     );

    //     console2.log(
    //         "Attempting first forgeBatch with lastIdx:",
    //         currentLastIdx + 1
    //     );

    //     uint256[2] memory proofA = [
    //         20491192805390485299153009773594534940189261866228447918068658471970481763042, // alphax
    //         9383485363053290200918347156157836566562967994039712273449902621266178545958 // alphay
    //     ];

    //     uint256[2][2] memory proofB = [
    //         [
    //             4252822878758300859123897981450591353533073413197771768651442665752259397132, // betax1
    //             21847035105528745403288232691147584728191162732299865338377159692350059136679 // betay1
    //         ],
    //         [
    //             6375614351688725206403948262868962793625744043794305715222011528459656738731, // betax2
    //             10505242626370262277552901082094356697409835680220590971873171140371331206856 // betay2
    //         ]
    //     ];

    //     uint256[2] memory proofC = [
    //         11559732032986387107991004021392285783925812861821192530917403151452391805634, // gammax1
    //         4082367875863433681332203403145435568316851327593401208105741076214120093531 // gammay1
    //     ];

    //     try
    //         sybilContract.forgeBatch(
    //             currentLastIdx + 1,
    //             newAccountRoot,
    //             newVouchRoot,
    //             newScoreRoot,
    //             newExitRoot,
    //             proofA,
    //             proofB,
    //             proofC
    //         )
    //     {
    //         console2.log("First forgeBatch successful");
    //     } catch Error(string memory reason) {
    //         console2.log("First forgeBatch failed with reason:", reason);
    //         console2.logBytes(txData);
    //     } catch {
    //         console2.log("First forgeBatch failed with unknown reason");
    //         console2.logBytes(txData);
    //     }
    // }

    function validCreateAccountDeposit() public pure returns (TxParams memory) {
        return TxParams({fromIdx: 0, loadAmountF: 2, amountF: 0, toIdx: 0});
    }

    function validDeposit() public pure returns (TxParams memory) {
        return TxParams({fromIdx: 256, loadAmountF: 100, amountF: 0, toIdx: 0});
    }

    function validVouch() public pure returns (TxParams memory) {
        return TxParams({fromIdx: 256, loadAmountF: 0, amountF: 1, toIdx: 256});
    }

    function validUnvouch() public pure returns (TxParams memory) {
        return TxParams({fromIdx: 256, loadAmountF: 0, amountF: 1, toIdx: 256});
    }

    function validExit() public pure returns (TxParams memory) {
        return TxParams({fromIdx: 256, loadAmountF: 0, amountF: 0, toIdx: 1});
    }

    function _float2Fix(uint40 floatVal) internal pure returns (uint256) {
        uint256 m = floatVal & 0x7FFFFFFFF;
        uint256 e = floatVal >> 35;

        uint256 exp = 10 ** e;
        uint256 fix = m * exp;

        return fix;
    }
}

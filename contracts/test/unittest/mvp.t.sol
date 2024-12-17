// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import "forge-std/Test.sol";
import "../../src/mvp/Sybil.sol";
import "../../src/interfaces/IMVPSybil.sol";
import "../_helpers/constants.sol";
import "../_helpers/MVPTransactionTypes.sol";
import "../../src/stub/VerifierRollupStub.sol";

contract MvpTest is Test, TransactionTypeHelper {
    Sybil public sybil;
    bytes32[] public hashes;

    function setUp() public {
        PoseidonUnit2 mockPoseidon2 = new PoseidonUnit2();
        PoseidonUnit3 mockPoseidon3 = new PoseidonUnit3();
        PoseidonUnit4 mockPoseidon4 = new PoseidonUnit4();
        emit log_address(address(mockPoseidon2));
        emit log_address(address(mockPoseidon3));
        emit log_address(address(mockPoseidon4));

        VerifierRollupStub verifierStub = new VerifierRollupStub(); 

        address[] memory verifiers = new address[](1);
        uint256[] memory maxTx = new uint256[](1);
        uint256[] memory nLevels = new uint256[](1);

        verifiers[0] = address(verifierStub);
        maxTx[0] = uint(256);
        nLevels[0] = uint(1);

        sybil = new Sybil();

        sybil.initialize(
            verifiers, 
            maxTx, 
            nLevels, 
            120, 
            address(mockPoseidon2), 
            address(mockPoseidon3), 
            address(mockPoseidon4)
        );
    }

    function testGetStateRoot() public {
        uint32 batchNum = 1;
        uint256 input = uint(1);
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];

        vm.prank(address(this));
        sybil.forgeBatch(
            256, 
            0xabc, 
            0, 
            0, 
            0, 
            0, 
            proofA,
            proofB,
            proofC,
            input
        );
        uint256 stateRoot = sybil.getStateRoot(batchNum);
        assertEq(stateRoot, 0xabc);
    }

    function testGetLastForgedBatch() public {
        uint32 lastForged = sybil.getLastForgedBatch();
        assertEq(lastForged, 0);

        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            256, 
            0xabc, 
            0, 
            0, 
            0, 
            0, 
            proofA,
            proofB,
            proofC,
            input
        );

        lastForged = sybil.getLastForgedBatch();
        assertEq(lastForged, 1);
    }

    function testGetL1TransactionQueue() public {
        TxParams memory params = validDeposit();

        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            256, 
            0xabc, 
            0, 
            0, 
            0, 
            0, 
            proofA,
            proofB,
            proofC,
            input
        );

        vm.prank(address(this));
        sybil.deposit {
            value: loadAmount
        }(params.fromIdx, params.loadAmountF);

        bytes memory txData = sybil.getL1TransactionQueue(1);
        bytes memory expectedTxData = abi.encodePacked(address(this), params.fromIdx, params.loadAmountF, params.amountF, params.toIdx);
        assertEq(txData, expectedTxData);
    }

    function testGetQueueLength() public {
        uint32 queueLength = sybil.getQueueLength();
        assertEq(queueLength, 1);

        TxParams memory params = validCreateAccountDeposit();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        vm.prank(address(this));
        sybil.createAccountDeposit {
            value: loadAmount
        }(params.loadAmountF);

        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            256, 
            0xabc, 
            0, 
            0, 
            0, 
            0, 
            proofA,
            proofB,
            proofC,
            input
        );

        queueLength = sybil.getQueueLength();
        assertEq(queueLength, 0);
    }

    function testSetForgeL1BatchTimeout() public {
        uint8 newTimeout = 255;
        vm.expectRevert(IMVPSybil.BatchTimeoutExceeded.selector);
        sybil.setForgeL1BatchTimeout(newTimeout);
    }

    function testClearQueue() public {
        TxParams memory params = validCreateAccountDeposit();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        vm.prank(address(this));
        sybil.createAccountDeposit {
            value: loadAmount
        }(params.loadAmountF);
    
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            256, 
            0xabc, 
            0, 
            0, 
            0, 
            0, 
            proofA,
            proofB,
            proofC,
            input
        );

        uint32 queueAfter = sybil.getQueueLength();
        assertEq(queueAfter, 0);

        queueAfter = sybil.getQueueLength();
        assertEq(sybil.getLastForgedBatch(),1);
        assertEq(queueAfter, 0);
    }

    // Events tests
    function testForgeBatchEventEmission() public {
        vm.expectEmit(true, true, true, true);
        emit Sybil.ForgeBatch(1, 0);

        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            256, 
            0xabc, 
            0, 
            0, 
            0, 
            0, 
            proofA,
            proofB,
            proofC,
            input
        );
    }

    function testL1UserTxEventEmission() public {
        TxParams memory params = validCreateAccountDeposit();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        vm.expectEmit(true, true, true, true);
        emit Sybil.L1UserTxEvent(1, 0, abi.encodePacked(address(this), params.fromIdx, params.loadAmountF, params.amountF, params.toIdx));

        vm.prank(address(this));
        sybil.createAccountDeposit {
            value: loadAmount
        }(params.loadAmountF);
    }

    function testInitializeEventEmission() public {
        PoseidonUnit2 mockPoseidon2 = new PoseidonUnit2();
        PoseidonUnit3 mockPoseidon3 = new PoseidonUnit3();
        PoseidonUnit4 mockPoseidon4 = new PoseidonUnit4();

        // Deploy verifier stub
        VerifierRollupStub verifierStub = new VerifierRollupStub(); 
        
        address[] memory verifiers = new address[](1);
        uint256[] memory maxTx = new uint256[](1);
        uint256[] memory nLevels = new uint256[](1);

        verifiers[0] = address(verifierStub);
        maxTx[0] = uint(256);
        nLevels[0] = uint(1);

        Sybil newSybil = new Sybil();
        newSybil.initialize(
            verifiers, 
            maxTx, 
            nLevels, 
            120, 
            address(mockPoseidon2), 
            address(mockPoseidon3), 
            address(mockPoseidon4)
        );
        emit Sybil.Initialize(120);
    }

        // CreateAccount transactions tests
    function testCreateDepositAccountTransaction() public {
        TxParams memory params = validCreateAccountDeposit();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        vm.prank(address(this));
        sybil.createAccountDeposit {
            value: loadAmount
        }(params.loadAmountF);
    }

    function testInvalidCreateDepositAccountTransaction() public {
        TxParams memory params = invalidCreateAccountDeposit();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        vm.prank(address(this));
        vm.expectRevert(IMVPSybil.LoadAmountDoesNotMatch.selector);
        sybil.createAccountDeposit {
            value: 2*loadAmount
        }(params.loadAmountF);
    }

    function testDepositTransaction() public {
        TxParams memory params = validDeposit();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        uint48 initialLastIdx = 256;
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0,
            proofA,
            proofB,
            proofC,
            input
        );

        vm.prank(address(this));
        sybil.deposit {
            value: loadAmount
        }(params.fromIdx, params.loadAmountF);
    }

    function testInvalidDepositTransactionWithLoadAmountDoesNotMatch() public {
        TxParams memory params = validDeposit();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        uint48 initialLastIdx = 256;
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0,
            proofA,
            proofB,
            proofC,
            input
        );

        vm.prank(address(this));
        vm.expectRevert(IMVPSybil.LoadAmountDoesNotMatch.selector);
        sybil.deposit {
            value: 2*loadAmount
        }(params.fromIdx, params.loadAmountF);
    }

    function testInvalidDepositTransactionWithInvalidFromIdx() public {
        TxParams memory params = invalidDeposit();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        uint48 initialLastIdx = 255;
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0,
            proofA,
            proofB,
            proofC,
            input
        );

        vm.prank(address(this));
        vm.expectRevert(IMVPSybil.InvalidFromIdx.selector);
        sybil.deposit {
            value: loadAmount
        }(params.fromIdx, params.loadAmountF);
    }

    function testVouch() public {
        TxParams memory params = validVouch();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        uint48 initialLastIdx = 256;
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));

        // forging to set the lastIdx
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0,
            proofA,
            proofB,
            proofC,
            input
        );

        vm.prank(address(this));
        sybil.vouch(params.fromIdx, params.toIdx);
    }

        function testInvalidVouchWithInvalidFromIdx() public {
        TxParams memory params = validVouch();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        uint48 initialLastIdx = 256;
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));

        // forging to set the lastIdx
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0,
            proofA,
            proofB,
            proofC,
            input
        );

        vm.prank(address(this));
        vm.expectRevert(IMVPSybil.InvalidFromIdx.selector);
        // fromIdx is 244 due to which the tx reverts
        sybil.vouch(244, params.toIdx);
    }

    function testInvalidVouchWithInvalidToIdx() public {
        TxParams memory params = validVouch();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        uint48 initialLastIdx = 256;
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));

        // forging to set the lastIdx
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0,
            proofA,
            proofB,
            proofC,
            input
        );

        vm.prank(address(this));
        vm.expectRevert(IMVPSybil.InvalidToIdx.selector);
        // toIdx is 244 due to which the tx reverts
        sybil.vouch(params.fromIdx, 244);
    }


    function testUnVouch() public {
        TxParams memory params = validUnVouch();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        uint48 initialLastIdx = 256;
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));

        // forging to set the lastIdx
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0,
            0, 
            0,
            proofA,
            proofB,
            proofC,
            input
        );
    
        vm.prank(address(this));
        sybil.unvouch(params.fromIdx, params.toIdx);
    }

        function testInvalidUnVouchWitInvalidFromIdx() public {
        TxParams memory params = validUnVouch();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        uint48 initialLastIdx = 256;
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));

        // forging to set the lastIdx
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0,
            proofA,
            proofB,
            proofC,
            input
        );

        vm.prank(address(this));
        vm.expectRevert(IMVPSybil.InvalidFromIdx.selector);
        // fromIdx is 244 due to which the tx reverts
        sybil.unvouch(244, params.toIdx);
    }

    function testInvalidUnVouchWitInvalidToIdx() public {
        TxParams memory params = validUnVouch();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);

        uint48 initialLastIdx = 256;
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));

        // forging to set the lastIdx
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0,
            proofA,
            proofB,
            proofC,
            input
        );

        vm.prank(address(this));
        vm.expectRevert(IMVPSybil.InvalidToIdx.selector);
        // toIdx is 244 due to which the tx reverts
        sybil.unvouch(params.fromIdx, 244);
    }

    // ForceExit transactions tests
    function testForceExitTransaction() public {
        TxParams memory params = validForceExit();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);
        uint48 initialLastIdx = 256;

        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0, 
            proofA,
            proofB,
            proofC,
            input
        );

        vm.prank(address(this));
        sybil.exit(params.fromIdx, params.amountF);
    }

        function testInvalidForceExitTransactionWithInvalidFromIdx() public {
        TxParams memory params = invalidForceExit();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);
        uint48 initialLastIdx = 256;

        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0, 
            proofA,
            proofB,
            proofC,
            input
        );

        vm.prank(address(this));
        vm.expectRevert(IMVPSybil.InvalidFromIdx.selector);
        sybil.exit(params.fromIdx, params.amountF);
    }

    // ForceExplode transactions tests
    function testExplodeMultiple() public {
        TxParams memory params = validForceExplode();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);
        uint48 initialLastIdx = 256;

        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0,  
            proofA,
            proofB,
            proofC,
            input
        );

        uint48 fromIdx = params.fromIdx;

        uint48[] memory toIdxs = new uint48[](1);
        toIdxs[0] = params.toIdx;
        vm.prank(address(this));
        sybil.explodeMultiple(fromIdx, toIdxs);
        
    }

    function testExplodeMultipleWithInvalidToIdx() public {
        TxParams memory params = invalidFromIdxForceExplode();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);
        uint48 initialLastIdx = 256;

        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0,  
            proofA,
            proofB,
            proofC,
            input
        );

        uint48 fromIdx = params.fromIdx;

        uint48[] memory toIdxs = new uint48[](1);
        toIdxs[0] = params.toIdx;
        vm.prank(address(this));
        vm.expectRevert(IMVPSybil.InvalidFromIdx.selector);
        sybil.explodeMultiple(fromIdx, toIdxs); 
    }

            function testExplodeMultipleWithInvalidFromIdx() public {
        TxParams memory params = invalidToIdxForceExplode();
        uint256 loadAmount = (params.loadAmountF) * 10 ** (18 - 8);
        uint48 initialLastIdx = 256;

        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        uint256 input = uint(1);

        vm.prank(address(this));
        sybil.forgeBatch(
            initialLastIdx, 
            0xabc, 
            0, 
            0, 
            0, 
            0,  
            proofA,
            proofB,
            proofC,
            input
        );

        uint48 fromIdx = params.fromIdx;

        uint48[] memory toIdxs = new uint48[](1);
        toIdxs[0] = params.toIdx;
        vm.prank(address(this));
        vm.expectRevert(IMVPSybil.InvalidToIdx.selector);
        sybil.explodeMultiple(fromIdx, toIdxs); 
    }

    function testInitializeWithInvalidPoseidonAddresses() public {
        PoseidonUnit2 mockPoseidon2 = new PoseidonUnit2();
        PoseidonUnit3 mockPoseidon3 = new PoseidonUnit3();
        PoseidonUnit4 mockPoseidon4 = new PoseidonUnit4();
        // Deploy verifier stub
        VerifierRollupStub verifierStub = new VerifierRollupStub(); 
        
        address[] memory verifiers = new address[](1);
        uint256[] memory maxTx = new uint256[](1);
        uint256[] memory nLevels = new uint256[](1);

        verifiers[0] = address(verifierStub);
        maxTx[0] = uint(256);
        nLevels[0] = uint(1);


        address invalidAddress = address(0);

        // Expect revert for invalid poseidon2Elements address
        Sybil newSybil = new Sybil();
        vm.expectRevert();
        newSybil.initialize(
            verifiers, 
            maxTx, 
            nLevels, 
            120, 
            invalidAddress, 
            address(mockPoseidon3), 
            address(mockPoseidon4)
        );

        // Expect revert for invalid poseidon3Elements address
        vm.expectRevert();
        newSybil.initialize(
            verifiers, 
            maxTx, 
            nLevels, 
            120, 
            address(mockPoseidon2), 
            invalidAddress, 
            address(mockPoseidon4)
        );

        // Expect revert for invalid poseidon4Elements address
        vm.expectRevert();
        newSybil.initialize(
            verifiers, 
            maxTx, 
            nLevels, 
            120, 
            address(mockPoseidon2), 
            address(mockPoseidon3),
            invalidAddress
        );
    }

        // Test initializing with invalid verifier address
    function testInitializeWithInvalidVerifierAddresses() public {
        PoseidonUnit2 mockPoseidon2 = new PoseidonUnit2();
        PoseidonUnit3 mockPoseidon3 = new PoseidonUnit3();
        PoseidonUnit4 mockPoseidon4 = new PoseidonUnit4();
        // Deploy verifier stub
        VerifierRollupStub verifierStub = new VerifierRollupStub(); 
        
        address[] memory verifiers = new address[](1);
        uint256[] memory maxTx = new uint256[](1);
        uint256[] memory nLevels = new uint256[](1);

        verifiers[0] = address(0);
        maxTx[0] = uint(256);
        nLevels[0] = uint(1);


        address invalidAddress = address(0);

        // Expect revert for invalid verifier address
        Sybil newSybil = new Sybil();
        vm.expectRevert(IMVPSybil.InvalidVerifierAddress.selector);
        newSybil.initialize(
            verifiers, 
            maxTx, 
            nLevels, 
            120, 
            invalidAddress, 
            address(mockPoseidon3), 
            address(mockPoseidon4)
        );
    }
    
    receive() external payable { }
}

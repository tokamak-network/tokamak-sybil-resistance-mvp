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
    
    receive() external payable { }
}

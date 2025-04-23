// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import "forge-std/Test.sol";
import "../src/NewSybil.sol";
import "../src/interfaces/INewSybil.sol";
import "./utils/Constants.sol";
import "./types/NewTransactionTypes.sol";
import "../src/Verifier.sol";
import "forge-std/console.sol";
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

    function setUp() public {
        PoseidonUnit2 mockPoseidon2 = new MockPoseidon2();
        PoseidonUnit3 mockPoseidon3 = new MockPoseidon3();
        // emit log_address(address(mockPoseidon2));
        // emit log_address(address(mockPoseidon3));

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

    function testGetStateRoot() public {
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        vm.prank(address(this));
        sybil.forgeBatch(
            0xabc,
            0, 
            0,
            proofA,
            proofB,
            proofC
        );
        uint32 batchNum = sybil.lastForgedBatch();
        uint256 stateRoot = sybil.accountRootMap(batchNum);
        assertEq(stateRoot, 0xabc);
    }

    function testGetLastForgedBatch() public {
        uint32 lastForged = sybil.lastForgedBatch();
        assertEq(lastForged, 0);

        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];

        vm.prank(address(this));
        sybil.forgeBatch(
            0xabc,
            0, 
            0,
            proofA,
            proofB,
            proofC
        );

        lastForged = sybil.lastForgedBatch();
        assertEq(lastForged, 1);
    }

    function testGetL1TransactionQueue() public {
        vm.prank(address(this));
        sybil.deposit {
            value: 1 ether
        }();

        bytes memory txData = sybil.unprocessedBatchesMap(uint32(2));
        uint256 identifer = 0;
        uint256 amount = 1 ether;
        bytes memory expectedTxData = abi.encodePacked(identifer, address(this), address(0), amount);
        assertEq(txData, expectedTxData);
    }

    function testGetQueueLength() public {
        uint32 queueLength = sybil.getQueueLength();
        assertEq(queueLength, 2);

        vm.prank(address(this));
        sybil.deposit {
            value: 1 ether
        }();

        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        
        vm.prank(address(this));
        sybil.forgeBatch(
            0xabc, 
            0, 
            0,
            proofA,
            proofB,
            proofC
        );

        queueLength = sybil.getQueueLength();
        assertEq(queueLength, 2);
    }

    function testForgeBatchEventEmission() public {
        vm.expectEmit(true, true, true, true);
        emit NewSybil.ForgeBatch(1, 0);

        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        
        vm.prank(address(this));
        sybil.forgeBatch(
            0xabc, 
            0, 
            0, 
            proofA,
            proofB,
            proofC
        );
    }

    function testL1UserTxEventEmission() public {
        vm.expectEmit(true, true, true, true);
        uint256 identifer = 0;
        uint256 amount = 1 ether;
        emit NewSybil.L1UserTxEvent(2, 0, abi.encodePacked(identifer, address(this), address(0), amount));

        vm.prank(address(this));
        sybil.deposit {
            value: 1 ether
        }();
    }
}
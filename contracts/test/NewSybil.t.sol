// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import "forge-std/Test.sol";
import "../src/NewSybil.sol";
import "../src/interfaces/INewSybil.sol";
import "./utils/Constants.sol";
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

contract MvpTest is Test {
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

    function testCreateDepositAccountTransaction() public {
        uint256 balance = sybil.balances(address(this));
        // balance zero means the account has not been created
        assertEq(balance, 0 ether);
        vm.prank(address(this));
        // account is created in this deposit function
        sybil.deposit{
            value: 1 ether
        }();
        balance = sybil.balances(address(this));
        assertEq(balance, 1 ether);
    }

    function testDepositTransaction() public {
        uint256[2] memory proofA = [uint(0),uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        // account is created in this deposit function
        vm.prank(address(this));
        sybil.deposit {
            value: 1 ether
        }();
        vm.prank(address(this));
        sybil.forgeBatch(
            0xabc, 
            0, 
            0, 
            proofA,
            proofB,
            proofC
        );
        // balance is added in this deposit function
        vm.prank(address(this));
        sybil.deposit {
            value: 1 ether
        }();
        uint256 balance = sybil.balances(address(this));
        assertEq(balance, 2 ether);
    }

    function testDepositTransactionWithLimitAmountExceeded() public {
        uint256 _LIMIT_AMOUNT = (1 << 129);
        uint num = 1 << 129;
        vm.deal(address(this), num);

        vm.prank(address(this));
        vm.expectRevert(INewSybil.LimitAmountExceeded.selector);
        sybil.deposit {
            value: _LIMIT_AMOUNT
        }();
    }

    function testDepositTransactionWithInsufficientETH() public {
        vm.prank(address(this));
        vm.expectRevert(INewSybil.InsufficientETH.selector);
        sybil.deposit();
        uint256 balance = sybil.balances(address(this));
        assertEq(balance, 0 ether);
    }

    function testVouch() public {
        vm.prank(address(this));
        sybil.deposit{
            value: 1 ether
        }();

        vm.deal(address(0x123), 1 ether);
        vm.prank(address(0x123));
        sybil.deposit{
            value: 1 ether
        }();
        vm.prank(address(this));
        sybil.vouch(address(0x123));

        assertEq(sybil.vouches(address(this), address(0x123)), true);
    }

    function testInvalidVouchWithSenderHasZeroBalance() public {
        vm.prank(address(this));
        vm.expectRevert(INewSybil.SenderHasZeroBalance.selector);
        sybil.vouch(address(0x123));
    }

    function testInvalidVouchWithReceiverHasZeroBalance() public {
        vm.prank(address(this));
        sybil.deposit{
            value: 1 ether
        }();

        vm.expectRevert(INewSybil.ReceiverHasZeroBalance.selector);
        vm.prank(address(this));
        sybil.vouch(address(0x123));
    }

    function testUnvouch() public {
        vm.prank(address(this));
        sybil.deposit{
            value: 1 ether
        }();

        vm.deal(address(0x123), 1 ether);
        vm.prank(address(0x123));
        sybil.deposit{
            value: 1 ether
        }();
        vm.prank(address(this));
        sybil.vouch(address(0x123));

        // first vouch for another address to unvouch it
        vm.prank(address(this));
        sybil.unvouch(address(0x123));

        assertEq(sybil.vouches(address(this), address(0x123)), false);
    }

}
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
        uint256[2] memory proofA = [uint(0), uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];
        vm.prank(address(this));
        sybil.forgeBatch(0xabc, 0, 0, proofA, proofB, proofC);
        uint32 batchNum = sybil.lastForgedBatch();
        uint256 stateRoot = sybil.accountRootMap(batchNum);
        assertEq(stateRoot, 0xabc);
    }

    function testGetLastForgedBatch() public {
        uint32 lastForged = sybil.lastForgedBatch();
        assertEq(lastForged, 0);

        uint256[2] memory proofA = [uint(0), uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];

        vm.prank(address(this));
        sybil.forgeBatch(0xabc, 0, 0, proofA, proofB, proofC);

        lastForged = sybil.lastForgedBatch();
        assertEq(lastForged, 1);
    }

    function testGetQueueLength() public {
        uint32 queueLength = sybil.getQueueLength();
        assertEq(queueLength, 2);

        vm.prank(address(this));
        sybil.deposit{value: 1 ether}();

        uint256[2] memory proofA = [uint(0), uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];

        vm.prank(address(this));
        sybil.forgeBatch(0xabc, 0, 0, proofA, proofB, proofC);

        queueLength = sybil.getQueueLength();
        assertEq(queueLength, 2);
    }

    function testForgeBatchEventEmission() public {
        vm.expectEmit(true, true, true, true);
        emit NewSybil.ForgeBatch(1, 0);

        uint256[2] memory proofA = [uint(0), uint(0)];
        uint256[2][2] memory proofB = [[uint(0), uint(0)], [uint(0), uint(0)]];
        uint256[2] memory proofC = [uint(0), uint(0)];

        vm.prank(address(this));
        sybil.forgeBatch(0xabc, 0, 0, proofA, proofB, proofC);
    }

    function testL1UserTxEventEmission() public {
        vm.expectEmit(true, true, true, true);
        uint8 identifier = 0;
        uint256 amount = 1 ether;
        emit NewSybil.L1UserTxEvent(
            2,
            0,
            abi.encode(identifier, address(this), address(0), amount)
        );

        vm.prank(address(this));
        sybil.deposit{value: 1 ether}();
    }

    function testCreateDepositAccountTransaction() public {
        uint256 balance = sybil.balances(address(this));
        assertEq(balance, 0 ether);

        vm.prank(address(this));
        sybil.deposit{value: 1 ether}();
        balance = sybil.balances(address(this));
        assertEq(balance, 1 ether);
    }

    function testDepositTransaction() public {
        vm.prank(address(this));
        sybil.deposit{value: 1 ether}();

        vm.prank(address(this));
        sybil.deposit{value: 1 ether}();

        uint256 balance = sybil.balances(address(this));
        assertEq(balance, 2 ether);
    }

    function testDepositTransactionWithLimitAmountExceeded() public {
        uint256 amount = (1 << 129);
        vm.deal(address(this), amount);

        vm.prank(address(this));
        vm.expectRevert(INewSybil.LimitAmountExceeded.selector);
        sybil.deposit{value: amount}();
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
        sybil.deposit{value: 1 ether}();

        vm.deal(address(0x123), 1 ether);
        vm.prank(address(0x123));
        sybil.deposit{value: 1 ether}();
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
        sybil.deposit{value: 1 ether}();

        vm.expectRevert(INewSybil.ReceiverHasZeroBalance.selector);
        vm.prank(address(this));
        sybil.vouch(address(0x123));
    }

    function testUnvouch() public {
        vm.prank(address(this));
        sybil.deposit{value: 1 ether}();

        vm.deal(address(0x123), 1 ether);
        vm.prank(address(0x123));
        sybil.deposit{value: 1 ether}();
        vm.prank(address(this));
        sybil.vouch(address(0x123));

        // first vouch for another address to unvouch it
        vm.prank(address(this));
        sybil.unvouch(address(0x123));

        assertEq(sybil.vouches(address(this), address(0x123)), false);
    }

    function testUnvouchWithNotVouched() public {
        assertEq(sybil.vouches(address(this), address(0x123)), false);

        vm.expectRevert(
            abi.encodeWithSelector(
                INewSybil.NotVouched.selector,
                address(this),
                address(0x123)
            )
        );
        vm.prank(address(this));
        sybil.unvouch(address(0x123));

        assertEq(sybil.vouches(address(this), address(0x123)), false);

    }
    function testWithdrawTransaction() public {
        vm.prank(address(this));
        sybil.deposit{value: 2 ether}();

        uint256 amount = 1 ether;
        sybil.withdraw(amount);
        assertEq(sybil.balances(address(this)), 1 ether);
    }

    function testWithdrawTransactionWithLimitAmountExceeded() public {
        uint256 amount = (1 << 129);
        vm.deal(address(this), amount);

        vm.prank(address(this));
        vm.expectRevert(INewSybil.LimitAmountExceeded.selector);
        sybil.withdraw(amount);
    }

    function testWithdrawTransactionWithInsufficientBalance() public {
        vm.prank(address(this));
        sybil.deposit{value: 1 ether}();
        vm.expectRevert(INewSybil.InsufficientBalance.selector);
        sybil.withdraw(1 ether);
    }

    function testInitializeWithInvalidPoseidonAddresses() public {
        PoseidonUnit2 mockPoseidon2 = new MockPoseidon2();
        PoseidonUnit3 mockPoseidon3 = new MockPoseidon3();
        Verifier verifierStub = new Verifier();

        address verifiers = address(verifierStub);
        uint256 maxTx = uint(256);
        uint256 nLevels = uint(1);

        address invalidAddress = address(0);

        NewSybil newSybil = new NewSybil();
        vm.expectRevert();
        newSybil.initialize(
            verifiers,
            maxTx,
            nLevels,
            invalidAddress,
            address(mockPoseidon3),
            address(this)
        );

        vm.expectRevert();
        newSybil.initialize(
            verifiers,
            maxTx,
            nLevels,
            address(mockPoseidon2),
            invalidAddress,
            address(this)
        );
    }

    function testInitializeWithInvalidVerifierAddresses() public {
        PoseidonUnit2 mockPoseidon2 = new MockPoseidon2();
        PoseidonUnit3 mockPoseidon3 = new MockPoseidon3();

        address verifier = address(0);
        uint256 maxTx = uint(256);
        uint256 nLevel = uint(1);

        NewSybil newSybil = new NewSybil();
        vm.expectRevert(INewSybil.InvalidVerifierAddress.selector);
        newSybil.initialize(
            verifier,
            maxTx,
            nLevel,
            address(mockPoseidon2),
            address(mockPoseidon3),
            address(this)
        );
    }

    function testExplodeMultiple() public {
        address[] memory addArray = new address[](4);
        addArray[0] = address(1);
        addArray[1] = address(2);
        addArray[2] = address(3);
        addArray[3] = address(4);
        address sender = address(this);

        vm.prank(sender);
        sybil.deposit{value: 2 ether}();

        for (uint256 i = 0; i < addArray.length; ++i) {
            vm.deal(addArray[i], 10 ether);
            vm.prank(addArray[i]);
            sybil.deposit{value: 2 ether}();
        }
        for (uint256 i = 0; i < addArray.length; ++i) {
            vm.prank(addArray[i]);
            sybil.vouch(sender);
        }

        vm.prank(sender);
        sybil.explodeMultiple(addArray);
        for (uint256 i = 0; i < addArray.length; ++i) {
            vm.prank(addArray[i]);
            assertEq(sybil.vouches(addArray[i], sender), false);
        }
    }

    function testExplodeMultipleWithNotVouched() public {
        address[] memory addArray = new address[](4);
        addArray[0] = address(1);
        addArray[1] = address(2);
        address sender = address(this);

        vm.prank(sender);
        vm.expectRevert(
            abi.encodeWithSelector(
                INewSybil.NotVouched.selector,
                sender,
                addArray[0]
            )
        );
        sybil.explodeMultiple(addArray);
        for (uint256 i = 0; i < addArray.length; ++i) {
            vm.prank(addArray[i]);
            assertEq(sybil.vouches(addArray[i], sender), false);
        }
    }

    function testProveScoreMerkleProof() public {
        uint32 numScoreRoot = 0;
        uint24 idx = 0;
        uint32 score = 100;
        uint256[] memory siblings = new uint256[](2);

        vm.prank(address(this));
        sybil.proveScoreMerkleProof(numScoreRoot, idx, score, siblings);
    }

    function testUpdateExplodeAmount() public {
        uint256 newExplodeAmount = 500;
        vm.prank(address(this));
        sybil.updateExplodeAmount(newExplodeAmount);

        assertEq(sybil.explodeAmount(), newExplodeAmount);
    }

    function testUpdateMinBalance() public {
        uint256 newMinBalance = 1000;
        vm.prank(address(this));
        sybil.updateMinBalance(newMinBalance);

        assertEq(sybil.minBalance(), newMinBalance);
    }

    function testUpdateExplodeAmountByNonAdmin() public {
        uint256 newExplodeAmount = 500;
        address user = address(0);

        vm.prank(user);
        vm.expectRevert();
        sybil.updateExplodeAmount(newExplodeAmount);
    }

    function testUpdateMinBalanceByNonAdmin() public {
        uint256 newMinBalance = 1000;
        address user = address(0);

        vm.prank(user);
        vm.expectRevert();
        sybil.updateMinBalance(newMinBalance);
    }

    receive() external payable {}
}

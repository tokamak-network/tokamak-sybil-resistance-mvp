pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/comparators.circom";
include "../../node_modules/circomlib/circuits/mux1.circom";
include "../../node_modules/circomlib/circuits/bitify.circom";
include "./balance-updater.circom";

template BatchTxStates(nLevels) {

    signal input MIN_BALANCE;    // Minimum balance requirement
    signal input EXPLODE_AMOUNT; // Explode amount
    
    // Inputs
    signal input txnType;       // Transaction type (0-5)
    signal input fromIdx;       // Sender index
    signal input toIdx;         // Receiver index
    signal input fromEthAddr;   // Sender ethereum address
    signal input toEthAddr;     // Receiver ethereum address
    signal input amount;        // Amount
    signal input balance2;      // Receiver balance

    // Outputs - Transaction type flags
    signal output isCreateAccount;  // Create account (txnType = 0)
    signal output isDeposit;        // Deposit (txnType = 1)
    signal output isWithdraw;       // Withdraw (txnType = 2)
    signal output isVouch;          // Vouch (txnType = 3)
    signal output isUnVouch;        // Unvouch (txnType = 4)
    signal output isExplode;        // Explode (txnType = 5)

    // Outputs - Calculated amounts
    signal output effectiveExplodeAmount;  // Actual applied amount

    // Outputs - Tree keys and functions
    signal output key1;             // Sender account key
    signal output key2;             // Receiver account key
    signal output key3;             // from->to vouch key
    signal output key4;             // to->from vouch key

    // Outputs - SMT processor function flags
    signal output P1_fnc0;
    signal output P1_fnc1;
    signal output P2_fnc0;
    signal output P2_fnc1;
    signal output P3_fnc0;
    signal output P3_fnc1;
    signal output P4_fnc0;
    signal output P4_fnc1;

    // Convert txnType to 3 bits
    component num2Bits = Num2Bits(3);
    num2Bits.in <== txnType;

    // txnType = 0 (000)
    signal temp0 <== (1 - num2Bits.out[0]) * (1 - num2Bits.out[1]);
    isCreateAccount <== temp0 * (1 - num2Bits.out[2]);

    // txnType = 1 (001)
    signal temp1 <== num2Bits.out[0] * (1 - num2Bits.out[1]);
    isDeposit <== temp1 * (1 - num2Bits.out[2]);

    // txnType = 2 (010)
    signal temp2 <== (1 - num2Bits.out[0]) * num2Bits.out[1];
    isWithdraw <== temp2 * (1 - num2Bits.out[2]);

    // txnType = 3 (011)
    signal temp3 <== num2Bits.out[0] * num2Bits.out[1];
    isVouch <== temp3 * (1 - num2Bits.out[2]);

    // txnType = 4 (100)
    signal temp4 <== (1 - num2Bits.out[0]) * (1 - num2Bits.out[1]);
    isUnVouch <== temp4 * num2Bits.out[2];

    // txnType = 5 (101)
    signal temp5 <== num2Bits.out[0] * (1 - num2Bits.out[1]);
    isExplode <== temp5 * num2Bits.out[2];

    // Check if Nop
    // nop <== 1 - (isCreateAccount + isDeposit + isWithdraw + isVouch + isUnVouch + isExplode);
    // 0 === nop * (1 - nop);

    // --- SMT Processor input setup ---
    // AccountTree keys
    key1 <== fromIdx;
    key2 <== toIdx;

    component fromIdxBits = Num2Bits(nLevels);
    component toIdxBits = Num2Bits(nLevels);
    fromIdxBits.in <== fromIdx;
    toIdxBits.in <== toIdx;

    component concatKey3 = Bits2Num(2*nLevels);
    component concatKey4 = Bits2Num(2*nLevels);

    for (var i = 0; i < nLevels; i++) {
        concatKey3.in[i] <== fromIdxBits.out[i];
        concatKey3.in[i + nLevels] <== toIdxBits.out[i];
        concatKey4.in[i] <== toIdxBits.out[i];
        concatKey4.in[i + nLevels] <== fromIdxBits.out[i];
    }

    // VouchTree keys
    key3 <== concatKey3.out;
    key4 <== concatKey4.out;

    // Processor functions
    // fnc[0]  fnc[1]
    // 0       0             NOP
    // 0       1             UPDATE
    // 1       0             INSERT
    // 1       1             DELETE

    // Setup SMT processor function flags
    P1_fnc0 <== isCreateAccount; // INSERT
    P1_fnc1 <== (1 - isCreateAccount) * (isDeposit + isWithdraw + isExplode); // INSERT if CreateAccount, otherwise UPDATE
    
    P2_fnc0 <== 0; // Never INSERT or DELETE
    P2_fnc1 <== isExplode; // UPDATE if Explode
    
    P3_fnc0 <== isVouch + isUnVouch + isExplode; // INSERT if Vouch
    P3_fnc1 <== (1 - isVouch) * (isUnVouch + isExplode); // DELETE if Unvouch or Explode
    
    P4_fnc0 <== isExplode; // DELETE if Explode
    P4_fnc1 <== isExplode; // NOP or DELETE
    
    component minAmount = LessThan(192);
    minAmount.in[0] <== EXPLODE_AMOUNT;
    minAmount.in[1] <== balance2 - MIN_BALANCE;
    
    component explodeAmountSelector = Mux1();
    explodeAmountSelector.c[0] <== EXPLODE_AMOUNT;
    explodeAmountSelector.c[1] <== balance2 - MIN_BALANCE;
    explodeAmountSelector.s <== minAmount.out;
    
    effectiveExplodeAmount <== explodeAmountSelector.out;
}
pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/comparators.circom";
include "../../node_modules/circomlib/circuits/mux1.circom";
include "../../node_modules/circomlib/circuits/bitify.circom";
include "./balance-updater.circom";

template BatchTxStates(nLevels) {
    // Inputs
    signal input txnType;       // Transaction type (0-6)
    signal input fromIdx;       // Sender index
    signal input toIdx;         // Receiver index

    // Outputs - Transaction type flags
    //signal output isNop;          // Nop (txnType = 0)
    signal output isCreateAccount;  // Create account (txnType = 1)
    signal output isDeposit;        // Deposit (txnType = 2)
    signal output isWithdraw;       // Withdraw (txnType = 3)
    signal output isVouch;          // Vouch (txnType = 4)
    signal output isUnVouch;        // Unvouch (txnType = 5)
    signal output isExplode;        // Explode (txnType = 6)

    // Outputs - Tree keys and functions
    signal output key1;             // Sender account key
    signal output key2;             // Receiver account key
    signal output key3;             // from->to vouch key

    // Outputs - SMT processor function flags
    signal output P1_fnc0;
    signal output P1_fnc1;
    signal output P2_fnc0;
    signal output P2_fnc1;
    signal output P3_fnc0;
    signal output P3_fnc1;

    // Convert txnType to 3 bits
    component num2Bits = Num2Bits(3);
    num2Bits.in <== txnType;

    // txnType = 0 (000)
    // signal temp0 <== (1 - num2Bits.out[0]) * (1 - num2Bits.out[1]);
    // isNop <== temp0 * (1 - num2Bits.out[2]);

    // txnType = 1 (001)
    signal temp1 <== num2Bits.out[0] * (1 - num2Bits.out[1]);
    isCreateAccount <== temp1 * (1 - num2Bits.out[2]);

    // txnType = 2 (010)
    signal temp2 <== (1 - num2Bits.out[0]) * num2Bits.out[1];
    isDeposit <== temp2 * (1 - num2Bits.out[2]);

    // txnType = 3 (011)
    signal temp3 <== num2Bits.out[0] * num2Bits.out[1];
    isWithdraw <== temp3 * (1 - num2Bits.out[2]);

    // txnType = 4 (100)
    signal temp4 <== (1 - num2Bits.out[0]) * (1 - num2Bits.out[1]);
    isVouch <== temp4 * num2Bits.out[2];

    // txnType = 5 (101)
    signal temp5 <== num2Bits.out[0] * (1 - num2Bits.out[1]);
    isUnVouch <== temp5 * num2Bits.out[2];

    // txnType = 6 (110)
    signal temp6 <== (1 - num2Bits.out[0]) * num2Bits.out[1];
    isExplode <== temp6 * num2Bits.out[2];

    // just one of the flags should be true
    // 1 === (isNop + isCreateAccount + isDeposit + isWithdraw + isVouch + isUnVouch + isExplode);

    // --- SMT Processor input setup ---
    // AccountTree keys
    key1 <== fromIdx;
    key2 <== toIdx;

    component fromIdxBits = Num2Bits(nLevels);
    component toIdxBits = Num2Bits(nLevels);
    fromIdxBits.in <== fromIdx;
    toIdxBits.in <== toIdx;

    // fromIdx|toIdx
    component concatKeyFromTo = Bits2Num(2*nLevels);
    for (var i = 0; i < nLevels; i++) {
        concatKeyFromTo.in[i] <== fromIdxBits.out[i];
        concatKeyFromTo.in[i + nLevels] <== toIdxBits.out[i];
    }

    // toIdx|fromIdx
    component concatKeyToFrom = Bits2Num(2*nLevels);
    for (var i = 0; i < nLevels; i++) {
        concatKeyToFrom.in[i] <== toIdxBits.out[i];
        concatKeyToFrom.in[i + nLevels] <== fromIdxBits.out[i];
    }

    component key3Mux = Mux1();
    key3Mux.c[0] <== concatKeyFromTo.out; // isExplode = 0
    key3Mux.c[1] <== concatKeyToFrom.out; // isExplode = 1
    key3Mux.s <== isExplode;

    // VouchTree keys
    key3 <== key3Mux.out;

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
}
pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/smt/smtprocessor.circom";
include "../../node_modules/circomlib/circuits/poseidon.circom";
include "../../node_modules/circomlib/circuits/gates.circom";
include "../../node_modules/circomlib/circuits/mux1.circom";

// include "./decode-tx.circom";
// include "./batch-tx.circom";
// include "./hash-inputs.circom";

template BatchMain(nTx, nLevels) {
    // Unique public signal
    signal output hashGlobalInputs;

    // private signals taking part of the hash-input
    signal input oldLastIdx;
    signal input oldAccountRoot;
    signal input oldVouchRoot;

    // Transaction inputs
    signal input txCompressedData[nTx];
    signal input fromIdx[nTx];
    signal input auxFromIdx[nTx];
    signal input toIdx[nTx];
    signal input auxToIdx[nTx];
    signal input amountF[nTx];
    signal input loadAmountF[nTx];
    signal input fromEthAddr[nTx];
    signal input toEthAddr[nTx];

    // Account State 1(from)
    signal input nonce1[nTx];
    signal input balance1[nTx];
    signal input ethAddr1[nTx]; // @TODO: check if this is not needed
    signal input siblings1[nTx][nLevels + 1];
    //signal input isOld0_1[nTx];
    //signal input oldKey1[nTx];
    //signal input oldValue1[nTx];

    // Account State 2(to)
    signal input nonce2[nTx];
    signal input balance2[nTx];
    signal input ethAddr2[nTx];
    signal input siblings2[nTx][nLevels + 1];
    //signal input isOld0_2[nTx];
    //signal input oldKey2[nTx];
    //signal input oldValue2[nTx];

    // Vouch Tree 1(from)
    signal input siblings3[nTx][2*nLevels + 1];
    // Vouch Tree 2(to)
    signal input siblings4[nTx][2*nLevels + 1];

    var i,j;

    component decodeTx[nTx];
    component batchTx[nTx];



}
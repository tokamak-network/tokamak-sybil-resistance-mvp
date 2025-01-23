pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/bitify.circom";
include "../../node_modules/circomlib/circuits/mux1.circom";

include "./lib/decode-float.circom";

template DecodeTx(nLevels){

    signal input txCompressedData;
    signal input amountF;
    //signal input toEthAddr;

    signal input fromEthAddr;
    signal input loadAmountF;

    signal input newAccount;

    signal input auxFromIdx;

    // fromEthAddr | fromIdx | loadAmountF | amountF | toIdx
    signal output L1TxFullData[160 + 48 + 40 + 40 + 48];

    signal input inIdx;
    signal output outIdx;

    // decode txCompressedData
    signal output fromIdx; // 48 0..47
    signal output toIdx; // 48 48..95
    //signal output nonce; // 40 96..135

    signal output amount;

    var i;

    // Parse txCompressedData (96 bits total: fromIdx[48], toIdx[48])
    component n2bData = Num2Bits(96);
    n2bData.in <== txCompressedData;

    // fromIdx
    component b2nFrom = Bits2Num(48);
    for (i = 0; i < 48; i++) {
        b2nFrom.in[i] <== n2bData.out[i];
    }
    b2nFrom.out ==> fromIdx;

    var paddingFrom = 0;
    for (i = nLevels; i < 48; i++) {
        paddingFrom += n2bData.out[i];
    }
    paddingFrom === 0;

    // toIdx
    component b2nTo = Bits2Num(48);
    for (i = 0; i < 48; i++) {
        b2nTo.in[i] <== n2bData.out[48 + i];
    }
    b2nTo.out ==> toIdx;

    var paddingTo = 0;
    for (i = nLevels; i < 48; i++) {
        paddingTo += n2bData.out[48 + i];
    }
    paddingTo === 0;

    // Parse amount
    component n2bAmount = Num2Bits(40);
    n2bAmount.in <== amountF;
    component dfAmount = DecodeFloatBin();
    for (i = 0; i < 40; i++) {
        dfAmount.in[i] <== n2bAmount.out[i];
    }
    dfAmount.out ==> amount;


    // Build L1TxFullData
    // Add fromEthAddr
    component n2bFromEthAddr = Num2Bits(160);
    n2bFromEthAddr.in <== fromEthAddr;
    for (i = 0; i < 160; i++) {
        L1TxFullData[160 - 1 - i] <== n2bFromEthAddr.out[i];
    }

    // Add fromIdx
    for (i = 0; i < 48; i++) {
        L1TxFullData[160 + 48 - 1 - i] <== n2bData.out[i];
    }

    // Add loadAmountF
    component n2bLoadAmountF = Num2Bits(40);
    n2bLoadAmountF.in <== loadAmountF;
    for (i = 0; i < 40; i++) {
        L1TxFullData[160 + 48 + 40 - 1 - i] <== n2bLoadAmountF.out[i];
    }

    // Add amountF
    for (i = 0; i < 40; i++) {
        L1TxFullData[160 + 48 + 40 + 40 - 1 - i] <== n2bAmount.out[i];
    }

    // Add toIdx
    for (i = 0; i < 48; i++) {
        L1TxFullData[160 + 48 + 40 + 40 + 48 - 1 - i] <== n2bData.out[48 + i];
    }

    component fromIdxIsZero = IsZero();
    fromIdxIsZero.in <== fromIdx;
    fromIdxIsZero.out === newAccount;

    outIdx <== inIdx + newAccount;

    // check auxFromIdx if it is an L1 tx and new account
    // force that index inserted for creating new accounts must be incremental
    component idxChecker = ForceEqualIfEnabled();
    idxChecker.in[0] <== auxFromIdx;
    idxChecker.in[1] <== outIdx;
    idxChecker.enabled <== newAccount;
}
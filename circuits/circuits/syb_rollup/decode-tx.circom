pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/bitify.circom";
include "../../node_modules/circomlib/circuits/comparators.circom";

template DecodeTx(nLevels) {
    // Input
    signal input txData; // (8 + nLevels + nLevels + 128) bits transaction data
    
    // Outputs
    signal output txnType;       // Transaction type (0-6, 1 byte)
    signal output fromIdx;       // Sender index (nLevels bits, nLevels/8 bytes)
    signal output toIdx;         // Receiver index (nLevels bits, nLevels/8 bytes)
    signal output amount;        // Amount (128 bits, 16 bytes)
    
    signal output bitsTxData[8 + nLevels + nLevels + 128]; //for HashGlobalInputs

    var txDataBitsLength = 8 + nLevels + nLevels + 128;
    
    // Convert txData to bits
    component txDataBits = Num2Bits(txDataBitsLength); // (8 + nLevels + nLevels + 128) bits
    txDataBits.in <== txData;

    // Extract txnType (lowest 8 bits)
    component txnType_bits2Num = Bits2Num(8);
    for (var i = 0; i < 8; i++) {
        txnType_bits2Num.in[i] <== txDataBits.out[i];
        bitsTxData[i] <== txDataBits.out[i];
    }
    txnType <== txnType_bits2Num.out;
    
    // Extract fromIdx (next nLevels bits)
    component fromIdx_bits2Num = Bits2Num(nLevels);
    for (var i = 0; i < nLevels; i++) {
        fromIdx_bits2Num.in[i] <== txDataBits.out[8 + i];
        bitsTxData[8 + i] <== txDataBits.out[8 + i];
    }
    fromIdx <== fromIdx_bits2Num.out;
    
    // Extract toIdx (next nLevels bits)
    component toIdx_bits2Num = Bits2Num(nLevels);
    for (var i = 0; i < nLevels; i++) {
        toIdx_bits2Num.in[i] <== txDataBits.out[8 + nLevels + i];
        bitsTxData[8 + nLevels + i] <== txDataBits.out[8 + nLevels + i];
    }
    toIdx <== toIdx_bits2Num.out;
    
    // Extract amount (highest 128 bits)
    component amount_bits2Num = Bits2Num(128);
    for (var i = 0; i < 128; i++) {
        amount_bits2Num.in[i] <== txDataBits.out[8 + nLevels + nLevels + i];
        bitsTxData[8 + nLevels + nLevels + i] <== txDataBits.out[8 + nLevels + nLevels + i];
    }
    amount <== amount_bits2Num.out;
    
    // Validate txnType (only values in 0-6 range are allowed)
    component isTxnTypeValid = LessThan(8);
    isTxnTypeValid.in[0] <== txnType;
    isTxnTypeValid.in[1] <== 7; // Only 0-6 are valid
    
    // txnType must be valid
    1 === isTxnTypeValid.out;
}
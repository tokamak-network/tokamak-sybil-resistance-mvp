pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/sha256/sha256.circom";
include "../../node_modules/circomlib/circuits/bitify.circom";

template HashInputs(nTx, nLevels) {
    var bitsRoots = 256;       // Root size (256 bits)
    var bitsTxsData = nTx * (8 + nLevels + nLevels + 128); // Transaction data size
    var totalBitsSha256 = 6*bitsRoots + bitsTxsData; // Total SHA256 input bits
    
    // Inputs
    signal input oldAccountRoot;   // Previous account root
    signal input oldVouchRoot;     // Previous voucher root
    signal input oldScoreRoot;     // Previous score root
    signal input newAccountRoot;   // New account root
    signal input newVouchRoot;     // New voucher root
    signal input newScoreRoot;     // New score root
    signal input txsData[bitsTxsData]; // Transaction data
    
    // Output
    signal output hashInputsOut;   // Final hash output
    
    var i;
    
    // Convert root values to bits
    component n2bOldAccountRoot = Num2Bits(256);
    n2bOldAccountRoot.in <== oldAccountRoot;
    
    component n2bOldVouchRoot = Num2Bits(256);
    n2bOldVouchRoot.in <== oldVouchRoot;

    component n2bOldScoreRoot = Num2Bits(256);
    n2bOldScoreRoot.in <== oldScoreRoot;
    
    component n2bNewAccountRoot = Num2Bits(256);
    n2bNewAccountRoot.in <== newAccountRoot;
    
    component n2bNewVouchRoot = Num2Bits(256);
    n2bNewVouchRoot.in <== newVouchRoot;
    
    component n2bNewScoreRoot = Num2Bits(256);
    n2bNewScoreRoot.in <== newScoreRoot;

    // Configure SHA256
    component inputsHasher = Sha256(totalBitsSha256);
    
    var offset = 0;
    
    // Add oldAccountRoot
    for (i = 0; i < bitsRoots; i++) {
        inputsHasher.in[offset + bitsRoots - 1 - i] <== n2bOldAccountRoot.out[i];
    }
    offset += bitsRoots;
    
    // Add oldVouchRoot
    for (i = 0; i < bitsRoots; i++) {
        inputsHasher.in[offset + bitsRoots - 1 - i] <== n2bOldVouchRoot.out[i];
    }
    offset += bitsRoots;

    // Add oldScoreRoot
    for (i = 0; i < bitsRoots; i++) {
        inputsHasher.in[offset + bitsRoots - 1 - i] <== n2bOldScoreRoot.out[i];
    }
    offset += bitsRoots;
    
    // Add newAccountRoot
    for (i = 0; i < bitsRoots; i++) {
        inputsHasher.in[offset + bitsRoots - 1 - i] <== n2bNewAccountRoot.out[i];
    }
    offset += bitsRoots;
    
    // Add newVouchRoot
    for (i = 0; i < bitsRoots; i++) {
        inputsHasher.in[offset + bitsRoots - 1 - i] <== n2bNewVouchRoot.out[i];
    }
    offset += bitsRoots;

    // Add newScoreRoot
    for (i = 0; i < bitsRoots; i++) {
        inputsHasher.in[offset + bitsRoots - 1 - i] <== n2bNewScoreRoot.out[i];
    }
    offset += bitsRoots;
    
    // Add txsData (data in format: txnType, fromIdx, toIdx, amount)
    for (i = 0; i < bitsTxsData; i++) {
        inputsHasher.in[offset + i] <== txsData[i];
    }
    
    // Convert hash result to number
    component n2bHashInputsOut = Bits2Num(256);
    for (i = 0; i < 256; i++) {
        n2bHashInputsOut.in[i] <== inputsHasher.out[255 - i];
    }
    
    hashInputsOut <== n2bHashInputsOut.out;
}
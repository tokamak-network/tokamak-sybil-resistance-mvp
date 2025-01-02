pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/poseidon.circom";
include "../../node_modules/circomlib/circuits/sha256/sha256.circom";
include "../../node_modules/circomlib/circuits/bitify.circom";

template HashInputs(nLevels, nTx) {

    var bitsIndexMax = 48; // MAX_NLEVELS
    //var bitsIndex = nLevels;
    var bitsRoots = 256;
    var bitsL1TxsData = nTx * (2*bitsIndexMax + 40 + 40 + 160);//nTx must be maxL1Tx

    signal input oldLastIdx;
    signal input newLastIdx;
    signal input oldAccountRoot;
    signal input oldVouchRoot;
    signal input newAccountRoot;
    signal input newVouchRoot;
    signal input newExitRoot;
    //signal input oldScoreRoot;
    //signal input newScoreRoot;
    signal input L1TxsData[bitsL1TxsData];

    //output
    signal output hashInputsOut;

    var i;

    // get bits from all inputs
    ////////
    // oldLastIdx
    component n2bOldLastIdx = Num2Bits(48);
    n2bOldLastIdx.in <== oldLastIdx;

    var paddingOldLastIdx = 0;
    for (i = nLevels; i < 48; i++) {
        paddingOldLastIdx += n2bOldLastIdx.out[i];
    }
    paddingOldLastIdx === 0;

    // newLastIdx
    component n2bNewLastIdx = Num2Bits(48);
    n2bNewLastIdx.in <== newLastIdx;

    var paddingNewLastIdx = 0;
    for (i = nLevels; i < 48; i++) {
        paddingNewLastIdx += n2bNewLastIdx.out[i];
    }
    paddingNewLastIdx === 0;

    // oldAccountRoot
    component n2bOldAccountRoot = Num2Bits(256);
    n2bOldAccountRoot.in <== oldAccountRoot;

    // oldVouchRoot
    component n2bOldVouchRoot = Num2Bits(256);
    n2bOldVouchRoot.in <== oldVouchRoot;

    // newAccountRoot
    component n2bNewAccountRoot = Num2Bits(256);
    n2bNewAccountRoot.in <== newAccountRoot;

    // newVouchRoot
    component n2bNewVouchRoot = Num2Bits(256);
    n2bNewVouchRoot.in <== newVouchRoot;

    // newExitRoot
    component n2bNewExitRoot = Num2Bits(256);
    n2bNewExitRoot.in <== newExitRoot;

    // build SHA256 with all inputs
    ////////
    var totalBitsSha256 = 2*bitsIndexMax + 5*bitsRoots + bitsL1TxsData;

    component inputsHasher = Sha256(totalBitsSha256);

    var offset = 0;

    // add oldLastIdx
    for (i = 0; i < bitsIndexMax; i++) {
        inputsHasher.in[bitsIndexMax - 1 - i] <== n2bOldLastIdx.out[i];
    }
    offset = offset + bitsIndexMax;

    // add newLastIdx
    for (i = 0; i < bitsIndexMax; i++) {
        inputsHasher.in[offset + bitsIndexMax - 1 - i] <== n2bNewLastIdx.out[i];
    }
    offset = offset + bitsIndexMax;

    // add oldAccountRoot
    for (i = 0; i < bitsRoots; i++) {
        inputsHasher.in[offset + bitsRoots - 1 - i] <== n2bOldAccountRoot.out[i];
    }
    offset = offset + bitsRoots;

    // add oldVouchRoot
    for (i = 0; i < bitsRoots; i++) {
        inputsHasher.in[offset + bitsRoots - 1 - i] <== n2bOldVouchRoot.out[i];
    }
    offset = offset + bitsRoots;

    // add newAccountRoot
    for (i = 0; i < bitsRoots; i++) {
        inputsHasher.in[offset + bitsRoots - 1 - i] <== n2bNewAccountRoot.out[i];
    }
    offset = offset + bitsRoots;

    // add newVouchRoot
    for (i = 0; i < bitsRoots; i++) {
        inputsHasher.in[offset + bitsRoots - 1 - i] <== n2bNewVouchRoot.out[i];
    }
    offset = offset + bitsRoots;

    // add newExitRoot
    for (i = 0; i < bitsRoots; i++) {
        inputsHasher.in[offset + bitsRoots - 1 - i] <== n2bNewExitRoot.out[i];
    }
    offset = offset + bitsRoots;

    // add L1TxsData
    for (i = 0; i < bitsL1TxsData; i++) {
        inputsHasher.in[offset + i] <== L1TxsData[i];
    }

    component n2bHashInputsOut = Bits2Num(256);
    for (i = 0; i < 256; i++) {
        n2bHashInputsOut.in[i] <== inputsHasher.out[255 - i];
    }

    hashInputsOut <== n2bHashInputsOut.out;
}
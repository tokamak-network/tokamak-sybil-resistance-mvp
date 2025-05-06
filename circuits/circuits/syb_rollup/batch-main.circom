pragma circom 2.0.0;

include "./decode-tx.circom";
include "./batch-tx.circom";
// include "./hash-inputs.circom";

template BatchMain(nTx, nLevels) {
    // Public output signals
    //signal output hashGlobalInputs;

    // Public system parameters set in the contract
    signal input MIN_BALANCE;
    signal input EXPLODE_AMOUNT;

    // Private signals that participate in hash inputs
    signal input oldLastIdx;
    signal input oldAccountRoot;
    signal input oldVouchRoot;

    // txnType | fromIdx | toIdx | amount
    signal input txData[nTx];

    // Account state 1 (sender)
    signal input balance1[nTx];
    signal input ethAddr1[nTx];
    signal input siblings1[nTx][nLevels + 1];
    signal input isOld0_1[nTx];
    signal input oldKey1[nTx];
    signal input oldValue1[nTx];

    // Account state 2 (receiver)
    signal input balance2[nTx];
    signal input ethAddr2[nTx];
    signal input siblings2[nTx][nLevels + 1];
    signal input isOld0_2[nTx];
    signal input oldKey2[nTx];
    signal input oldValue2[nTx];

    // Vouch Tree 1 (from->to)
    signal input siblings3[nTx][2*nLevels + 1];
    signal input isOld0_3[nTx];
    signal input oldKey3[nTx];
    signal input oldValue3[nTx];
    
    // Vouch Tree 2 (to->from)
    signal input siblings4[nTx][2*nLevels + 1];
    signal input isOld0_4[nTx];
    signal input oldKey4[nTx];
    signal input oldValue4[nTx];

    var i, j;

    component decodeTx[nTx];
    component batchTx[nTx];

    // A - Check binary signals
    for (i = 0; i < nTx; i++) {
        isOld0_1[i] * (isOld0_1[i] - 1) === 0;
        isOld0_2[i] * (isOld0_2[i] - 1) === 0;
        isOld0_3[i] * (isOld0_3[i] - 1) === 0;
        isOld0_4[i] * (isOld0_4[i] - 1) === 0;
    }

    // B - Decode transactions
    for (i = 0; i < nTx; i++) {
        decodeTx[i] = DecodeTx(nLevels);
        decodeTx[i].txData <== txData[i];
    }

    // C - Process batch transactions
    for (i = 0; i < nTx; i++) {
        batchTx[i] = BatchTx(nLevels);
        batchTx[i].fromIdx <== decodeTx[i].fromIdx;
        batchTx[i].toIdx <== decodeTx[i].toIdx;
        batchTx[i].amount <== decodeTx[i].amount;
        batchTx[i].MIN_BALANCE <== MIN_BALANCE;
        batchTx[i].EXPLODE_AMOUNT <== EXPLODE_AMOUNT;
        batchTx[i].txnType <== decodeTx[i].txnType;

        // Sender state
        batchTx[i].balance1 <== balance1[i];
        batchTx[i].fromEthAddr <== ethAddr1[i];
        for (j = 0; j < nLevels + 1; j++) {
            batchTx[i].siblings1[j] <== siblings1[i][j];
        }
        batchTx[i].isOld0_1 <== isOld0_1[i];
        batchTx[i].oldKey1 <== oldKey1[i];
        batchTx[i].oldValue1 <== oldValue1[i];

        // Receiver state
        batchTx[i].balance2 <== balance2[i];
        batchTx[i].toEthAddr <== ethAddr2[i];
        for (j = 0; j < nLevels + 1; j++) {
            batchTx[i].siblings2[j] <== siblings2[i][j];
        }
        batchTx[i].isOld0_2 <== isOld0_2[i];
        batchTx[i].oldKey2 <== oldKey2[i];
        batchTx[i].oldValue2 <== oldValue2[i];

        // Vouch state
        for (j = 0; j < 2*nLevels + 1; j++) {
            batchTx[i].siblings3[j] <== siblings3[i][j];
            batchTx[i].siblings4[j] <== siblings4[i][j];
        }
        batchTx[i].isOld0_3 <== isOld0_3[i];
        batchTx[i].oldKey3 <== oldKey3[i];
        batchTx[i].oldValue3 <== oldValue3[i];
        batchTx[i].isOld0_4 <== isOld0_4[i];
        batchTx[i].oldKey4 <== oldKey4[i];
        batchTx[i].oldValue4 <== oldValue4[i];

        // Roots
        if (i == 0) {
            batchTx[i].oldAccountRoot <== oldAccountRoot;
            batchTx[i].oldVouchRoot <== oldVouchRoot;
        } else {
            batchTx[i].oldAccountRoot <== batchTx[i-1].newAccountRoot;
            batchTx[i].oldVouchRoot <== batchTx[i-1].newVouchRoot;
        }
    }

    // // D - Calculate global inputs hash
    // component hasherInputs = HashInputs(nLevels, nTx);

    // hasherInputs.oldAccountRoot <== oldAccountRoot;
    // hasherInputs.oldVouchRoot <== oldVouchRoot;
    // hasherInputs.newAccountRoot <== batchTx[nTx-1].newAccountRoot;
    // hasherInputs.newVouchRoot <== batchTx[nTx-1].newVouchRoot;

    // // Set L1 transaction data
    // var txDataBits = (23 * 8); // txnType[8], fromIdx[24], toIdx[24], amount[128]
    // for (i = 0; i < nTx; i++) {
    //     for (j = 0; j < txDataBits; j++) {
    //         hasherInputs.TxsData[i*txDataBits + j] <== decodeTx[i].bitsTxData[j];
    //     }
    // }

    // hashGlobalInputs <== hasherInputs.hashInputsOut;

    // Used temporarily instead of hashGlobalInputs
    signal output newAccountRoot;
    signal output newVouchRoot;

    newAccountRoot <== batchTx[nTx-1].newAccountRoot;
    newVouchRoot <== batchTx[nTx-1].newVouchRoot;
}

component main = BatchMain(1, 24);
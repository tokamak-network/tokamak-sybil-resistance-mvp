pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/gates.circom";
include "../../node_modules/circomlib/circuits/mux1.circom";

include "./decode-tx.circom";
include "./batch-tx.circom";
include "./hash-inputs.circom";

template BatchMain(nTx, nLevels) {
    // public signal
    signal output hashGlobalInputs;

    // public system parameter set in the contract
    signal input EXPLODE_AMOUNT;

    // private signals taking part of the hash-input
    signal input oldLastIdx;
    signal input oldAccountRoot;
    signal input oldVouchRoot;

    // Transaction inputs
    signal input txCompressedData[nTx];
    signal input fromIdx[nTx];
    signal input auxFromIdx[nTx];
    signal input toIdx[nTx];
    signal input amountF[nTx];
    signal input loadAmountF[nTx];
    signal input fromEthAddr[nTx];
    signal input toEthAddr[nTx];
    signal input newAccount[nTx];

    // Account State 1(from)
    signal input balance1[nTx];
    signal input ethAddr1[nTx];
    signal input siblings1[nTx][nLevels + 1];
    signal input isOld0_1[nTx];
    signal input oldKey1[nTx];
    signal input oldValue1[nTx];

    // Account State 2(to)
    signal input balance2[nTx];
    signal input ethAddr2[nTx];
    signal input siblings2[nTx][nLevels + 1];
    signal input newExit[nTx];
    signal input isOld0_2[nTx];
    signal input oldKey2[nTx];
    signal input oldValue2[nTx];

    // Vouch Tree 1(from)
    signal input siblings3[nTx][2*nLevels + 1];
    signal input isOld0_3[nTx];
    signal input oldKey3[nTx];
    signal input oldValue3[nTx];
    
    // Vouch Tree 2(to)
    signal input siblings4[nTx][2*nLevels + 1];
    signal input isOld0_4[nTx];
    signal input oldKey4[nTx];
    signal input oldValue4[nTx];

    var i,j;

    component decodeTx[nTx];
    component batchTx[nTx];

    // A - Check binary signals
    for (i = 0; i < nTx; i++) {
        newAccount[i] * (newAccount[i] - 1) === 0;
        isOld0_1[i] * (isOld0_1[i] - 1) === 0;
        isOld0_2[i] * (isOld0_2[i] - 1) === 0;
        isOld0_3[i] * (isOld0_3[i] - 1) === 0;
        isOld0_4[i] * (isOld0_4[i] - 1) === 0;
    }

    // B - decode-tx : decode transactions
    for (i = 0; i < nTx; i++) {
        decodeTx[i] = DecodeTx(nLevels);

        if (i == 0) {
            decodeTx[i].inIdx <== oldLastIdx;
        } else {
            decodeTx[i].inIdx <== decodeTx[i-1].outIdx;
        }

        decodeTx[i].txCompressedData <== txCompressedData[i];
        decodeTx[i].amountF <== amountF[i];
        decodeTx[i].fromEthAddr <== fromEthAddr[i];
        decodeTx[i].loadAmountF <== loadAmountF[i];
        decodeTx[i].newAccount <== newAccount[i];
        decodeTx[i].auxFromIdx <== auxFromIdx[i];
    }

    //C - batch-tx : process batch transactions
    for (i = 0; i < nTx; i++) {
        batchTx[i] = BatchTx(nLevels);

        batchTx[i].fromIdx <== fromIdx[i];
        batchTx[i].toIdx <== toIdx[i];
        batchTx[i].auxFromIdx <== auxFromIdx[i];
        batchTx[i].amount <== decodeTx[i].amount;
        batchTx[i].loadAmountF <== loadAmountF[i];
        batchTx[i].newAccount <== newAccount[i];
        batchTx[i].newExit <== newExit[i];
        batchTx[i].fromEthAddr <== fromEthAddr[i];
        batchTx[i].EXPLODE_AMOUNT <== EXPLODE_AMOUNT;

        // State 1
        batchTx[i].balance1 <== balance1[i];
        batchTx[i].ethAddr1 <== ethAddr1[i];
        for (j = 0; j < nLevels + 1; j++) {
            batchTx[i].siblings1[j] <== siblings1[i][j];
        }
        batchTx[i].isOld0_1 <== isOld0_1[i];
        batchTx[i].oldKey1 <== oldKey1[i];
        batchTx[i].oldValue1 <== oldValue1[i];

        // State 2
        batchTx[i].balance2 <== balance2[i];
        batchTx[i].ethAddr2 <== ethAddr2[i];
        for (j = 0; j < nLevels + 1; j++) {
            batchTx[i].siblings2[j] <== siblings2[i][j];
        }
        batchTx[i].isOld0_2 <== isOld0_2[i];
        batchTx[i].oldKey2 <== oldKey2[i];
        batchTx[i].oldValue2 <== oldValue2[i];

        // Vouch states
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
            batchTx[i].oldExitRoot <== 0;
            batchTx[i].oldVouchRoot <== oldVouchRoot;
        } else {
            batchTx[i].oldAccountRoot <== batchTx[i-1].newAccountRoot;
            batchTx[i].oldExitRoot <== batchTx[i-1].newExitRoot;
            batchTx[i].oldVouchRoot <== batchTx[i-1].newVouchRoot;
        }
    }

    //D - hash-inputs : compute global hash input
    component hasherInputs = HashInputs(nLevels, nTx);

    hasherInputs.oldLastIdx <== oldLastIdx;
    hasherInputs.newLastIdx <== decodeTx[nTx-1].outIdx;
    hasherInputs.oldAccountRoot <== oldAccountRoot;
    hasherInputs.oldVouchRoot <== oldVouchRoot;
    hasherInputs.newAccountRoot <== batchTx[nTx-1].newAccountRoot;
    hasherInputs.newVouchRoot <== batchTx[nTx-1].newVouchRoot;
    hasherInputs.newExitRoot <== batchTx[nTx-1].newExitRoot;

    // Set L1 transactions data
    var bitsL1TxData = (160 + 48 + 40 + 40 + 48); // fromEthAddr[160], fromIdx[48], loadAmountF[40], amountF[40], toIdx[48]
    for (i = 0; i < nTx; i++) {
        for (j = 0; j < bitsL1TxData; j++) {
            hasherInputs.L1TxsData[i*bitsL1TxData + j] <== decodeTx[i].L1TxFullData[j];
        }
    }

    hashGlobalInputs <== hasherInputs.hashInputsOut;
}
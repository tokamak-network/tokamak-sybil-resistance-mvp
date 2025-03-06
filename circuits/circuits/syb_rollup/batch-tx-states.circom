pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/comparators.circom";
include "../../node_modules/circomlib/circuits/mux2.circom";
include "../../node_modules/circomlib/circuits/mux1.circom";

template BatchTxStates() {
    // Inputs
    signal input fromIdx;          // 48 bits
    signal input toIdx;            // 48 bits
    signal input fromEthAddr;      // 160 bits
    //signal input toEthAddr;        // 160 bits
    signal input ethAddr1;         // 160 bits
    signal input auxFromIdx;       // 48 bits
    signal input amount;           // 192 bits
    signal input loadAmount;       // 192 bits
    signal input newExit;          // bool
    signal input newAccount;       // bool
    signal input balance;          // 192 bits
    signal input EXPLODE_AMOUNT;   // Penalty amount set in the system. It is set in the contract.
    //signal input isExplode;        // bool

    // Outputs
    signal output isP1Insert;      // bool
    signal output isP2Insert;      // bool
    signal output key1;            // 48 bits (Account tree)
    signal output key2;            // 48 bits (Account tree)
    signal output key3;            // 96 bits (Vouch tree: fromIdx|toIdx)
    signal output key4;            // 96 bits (Vouch tree: toIdx|fromIdx)
    signal output P1_fnc0;         // bool
    signal output P1_fnc1;         // bool
    signal output P2_fnc0;         // bool
    signal output P2_fnc1;         // bool
    signal output P3_fnc0;         // bool
    signal output P3_fnc1;         // bool
    signal output P4_fnc0;         // bool
    signal output P4_fnc1;         // bool
    signal output isExit;          // bool
    signal output nop;             // bool
    signal output nullifyLoadAmount;// bool
    signal output nullifyAmount;   // bool
    signal output effectiveAmount; // 192 bits

    // Select finalFromIdx
    signal finalFromIdx;
    component selectFromIdx = Mux1();
    selectFromIdx.c[0] <== fromIdx;
    selectFromIdx.c[1] <== auxFromIdx;
    selectFromIdx.s <== newAccount;
    finalFromIdx <== selectFromIdx.out;

    // Check if finalFromIdx is 0 (NOP check)
    component finalFromIdxIsZero = IsZero();
    finalFromIdxIsZero.in <== finalFromIdx;
    signal isFinalFromIdx;
    isFinalFromIdx <== 1 - finalFromIdxIsZero.out;
    nop <== finalFromIdxIsZero.out;

    var EXIT_IDX = 1;
    // Check if tx is an exit
    component checkIsExit = IsEqual();
    checkIsExit.in[0] <== EXIT_IDX;
    checkIsExit.in[1] <== toIdx;
    isExit <== checkIsExit.out;

    component toIdxIsZero = IsZero();
    toIdxIsZero.in <== toIdx;
    signal isToIdx;
    isToIdx <== 1 - toIdxIsZero.out;

    // Check if amount/loadAmount is non-zero
    component amountIsZero = IsZero();
    amountIsZero.in <== amount;
    signal isAmount;
    isAmount <== 1 - amountIsZero.out;

    component loadAmountIsZero = IsZero();
    loadAmountIsZero.in <== loadAmount;
    signal isLoadAmount;
    isLoadAmount <== 1 - loadAmountIsZero.out;

    // Identify transaction types based on amount value
    component amountIsOne = IsEqual();
    amountIsOne.in[0] <== amount;
    amountIsOne.in[1] <== 1;

    component amountIsTwo = IsEqual();
    amountIsTwo.in[0] <== amount;
    amountIsTwo.in[1] <== 2;

    // VouchTx: finalFromIdx != 0 && toIdx != 0 && amount == 1 && !isExit && !nop
    signal vouchTemp1 <== isFinalFromIdx * isToIdx;
    signal vouchTemp2 <== vouchTemp1 * amountIsOne.out;
    signal vouchTemp3 <== vouchTemp2 * (1 - isExit);
    signal output isVouchTx <== vouchTemp3 * (1 - nop);

    // DeleteVouchTx: finalFromIdx != 0 && toIdx != 0 && amount == 0 && !isExit && !nop
    signal deleteTemp1 <== isFinalFromIdx * isToIdx;
    signal deleteTemp2 <== deleteTemp1 * (1 - isAmount);
    signal deleteTemp3 <== deleteTemp2 * (1 - isExit);
    signal isDeleteVouchTx <== deleteTemp3 * (1 - nop);

    // Transfer: finalFromIdx != 0 && toIdx != 0 && amount == 2 && !isExit && !nop
    signal transferTemp1 <== isFinalFromIdx * isToIdx;
    signal transferTemp2 <== transferTemp1 * amountIsTwo.out;
    signal transferTemp3 <== transferTemp2 * (1 - isExit);
    signal isTransfer <== transferTemp3 * (1 - nop);


    // Account tree keys (P1, P2)
    key1 <== finalFromIdx;
    key2 <== toIdx;

    // Vouch tree keys (P3, P4)
    component fromIdxBits = Num2Bits(48);
    component toIdxBits = Num2Bits(48);
    fromIdxBits.in <== finalFromIdx;
    toIdxBits.in <== toIdx;

    component concatKey3 = Bits2Num(96);
    component concatKey4 = Bits2Num(96);
    
    for (var i = 0; i < 48; i++) {
        concatKey3.in[i] <== fromIdxBits.out[i];
        concatKey3.in[i + 48] <== toIdxBits.out[i];
        concatKey4.in[i] <== toIdxBits.out[i];
        concatKey4.in[i + 48] <== fromIdxBits.out[i];
    }
    
    key3 <== concatKey3.out;
    key4 <== concatKey4.out;

    // Processor functions
    // fnc[0]  fnc[1]
    // 0       0             NOP
    // 0       1             UPDATE
    // 1       0             INSERT
    // 1       1             DELETE

    // Account tree (P1)
    isP1Insert <== newAccount;
    P1_fnc0 <== isP1Insert * isFinalFromIdx;
    P1_fnc1 <== (1 - isP1Insert) * isFinalFromIdx;

    // Account/Exit tree (P2)
    isP2Insert <== isExit * newExit;
    P2_fnc0 <== isP2Insert * isFinalFromIdx;
    P2_fnc1 <== (1 - isP2Insert) * isFinalFromIdx;

    log(isFinalFromIdx);
    log((isVouchTx + isDeleteVouchTx) * (1 - nop));
    log(isDeleteVouchTx);

    // Vouch tree (P3, P4)
    // P3: handles fromIdx|toIdx vouch
    P3_fnc0 <== 0; // Never INSERT
    P3_fnc1 <== (isVouchTx + isDeleteVouchTx) * (1 - nop); // Always UPDATE (1->Vouch, 0->unVouch)
    // P4: handles toIdx|fromIdx vouch
    P4_fnc0 <== 0; // Never INSERT
    P4_fnc1 <== isDeleteVouchTx * (1 - nop); // UPDATE (delete for deleteVouch/transfer)

    // Amount processing for transfer
    component minAmount = LessThan(192);
    minAmount.in[0] <== balance;
    minAmount.in[1] <== EXPLODE_AMOUNT;

    component amountSelector = Mux1();
    amountSelector.c[0] <== EXPLODE_AMOUNT;
    amountSelector.c[1] <== balance;
    amountSelector.s <== minAmount.out;

    component effectiveAmountSelector = Mux1();
    effectiveAmountSelector.c[0] <== amount;
    effectiveAmountSelector.c[1] <== amountSelector.out;
    effectiveAmountSelector.s <== isTransfer;

    effectiveAmount <== effectiveAmountSelector.out;

    // Nullifier logic
    signal shouldCheckEthAddr;
    // Check Ethereum address only if amount is exists and not a new account
    shouldCheckEthAddr <== (1 - newAccount) * isAmount;

    // Check that the transaction's fromEthAddr matches the real account's ethAddr1
    component checkFromEthAddr = IsEqual();
    checkFromEthAddr.in[0] <== fromEthAddr;
    checkFromEthAddr.in[1] <== ethAddr1;

    // Apply a nullifier if the Ethereum address doesn't match
    signal applyNullifier;
    applyNullifier <== shouldCheckEthAddr * (1 - checkFromEthAddr.out);

    // Invalidate amount or loadAmount
    nullifyAmount <== applyNullifier * isAmount;
    nullifyLoadAmount <== applyNullifier * isLoadAmount;
}
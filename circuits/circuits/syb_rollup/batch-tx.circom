pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/smt/smtprocessor.circom";
include "../../node_modules/circomlib/circuits/gates.circom";
include "../../node_modules/circomlib/circuits/mux1.circom";
include "../../node_modules/circomlib/circuits/bitify.circom";

include "./balance-updater.circom";
include "./batch-tx-states.circom";
include "./lib/hash-state.circom";
include "./lib/decode-float.circom";

template BatchTx(nLevels) {
    // Transaction inputs
    signal input fromIdx;
    signal input toIdx;
    signal input auxFromIdx;
    //signal input toEthAddr;
    signal input amount;
    signal input loadAmountF;
    signal input newAccount;
    signal input newExit;
    signal input fromEthAddr;
    signal input EXPLODE_AMOUNT;   // For batch-tx-states

    // State 1 (sender)
    signal input balance1;
    signal input ethAddr1;
    signal input siblings1[nLevels + 1];
    signal input isOld0_1;
    signal input oldKey1;
    signal input oldValue1;

    // State 2 (receiver)
    signal input balance2;
    signal input ethAddr2;
    signal input siblings2[nLevels + 1];
    signal input isOld0_2;
    signal input oldKey2;
    signal input oldValue2;

    // Vouch states
    signal input siblings3[2*nLevels + 1];
    signal input isOld0_3;
    signal input oldKey3;
    signal input oldValue3;
    signal input siblings4[2*nLevels + 1];
    signal input isOld0_4;
    signal input oldKey4;
    signal input oldValue4;

    // Roots
    signal input oldAccountRoot;
    signal input oldExitRoot;
    signal input oldVouchRoot;

    // Outputs
    signal output isAmountNullified;
    signal output newAccountRoot;
    signal output newExitRoot;
    signal output newVouchRoot;

    var i;

    // A - Compute Transaction States
    // Decode loadAmountF (float40)
    signal loadAmount;

    component n2bloadAmountF = Num2Bits(40);
    n2bloadAmountF.in <== loadAmountF;

    component dfLoadAmount = DecodeFloatBin();
    for (i = 0; i < 40; i++) {
        dfLoadAmount.in[i] <== n2bloadAmountF.out[i];
    }

    loadAmount <== dfLoadAmount.out;

    // Compute states
    component states = BatchTxStates();
    states.fromIdx <== fromIdx;
    states.toIdx <== toIdx;
    states.fromEthAddr <== fromEthAddr;
    states.ethAddr1 <== ethAddr1;
    states.auxFromIdx <== auxFromIdx;
    states.amount <== amount;
    states.loadAmount <== loadAmount;
    states.newExit <== newExit;
    states.newAccount <== newAccount;
    states.balance <== balance1;
    states.EXPLODE_AMOUNT <== EXPLODE_AMOUNT;

    // B - Compute Old State Hashes
    component oldSt1Hash = HashState();
    oldSt1Hash.balance <== balance1;
    oldSt1Hash.ethAddr <== ethAddr1;

    component oldSt2Hash = HashState();
    oldSt2Hash.balance <== balance2;
    oldSt2Hash.ethAddr <== ethAddr2;

    // C - Update Balances
    component balanceUpdater = BalanceUpdater();
    balanceUpdater.oldStBalanceSender <== balance1;
    balanceUpdater.oldStBalanceReceiver <== balance2;
    balanceUpdater.amount <== states.effectiveAmount;
    balanceUpdater.loadAmount <== loadAmount;
    balanceUpdater.nullifyLoadAmount <== states.nullifyLoadAmount;
    balanceUpdater.nullifyAmount <== states.nullifyAmount;
    balanceUpdater.nop <== states.nop;

    isAmountNullified <== states.nullifyAmount + states.nullifyLoadAmount;

    // D - Compute New State Hashes
    component newSt1Hash = HashState();
    newSt1Hash.balance <== balanceUpdater.newStBalanceSender;
    newSt1Hash.ethAddr <== ethAddr1;

    component newSt2Hash = HashState();
    newSt2Hash.balance <== balanceUpdater.newStBalanceReceiver;
    newSt2Hash.ethAddr <== ethAddr2;

    // E - SMT Processors
    // Account tree processor 1 (sender)
    component processor1 = SMTProcessor(nLevels + 1);
    processor1.oldRoot <== oldAccountRoot;
    for (i = 0; i < nLevels + 1; i++) {
        processor1.siblings[i] <== siblings1[i];
    }
    processor1.oldKey <== oldKey1;
    processor1.oldValue <== oldValue1;
    processor1.isOld0 <== isOld0_1;
    processor1.newKey <== states.key1;
    processor1.newValue <== newSt1Hash.out;
    processor1.fnc[0] <== states.P1_fnc0;
    processor1.fnc[1] <== states.P1_fnc1;

    // Select processor 2 root input
    // If tx is an 'Exit' select 'oldExitRoot', otherwise
    // select output root of processor 1 (account root)
    component selectP2Root = Mux1();
    selectP2Root.c[0] <== processor1.newRoot;
    selectP2Root.c[1] <== oldExitRoot;
    selectP2Root.s <== states.isExit;

    // Account/Exit tree processor 2 (receiver)
    component processor2 = SMTProcessor(nLevels + 1);
    processor2.oldRoot <== selectP2Root.out;
    for (i = 0; i < nLevels + 1; i++) {
        processor2.siblings[i] <== siblings2[i];
    }
    processor2.oldKey <== oldKey2;
    processor2.oldValue <== oldValue2;
    processor2.isOld0 <== isOld0_2;
    processor2.newKey <== states.key2;
    processor2.newValue <== newSt2Hash.out;
    processor2.fnc[0] <== states.P2_fnc0 * balanceUpdater.isP2Nop;
    processor2.fnc[1] <== states.P2_fnc1 * balanceUpdater.isP2Nop;

    // Vouch tree processor 3 (fromIdx|toIdx)
    component processor3 = SMTProcessor(2*nLevels + 1);
    processor3.oldRoot <== oldVouchRoot;
    for (i = 0; i < 2*nLevels + 1; i++) {
        processor3.siblings[i] <== siblings3[i];
    }
    processor3.oldKey <== oldKey3;
    processor3.oldValue <== oldValue3;
    processor3.isOld0 <== isOld0_3;
    processor3.newKey <== states.key3;
    processor3.newValue <== states.isVouchTx; // 1 for vouch, 0 for unvouch
    processor3.fnc[0] <== states.P3_fnc0;
    processor3.fnc[1] <== states.P3_fnc1;

    // Vouch tree processor 4 (toIdx|fromIdx)
    component processor4 = SMTProcessor(2*nLevels + 1);
    processor4.oldRoot <== processor3.newRoot;
    for (i = 0; i < 2*nLevels + 1; i++) {
        processor4.siblings[i] <== siblings4[i];
    }
    processor4.oldKey <== oldKey4;
    processor4.oldValue <== oldValue4;
    processor4.isOld0 <== isOld0_4;
    processor4.newKey <== states.key4;
    processor4.newValue <== 0; // p4 only works for vouch deletion(Transfer Tx)
    processor4.fnc[0] <== states.P4_fnc0;
    processor4.fnc[1] <== states.P4_fnc1;

    // F - Select Output Roots
    // New account root - if tx is exit, use processor1's root, else use processor2's root
    component selectAccountRoot = Mux1();
    selectAccountRoot.c[0] <== processor2.newRoot;
    selectAccountRoot.c[1] <== processor1.newRoot;
    selectAccountRoot.s <== states.isExit;

    newAccountRoot <== selectAccountRoot.out;

    // New exit root - if tx is exit, use processor2's root, else keep old exit root
    component selectExitRoot = Mux1();
    selectExitRoot.c[0] <== oldExitRoot;
    selectExitRoot.c[1] <== processor2.newRoot;
    selectExitRoot.s <== states.isExit;

    newExitRoot <== selectExitRoot.out;

    // New vouch root comes directly from processor4
    newVouchRoot <== processor4.newRoot;
}
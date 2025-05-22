pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/smt/smtprocessor.circom";
include "../../node_modules/circomlib/circuits/poseidon.circom";
include "../../node_modules/circomlib/circuits/mux1.circom";

include "./lib/hash-state.circom";
include "./balance-updater.circom";
include "./batch-tx-states.circom";

template BatchTx(nLevels) {
    // Inputs - Transaction data
    signal input fromIdx;           // Sender index
    signal input toIdx;             // Receiver index
    signal input amount;            // Transaction amount
    signal input txnType;           // Transaction type (0-6)

    // Inputs - Account state (sender)
    signal input balance1;          // Sender balance
    signal input siblings1[nLevels + 1];
    signal input isOld0_1;
    signal input oldKey1;
    signal input oldValue1;

    // Inputs - Account state (receiver)
    signal input balance2;          // Receiver balance
    signal input siblings2[nLevels + 1];
    signal input isOld0_2;
    signal input oldKey2;
    signal input oldValue2;

    // Inputs - Vouch state (from->to)
    signal input siblings3[2*nLevels + 1];
    signal input isOld0_3;
    signal input oldKey3;
    signal input oldValue3;

    // Inputs - Previous roots
    signal input oldAccountRoot;    // Previous account root
    signal input oldVouchRoot;      // Previous voucher root

    // Outputs - New roots
    signal output newAccountRoot;   // New account root
    signal output newVouchRoot;     // New voucher root

    var i;

    // Calculate states and balance updates
    component states = BatchTxStates(nLevels);
    states.txnType <== txnType;
    states.fromIdx <== fromIdx;
    states.toIdx <== toIdx;

    // Calculate balance updates
    component balanceUpdater = BalanceUpdater();
    balanceUpdater.oldStBalanceSender <== balance1;
    balanceUpdater.oldStBalanceReceiver <== balance2;
    balanceUpdater.amount <== amount;
    balanceUpdater.isCreateAccount <== states.isCreateAccount;
    balanceUpdater.isDeposit <== states.isDeposit;
    balanceUpdater.isWithdraw <== states.isWithdraw;
    balanceUpdater.isExplode <== states.isExplode;

    // Hash new account states
    component newSt1Hash = HashState();
    newSt1Hash.balance <== balanceUpdater.newStBalanceSender;

    component newSt2Hash = HashState();
    newSt2Hash.balance <== balanceUpdater.newStBalanceReceiver;

    // A - Account processor 1 (sender)
    component processor1 = SMTProcessor(nLevels+1);
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

    // B - Account processor 2 (receiver)
    component processor2 = SMTProcessor(nLevels+1);
    processor2.oldRoot <== processor1.newRoot;
    for (i = 0; i < nLevels + 1; i++) {
        processor2.siblings[i] <== siblings2[i];
    }
    processor2.oldKey <== oldKey2;
    processor2.oldValue <== oldValue2;
    processor2.isOld0 <== isOld0_2;
    processor2.newKey <== states.key2;
    processor2.newValue <== newSt2Hash.out;
    processor2.fnc[0] <== states.P2_fnc0;
    processor2.fnc[1] <== states.P2_fnc1;

    // C - Vouch processor
    component processor3 = SMTProcessor(2*nLevels+1);
    processor3.oldRoot <== oldVouchRoot;
    for (i = 0; i < 2*nLevels + 1; i++) {
        processor3.siblings[i] <== siblings3[i];
    }
    processor3.oldKey <== oldKey3;
    processor3.oldValue <== oldValue3;
    processor3.isOld0 <== isOld0_3;
    processor3.newKey <== states.key3;
    processor3.newValue <== states.isVouch; // Insert only on Vouch, Delete only on Unvouch/Explode
    processor3.fnc[0] <== states.P3_fnc0;
    processor3.fnc[1] <== states.P3_fnc1;

    // Set final root values
    newAccountRoot <== processor2.newRoot;
    newVouchRoot <== processor3.newRoot;
}
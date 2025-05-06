pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/mux1.circom";

template BalanceUpdater() {
    // Inputs
    signal input oldStBalanceSender;
    signal input oldStBalanceReceiver;
    signal input amount;
    signal input effectiveExplodeAmount;
    signal input isCreateAccount;
    signal input isDeposit;
    signal input isWithdraw;
    signal input isExplode;
    //signal input nop;

    // Outputs
    signal output newStBalanceSender;
    signal output newStBalanceReceiver;
    
    // Create/Deposit: Existing balance + amount
    signal depositBalance;
    depositBalance <== oldStBalanceSender + amount;
    
    // Withdraw: Existing balance - amount
    signal withdrawBalance;
    withdrawBalance <== oldStBalanceSender - amount;
    
    // Explode: Existing balance + amount (taken from receiver)
    signal explodedBalance;
    explodedBalance <== oldStBalanceSender + effectiveExplodeAmount;
    
    // Selection based on transaction type
    component selectBalSender = Mux1();
    component selectType = Mux1();
    
    // Deposit or account creation
    signal isDepositOrCreate;
    isDepositOrCreate <== isDeposit + isCreateAccount;
    
    selectType.c[0] <== withdrawBalance;
    selectType.c[1] <== depositBalance;
    selectType.s <== isDepositOrCreate;
    
    selectBalSender.c[0] <== selectType.out;
    selectBalSender.c[1] <== explodedBalance;
    selectBalSender.s <== isExplode;
    
    // // Maintain existing balance if NOP
    // component senderMux = Mux1();
    // senderMux.c[0] <== selectBalSender.out;
    // senderMux.c[1] <== oldStBalanceSender;
    // senderMux.s <== nop;
    //newStBalanceSender <== senderMux.out;

    newStBalanceSender <== selectBalSender.out;
    
    // Explode: Existing balance - amount
    signal explodedReceiverBalance;
    explodedReceiverBalance <== oldStBalanceReceiver - effectiveExplodeAmount;
    
    // Subtract amount only if Explode, otherwise no change
    component selectBalReceiver = Mux1();
    selectBalReceiver.c[0] <== oldStBalanceReceiver;
    selectBalReceiver.c[1] <== explodedReceiverBalance;
    selectBalReceiver.s <== isExplode;

    newStBalanceReceiver <== selectBalReceiver.out;
    
    // // Maintain existing balance if NOP
    // Receiver balance update logic (only changes in explode)
    //component receiverMux = Mux1();
    // receiverMux.c[0] <== selectBalReceiver.out;
    // receiverMux.c[1] <== oldStBalanceReceiver;
    // receiverMux.s <== nop;
    // newStBalanceReceiver <== receiverMux.out;
}
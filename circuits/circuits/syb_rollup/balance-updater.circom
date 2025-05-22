pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/mux1.circom";

template BalanceUpdater() {
    // Inputs
    signal input oldStBalanceSender;
    signal input oldStBalanceReceiver;
    signal input amount;
    signal input isCreateAccount;
    signal input isDeposit;
    signal input isWithdraw;
    signal input isExplode;

    // Outputs
    signal output newStBalanceSender;
    signal output newStBalanceReceiver;
    
    // Create/Deposit/Explode: Existing balance + amount
    signal depositBalance;
    depositBalance <== oldStBalanceSender + amount;
    
    // Withdraw: Existing balance - amount
    signal withdrawBalance;
    withdrawBalance <== oldStBalanceSender - amount;
    
    // Selection based on transaction type
    component selectBalSender = Mux1();
    component selectType = Mux1();
    
    // Deposit or account creation
    signal isDepositOrCreateOrExplode;
    isDepositOrCreateOrExplode <== isDeposit + isCreateAccount + isExplode;
    
    selectType.c[0] <== oldStBalanceSender;
    selectType.c[1] <== withdrawBalance;
    selectType.s <== isWithdraw;

    selectBalSender.c[0] <== selectType.out;
    selectBalSender.c[1] <== depositBalance;
    selectBalSender.s <== isDepositOrCreateOrExplode;

    newStBalanceSender <== selectBalSender.out;
    
    // Explode: Existing balance - amount
    signal explodedReceiverBalance;
    explodedReceiverBalance <== oldStBalanceReceiver - amount;
    
    // Subtract amount only if Explode, otherwise no change
    component selectBalReceiver = Mux1();
    selectBalReceiver.c[0] <== oldStBalanceReceiver;
    selectBalReceiver.c[1] <== explodedReceiverBalance;
    selectBalReceiver.s <== isExplode;

    newStBalanceReceiver <== selectBalReceiver.out;
}
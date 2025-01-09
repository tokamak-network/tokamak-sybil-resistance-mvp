pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/bitify.circom";
include "../../node_modules/circomlib/circuits/comparators.circom";

/**
 * Updates balances for a transaction
 */
template BalanceUpdater() {
    signal input oldStBalanceSender;
    signal input oldStBalanceReceiver;
    signal input amount;
    signal input loadAmount;
    signal input nop;
    signal input nullifyLoadAmount;
    signal input nullifyAmount;

    signal output newStBalanceSender;
    signal output newStBalanceReceiver;
    signal output isP2Nop;
    signal output isAmountNullified;

    signal underflowOk;             // 1 if sender balance is > 0
    signal effectiveAmount1;        // original amount to transfer. Set to 0 if tx is NOP
    signal effectiveAmount2;        // tx amount once nullifyAmount is applied
    signal effectiveAmount3;        // tx amount once checked if sender has enough balance
    signal effectiveLoadAmount1;    // original loadAmount to load
    signal effectiveLoadAmount2;    // tx loadAmount once nullifyLoadAmount is applied

    // compute effective loadAmount and amount
    effectiveLoadAmount1 <== loadAmount;
    effectiveLoadAmount2 <== effectiveLoadAmount1*(1-nullifyLoadAmount);
    effectiveAmount1 <== amount*(1-nop);                     //nop makes amount 0
    effectiveAmount2 <== effectiveAmount1*(1-nullifyAmount); //nullifyAmount makes amount 0

    // check balance sender
    // Overflow check:
    // - smart contract does not allow deposits over 2^128
    // - smart contract does not allow transfers over 2^192
    // - it is assumed that maximum balance accumulated would be 2^192
    // Underflow check:
    // - assuming 192 bits as maximum allowed balance for a single account
    // - bit 193 is set to 1
    // - if account has not enough balance, bit 193 will be 0
    component n2bSender = Num2Bits(193);
    n2bSender.in <== (1<<192) + oldStBalanceSender + effectiveLoadAmount2 - effectiveAmount2;

    underflowOk <== n2bSender.out[192];

    effectiveAmount3 <== underflowOk*effectiveAmount2;

    newStBalanceSender <== oldStBalanceSender + effectiveLoadAmount2 - effectiveAmount3;
    newStBalanceReceiver <== oldStBalanceReceiver + effectiveAmount3;

    component effectiveAmountIsZero = IsZero();
    effectiveAmountIsZero.in <== effectiveAmount1;

    isAmountNullified <== 1 - (1 - nullifyAmount)*underflowOk;

    // Set NOP function on processor 2 (receiver account) if original amount to transfer is 0
    isP2Nop <== (1 - effectiveAmountIsZero.out);
}
pragma circom 2.0.0;

include "../../../node_modules/circomlib/circuits/poseidon.circom";

template HashState() {
    signal input nonce;
    signal input balance;
    signal input ethAddr;

    signal output out;

    //signal e0; // build e0 element

    //e0 <== nonce + ethAddr * (1 << 40);

    component hash = Poseidon(3);

    hash.inputs[0] <== nonce;
    hash.inputs[1] <== balance;
    hash.inputs[2] <== ethAddr;

    hash.out ==> out;
}
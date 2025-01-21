pragma circom 2.0.0;

include "../../../node_modules/circomlib/circuits/poseidon.circom";

template HashState() {
    //signal input nonce;
    signal input balance;
    signal input ethAddr;

    signal output out;

    component hash = Poseidon(2);

    //hash.inputs[0] <== nonce;
    hash.inputs[0] <== balance;
    hash.inputs[1] <== ethAddr;

    hash.out ==> out;
}
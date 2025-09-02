pragma circom 2.0.0;
include "../node_modules/circomlib/circuits/poseidon.circom";

template DualMux(){
    signal input in[2];
    signal input s;
    signal output out[2];
    
    s*(s-1) === 0;

    out[0] <== (in[1] - in[0])*s + in[0];
    out[1] <== (in[0] - in[1])*s + in[1];
}

template GetMerkleRoot(MERKLE_TREE_DEPTH){
    signal input leaf;
    signal input paths2_root[MERKLE_TREE_DEPTH];
    signal input paths2_root_pos[MERKLE_TREE_DEPTH];

    signal output out;

    component selectors[MERKLE_TREE_DEPTH];
    component hashers[MERKLE_TREE_DEPTH];

    for(var i = 0; i < MERKLE_TREE_DEPTH; i++){
        selectors[i] = DualMux();
        selectors[i].in[0] <== i == 0 ? leaf : hashers[i-1].out;
        selectors[i].in[1] <== paths2_root[i];
        selectors[i].s <== paths2_root_pos[i];

        hashers[i] = Poseidon(2);
        hashers[i].inputs[0] <== selectors[i].out[0];
        hashers[i].inputs[1] <== selectors[i].out[1];
    }

    out <== hashers[MERKLE_TREE_DEPTH-1].out;
}
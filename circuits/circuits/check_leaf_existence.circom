pragma circom 2.0.0;
include "./get_merkle_root.circom";
include "../node_modules/circomlib/circuits/poseidon.circom";

template LeafExistence(MERKLE_TREE_DEPTH, PREIMAGE_LENGTH){
    signal input preimage[PREIMAGE_LENGTH];
    signal input root;
    signal input paths2_root_pos[MERKLE_TREE_DEPTH];
    signal input paths2_root[MERKLE_TREE_DEPTH];

    component leaf = Poseidon(PREIMAGE_LENGTH);
    for(var i = 0; i < PREIMAGE_LENGTH; i++){
        leaf.inputs[i] <== preimage[i];
    }

    component computed_root = GetMerkleRoot(MERKLE_TREE_DEPTH);
    computed_root.leaf <== leaf.out;

    for (var w = 0; w < MERKLE_TREE_DEPTH; w++){
        computed_root.paths2_root[w] <== paths2_root[w];
        computed_root.paths2_root_pos[w] <== paths2_root_pos[w];
    }

    // equality constraint: input tx root === computed tx root 
    root === computed_root.out;
}
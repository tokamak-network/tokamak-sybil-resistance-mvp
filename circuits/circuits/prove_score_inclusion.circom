pragma circom 2.0.0;

include "../node_modules/circomlib/circuits/smt/smtverifier.circom";

// This template is a wrapper around the circomlib SMTVerifier to prove
// the inclusion of a score in the tree.
// It takes the user's data as input and wires it to the SMTVerifier.
template ProveScoreInclusion(nLevels) {
    // Inputs from the user/prover
    signal input idx;
    signal input score;
    signal input root;
    signal input siblings[nLevels];

    // Instantiate the SMTVerifier from circomlib
    component verifier = SMTVerifier(nLevels);

    verifier.enabled <== 1;
    verifier.root <== root;
    for (var i = 0; i < nLevels; i++) {
        verifier.siblings[i] <== siblings[i];
    }

    verifier.oldKey <== 0;
    verifier.oldValue <== 0;
    verifier.isOld0 <== 0;
    verifier.key <== idx;
    verifier.value <== score;
    verifier.fnc <== 0; // fnc to use inclusion proof
}
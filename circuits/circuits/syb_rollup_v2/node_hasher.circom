pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/poseidon.circom";
include "../../node_modules/circomlib/circuits/comparators.circom";

// NodeHasher: Computes NodeHash_G(v) for a vertex in the GraphTree
//
// NodeHash Algorithm:
// - First block B_0: [v, d, u_0, ..., u_13]  (14 neighbors)
// - acc = Poseidon_16(B_0)
// - For remaining neighbors, process in chunks of 15:
//   - B_b = [acc, u_14+(b-1)*15, ..., u_14+(b-1)*15+14]
//   - acc = Poseidon_16(B_b)
//
// Key Features:
// - Enforces that nbr_arr[i] == 0 when i >= d (proper padding verification)
// - Enforces strictly ascending order: nbr_arr[i] < nbr_arr[i+1] for i in 0..d-2
// - NO domain separation tags
// - First block: 14 neighbors, Continuation blocks: 15 neighbors each
//
// Parameters:
//   maxDeg - Maximum degree a node can have (fixed at compile time)
//
// Inputs:
//   v - Vertex ID
//   d - Degree of vertex (number of actual neighbors)
//   nbr_arr[padLen] - Neighbor array MUST be padded with zeros when i >= d
//                     where padLen = 14 + 15*numR
//
// Output:
//   hash - NodeHash_G(v)
//
template NodeHasher(maxDeg) {
    assert(maxDeg >= 1);
    // TODO: do we need to check anything with maxDeg?

    signal input v;               // Vertex ID
    signal input d;               // Degree (actual number of neighbors)

    // Calculate number of continuation rounds needed
    // numR = ceil((maxDeg - 14) / 15)
    var numR = maxDeg <= 14 ? 0 : (maxDeg - 14 + 14) \ 15;

    // padLen = 14 + 15 * numR (exactly fits into numR + 1 hashing rounds)
    var padLen = 14 + 15 * numR;

    signal input nbr_arr[padLen]; // Neighbor array (must be properly padded)
    signal output hash;

    component isInPadding[padLen];                  // Checks if index i >= d (in padding region)
    component isNextInPadding[padLen - 1];          // Checks if index i+1 >= d
    component isStrictlyAscending[padLen - 1];      // Checks if nbr_arr[i] < nbr_arr[i+1]

    // INPUT VALIDATION CHECKS
    // 1. Zero-padding check: nbr_arr[i] == 0 when i >= d
    // 2. Strictly ascending check: nbr_arr[i] < nbr_arr[i+1] when both i and i+1 < d
    // Array regions:
    //   [0 ... d-1]         : Valid neighbors (must be strictly ascending)
    //   [d ... padLen-1]    : Padding zeros   (must all be 0)
    for (var i = 0; i < padLen; i++) {
        // Check if i >= d (in padding region)
        isInPadding[i] = GreaterEqThan(32);
        isInPadding[i].in[0] <== i;
        isInPadding[i].in[1] <== d;

        // When i >= d, nbr_arr[i] must be 0
        isInPadding[i].out * nbr_arr[i] === 0;

        // Strictly ascending check (skip last element since there's no i+1)
        if (i < padLen - 1) {
            // Check if i+1 >= d
            isNextInPadding[i] = GreaterEqThan(32);
            isNextInPadding[i].in[0] <== i + 1;
            isNextInPadding[i].in[1] <== d;

            // Check if nbr_arr[i] < nbr_arr[i+1]
            isStrictlyAscending[i] = LessThan(252);
            isStrictlyAscending[i].in[0] <== nbr_arr[i];
            isStrictlyAscending[i].in[1] <== nbr_arr[i + 1];

            // Enforce ascending only when both i and i+1 are valid neighbors (i+1 < d)
            // (1 - isNextInPadding[i].out) means: i+1 < d
            (1 - isNextInPadding[i].out) * (1 - isStrictlyAscending[i].out) === 0;
        }
    }

    // HASHING: First block [v, d, nbr[0..13]]
    component firstHash = Poseidon(16);
    firstHash.inputs[0] <== v;
    firstHash.inputs[1] <== d;
    for (var i = 0; i < 14; i++) {
        firstHash.inputs[2 + i] <== nbr_arr[i];
    }

    // HASHING: Continuation blocks (15 neighbors each)
    signal acc[numR + 1];
    acc[0] <== firstHash.out;

    component contHash[numR];
    for (var round = 0; round < numR; round++) {
        contHash[round] = Poseidon(16);
        contHash[round].inputs[0] <== acc[round];

        var startIdx = 14 + round * 15;
        for (var i = 0; i < 15; i++) {
            var idx = startIdx + i;
            contHash[round].inputs[1 + i] <== nbr_arr[idx];
        }

        acc[round + 1] <== contHash[round].out;
    }

    hash <== acc[numR];
}

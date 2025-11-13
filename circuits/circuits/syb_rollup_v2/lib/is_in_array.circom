pragma circom 2.0.0;

include "../../../node_modules/circomlib/circuits/comparators.circom";

// IsInArray: Checks if a target value is present in an array
// 
// This template checks each position in the array and counts how many times
// the target appears. It's useful for checking if an edge already exists
// when adding a new edge to a graph (by checking if out == 0).
//
// Parameters:
//   n: Length of the array to check
//
// Inputs:
//   arr[n]: The array to search through
//   target: The value to search for
//
// Outputs:
//   out: Count of how many times target appears in the array
//        0 = target not found
//        1+ = target found (count of occurrences)
//
// Note: When used with NodeHasher-validated arrays, duplicates are already
//       prevented by NodeHasher's strictly ascending constraint.

template IsInArray(n) {
    signal input arr[n];
    signal input target;
    signal output out;
    
    // Check each position in the array
    component isEqual[n];
    signal matches[n];
    
    for (var i = 0; i < n; i++) {
        isEqual[i] = IsEqual();
        isEqual[i].in[0] <== arr[i];
        isEqual[i].in[1] <== target;
        matches[i] <== isEqual[i].out;
    }
    
    // Sum all matches - should be 0 if target is not in array
    signal sum[n + 1];
    sum[0] <== 0;
    
    for (var i = 0; i < n; i++) {
        sum[i + 1] <== sum[i] + matches[i];
    }
    
    // sum[n] == 0: target is not in the array
    // sum[n] >= 1: target is in the array (count of occurrences)
    // Note: Duplicate prevention is handled by NodeHasher's strictly ascending check
    out <== sum[n];
}


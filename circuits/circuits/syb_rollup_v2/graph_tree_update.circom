pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/smt/smtprocessor.circom";
include "../../node_modules/circomlib/circuits/comparators.circom";
include "./node_hasher.circom";

// GraphTreeUpdate: Updates the GraphTree when adding an edge {u,v}
//
// This circuit:
// 1. Verifies preconditions (u != v, degrees within bounds)
// 2. Computes old NbrHash for u and v (before adding edge)
// 3. Computes new NbrHash for u and v (after adding edge)
// 4. Updates the Merkle tree for both leaves u and v
//
// Parameters:
//   nLevels - Depth of the Merkle tree (tree can hold 2^nLevels vertices)
//   maxDeg  - Maximum degree allowed for any vertex
//
// Inputs:
//   u, v - The two vertices to connect (edge endpoints)
//   
//   oldDegU, oldDegV - Old degrees (before adding edge)
//   newDegU, newDegV - New degrees (after adding edge, should be oldDeg + 1)
//   
//   oldNbrArrU[padLen], oldNbrArrV[padLen] - Old neighbor arrays (before edge)
//   newNbrArrU[padLen], newNbrArrV[padLen] - New neighbor arrays (after edge)
//   
//   siblingsU[nLevels+1] - Merkle proof for vertex u
//   siblingsV[nLevels+1] - Merkle proof for vertex v
//   
//   oldRoot - Old GraphTree root (before adding edge)
//
// Outputs:
//   newRoot - New GraphTree root (after adding edge)
//
template GraphTreeUpdate(nLevels, maxDeg) {
    // Calculate padLen for neighbor arrays
    var numR = (maxDeg + 14) \ 15;
    var padLen = 15 * numR;

    // ===== INPUTS =====
    signal input u;
    signal input v;
    
    signal input oldDegU;
    signal input oldDegV;
    signal input newDegU;
    signal input newDegV;
    
    signal input oldNbrArrU[padLen];
    signal input oldNbrArrV[padLen];
    signal input newNbrArrU[padLen];
    signal input newNbrArrV[padLen];
    
    signal input siblingsU[nLevels + 1];
    signal input siblingsV[nLevels + 1];
    
    signal input oldRoot;
    
    // ===== OUTPUTS =====
    signal output newRoot;
    
    // ===== PRECONDITION CHECKS =====
    
    // 1. Check u != v
    component uNotEqV = IsEqual();
    uNotEqV.in[0] <== u;
    uNotEqV.in[1] <== v;
    uNotEqV.out === 0; // Must be different
    
    // 2. Check newDegU = oldDegU + 1
    newDegU === oldDegU + 1;
    
    // 3. Check newDegV = oldDegV + 1
    newDegV === oldDegV + 1;
    
    // 4. Check newDegU <= maxDeg
    component checkMaxDegU = LessEqThan(32);
    checkMaxDegU.in[0] <== newDegU;
    checkMaxDegU.in[1] <== maxDeg;
    checkMaxDegU.out === 1;
    
    // 5. Check newDegV <= maxDeg
    component checkMaxDegV = LessEqThan(32);
    checkMaxDegV.in[0] <== newDegV;
    checkMaxDegV.in[1] <== maxDeg;
    checkMaxDegV.out === 1;
    
    // COMPUTE OLD HASHES (before adding edge)
    
    component oldHashU = NodeHasher(maxDeg);
    oldHashU.d <== oldDegU;
    for (var i = 0; i < padLen; i++) {
        oldHashU.nbr_arr[i] <== oldNbrArrU[i];
    }
    
    component oldHashV = NodeHasher(maxDeg);
    oldHashV.d <== oldDegV;
    for (var i = 0; i < padLen; i++) {
        oldHashV.nbr_arr[i] <== oldNbrArrV[i];
    }
    
    //COMPUTE NEW HASHES (after adding edge)
    
    component newHashU = NodeHasher(maxDeg);
    newHashU.d <== newDegU;
    for (var i = 0; i < padLen; i++) {
        newHashU.nbr_arr[i] <== newNbrArrU[i];
    }
    
    component newHashV = NodeHasher(maxDeg);
    newHashV.d <== newDegV;
    for (var i = 0; i < padLen; i++) {
        newHashV.nbr_arr[i] <== newNbrArrV[i];
    }
    
    // UPDATE MERKLE TREE
    
    // First update: Update vertex u's leaf
    // Function: UPDATE (fnc = [0, 1])
    component processorU = SMTProcessor(nLevels + 1);
    processorU.oldRoot <== oldRoot;
    for (var i = 0; i < nLevels + 1; i++) {
        processorU.siblings[i] <== siblingsU[i];
    }
    processorU.oldKey <== u;
    processorU.oldValue <== oldHashU.hash;
    processorU.isOld0 <== 0; // Not inserting new leaf, updating existing
    processorU.newKey <== u;
    processorU.newValue <== newHashU.hash;
    processorU.fnc[0] <== 0; // UPDATE operation
    processorU.fnc[1] <== 1; // UPDATE operation
    
    // Second update: Update vertex v's leaf
    // Function: UPDATE (fnc = [0, 1])
    component processorV = SMTProcessor(nLevels + 1);
    processorV.oldRoot <== processorU.newRoot; // Chain from first update
    for (var i = 0; i < nLevels + 1; i++) {
        processorV.siblings[i] <== siblingsV[i];
    }
    processorV.oldKey <== v;
    processorV.oldValue <== oldHashV.hash;
    processorV.isOld0 <== 0; // Not inserting new leaf, updating existing
    processorV.newKey <== v;
    processorV.newValue <== newHashV.hash;
    processorV.fnc[0] <== 0; // UPDATE operation
    processorV.fnc[1] <== 1; // UPDATE operation
    
    // Output the final root
    newRoot <== processorV.newRoot;
}


import fs from "fs";
import path from "path";
import { describe, it, before, after } from "mocha";
import assert from "assert";
import { wasm as tester } from "circom_tester";
import { buildPoseidon } from "circomlibjs";
import { fileURLToPath } from "url";
import { SmtTree } from "../utils/smt.js";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

describe("GraphTreeUpdate circuit test", function () {
  this.timeout(300000);

  const N_LEVELS = 4; // Tree depth (supports 2^4 = 16 vertices)
  const MAX_DEG = 15 * 4; // Maximum degree: 60
  let circuit;
  let circuitTmpPath;
  let poseidon;
  let F;

  // Calculate padLen based on maxDeg
  function calculatePadLen(maxDeg) {
    const numR = Math.ceil(maxDeg / 15);
    return 15 * numR;
  }

  const PAD_LEN = calculatePadLen(MAX_DEG);

  before(async () => {
    // Initialize Poseidon hasher
    poseidon = await buildPoseidon();
    F = poseidon.F;

    // Create circuit with nLevels=4, maxDeg=60
    const circuitSrc = `
            pragma circom 2.0.0;
            include "../circuits/syb_rollup_v2/graph_tree_update.circom";
            component main = GraphTreeUpdate(${N_LEVELS}, ${MAX_DEG});
        `;
    circuitTmpPath = path.join(__dirname, "graph-tree-update-temp.circom");
    fs.writeFileSync(circuitTmpPath, circuitSrc, "utf8");

    circuit = await tester(circuitTmpPath, {
      reduceConstraints: false,
      include: path.join(__dirname, "../"),
    });
    await circuit.loadConstraints();
    console.log(`\n✓ GraphTreeUpdate circuit compiled`);
    console.log(`  nLevels=${N_LEVELS}, maxDeg=${MAX_DEG}`);
    console.log(`  padLen=${PAD_LEN}`);
    console.log(`  Constraints: ${circuit.constraints.length}\n`);
  });

  after(() => {
    if (fs.existsSync(circuitTmpPath)) {
      fs.unlinkSync(circuitTmpPath);
    }
  });

  // Helper function to compute NbrHash
  function computeNbrHash(d, neighbors) {
    const paddedNbrs = [...neighbors];
    while (paddedNbrs.length < PAD_LEN) {
      paddedNbrs.push(0);
    }

    // First block: B_0 = [d, nbr[0..14]] (15 neighbors)
    const firstBlock = [d];
    for (let i = 0; i < 15; i++) {
      firstBlock.push(paddedNbrs[i] || 0);
    }

    let acc = F.toString(poseidon(firstBlock));

    // Continuation blocks (15 neighbors each)
    const numR = Math.ceil(MAX_DEG / 15);

    for (let round = 1; round < numR; round++) {
      const block = [BigInt(acc)];
      const startIdx = 15 + (round - 1) * 15;

      for (let i = 0; i < 15; i++) {
        const idx = startIdx + i;
        block.push(paddedNbrs[idx] || 0);
      }

      acc = F.toString(poseidon(block));
    }

    return acc;
  }

  // Helper to pad neighbor array
  function padNeighbors(neighbors) {
    const padded = [...neighbors];
    while (padded.length < PAD_LEN) {
      padded.push(0);
    }
    return padded.map((x) => x.toString());
  }

  // Helper to ensure siblings array has exactly nLevels + 1 elements
  // SmtTree.getSiblings returns nLevels elements, but SMTProcessor needs nLevels + 1
  function ensureSiblingsLength(siblings) {
    const padded = [...siblings];
    while (padded.length < N_LEVELS + 1) {
      padded.push(0);
    }
    return padded.map((x) => x.toString());
  }

  /**
   * TEST CASES
   * 
   * VALID UPDATES
   * [X] update GraphTree when adding edge {1,2} (no existing edges)
   * [X] update GraphTree when adding edge {2,5} with existing edges
   * 
   * INVALID UPDATES - Precondition Failures
   * [X] fail when u equals v (no self-loops)
   * [X] fail when u or v is 0 (reserved index)
   * [X] fail when u or v exceeds 2^nLevels (out of tree bounds)
   * [X] fail when newDegU != oldDegU + 1 (invalid degree increment for U)
   * [X] fail when newDegV != oldDegV + 1 (invalid degree increment for V)
   * [X] fail when degree exceeds maxDeg
   * [X] fail when edge already exists (duplicate edge prevention)
   */

  it("should update GraphTree when adding edge {1,2}", async () => {
    // Initial state: vertices 1 and 2 have no edges
    const u = 1;
    const v = 2;

    // Old state (before adding edge)
    const oldDegU = 0;
    const oldDegV = 0;
    const oldNbrArrU = [];
    const oldNbrArrV = [];

    // New state (after adding edge {1,2})
    const newDegU = 1;
    const newDegV = 1;
    const newNbrArrU = [2]; // vertex 1 now connected to vertex 2
    const newNbrArrV = [1]; // vertex 2 now connected to vertex 1

    // Compute hashes
    const oldHashU = BigInt(computeNbrHash(oldDegU, oldNbrArrU));
    const oldHashV = BigInt(computeNbrHash(oldDegV, oldNbrArrV));
    const newHashU = BigInt(computeNbrHash(newDegU, newNbrArrU));
    const newHashV = BigInt(computeNbrHash(newDegV, newNbrArrV));

    // Build initial tree with old hashes at leaves 0 and 1
    const tree = new SmtTree(N_LEVELS);
    await tree.init();
    await tree.insert(u, oldHashU);
    await tree.insert(v, oldHashV);

    const oldRoot = await tree.getRoot();

    // Get Merkle proof for U from original tree
    const siblingsU = ensureSiblingsLength(await tree.getSiblings(u));

    // Update U to get intermediate tree state
    await tree.update(u, newHashU);

    // Get Merkle proof for V from tree after U update
    const siblingsV = ensureSiblingsLength(await tree.getSiblings(v));

    // Prepare circuit input
    const input = {
      u: u.toString(),
      v: v.toString(),
      oldDegU: oldDegU.toString(),
      oldDegV: oldDegV.toString(),
      newDegU: newDegU.toString(),
      newDegV: newDegV.toString(),
      oldNbrArrU: padNeighbors(oldNbrArrU),
      oldNbrArrV: padNeighbors(oldNbrArrV),
      newNbrArrU: padNeighbors(newNbrArrU),
      newNbrArrV: padNeighbors(newNbrArrV),
      siblingsU: siblingsU,
      siblingsV: siblingsV,
      oldRoot: F.toString(oldRoot),
    };

    // Calculate witness
    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitNewRoot = w[1].toString();

    // Compute expected new root by updating V (U already updated for proof generation)
    await tree.update(v, newHashV);
    const expectedNewRoot = F.toString(await tree.getRoot());

    assert.equal(circuitNewRoot, expectedNewRoot);
  });

  

  it("should update GraphTree when adding edge {2,5} with existing edges", async () => {
    // Vertex 2 already has edges to [1, 3]
    // Vertex 5 already has edges to [1, 4]
    // Now adding edge {2,5}
    const u = 2;
    const v = 5;

    // Old state (before adding edge)
    const oldDegU = 2;
    const oldDegV = 2;
    const oldNbrArrU = [1, 3];
    const oldNbrArrV = [1, 4];

    // New state (after adding edge {2,5})
    const newDegU = 3;
    const newDegV = 3;
    const newNbrArrU = [1, 3, 5]; // Added 5
    const newNbrArrV = [1, 4, 5]; // Added 2 (should be sorted, so [1, 2, 4])

    // Fix: neighbors must be sorted
    const newNbrArrVSorted = [1, 2, 4];

    // Compute hashes
    const oldHashU = BigInt(computeNbrHash(oldDegU, oldNbrArrU));
    const oldHashV = BigInt(computeNbrHash(oldDegV, oldNbrArrV));
    const newHashU = BigInt(computeNbrHash(newDegU, newNbrArrU));
    const newHashV = BigInt(computeNbrHash(newDegV, newNbrArrVSorted));

    // Build initial tree with old hashes
    const tree = new SmtTree(N_LEVELS);
    await tree.init();
    await tree.insert(u, oldHashU);
    await tree.insert(v, oldHashV);

    const oldRoot = await tree.getRoot();

    // Get Merkle proof for U from original tree
    const siblingsU = ensureSiblingsLength(await tree.getSiblings(u));

    // Update U to get intermediate tree state
    await tree.update(u, newHashU);

    // Get Merkle proof for V from tree after U update
    const siblingsV = ensureSiblingsLength(await tree.getSiblings(v));

    // Prepare circuit input
    const input = {
      u: u.toString(),
      v: v.toString(),
      oldDegU: oldDegU.toString(),
      oldDegV: oldDegV.toString(),
      newDegU: newDegU.toString(),
      newDegV: newDegV.toString(),
      oldNbrArrU: padNeighbors(oldNbrArrU),
      oldNbrArrV: padNeighbors(oldNbrArrV),
      newNbrArrU: padNeighbors(newNbrArrU),
      newNbrArrV: padNeighbors(newNbrArrVSorted),
      siblingsU: siblingsU,
      siblingsV: siblingsV,
      oldRoot: F.toString(oldRoot),
    };

    // Calculate witness
    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitNewRoot = w[1].toString();

    // Compute expected new root by updating V (U already updated for proof generation)
    await tree.update(v, newHashV);
    const expectedNewRoot = F.toString(await tree.getRoot());

    assert.equal(circuitNewRoot, expectedNewRoot);
  });

  it("should fail when u equals v", async () => {
    const u = 3;
    const v = 3; // Same as u!

    const oldDegU = 1;
    const oldDegV = 1;
    const oldNbrArrU = [5];
    const oldNbrArrV = [5];

    const newDegU = 2;
    const newDegV = 2;
    const newNbrArrU = [3, 5];
    const newNbrArrV = [3, 5];

    // Build tree
    const oldHashU = BigInt(computeNbrHash(oldDegU, oldNbrArrU));
    const oldHashV = BigInt(computeNbrHash(oldDegV, oldNbrArrV));

    const tree = new SmtTree(N_LEVELS);
    await tree.init();
    await tree.insert(u, oldHashU);

    const oldRoot = await tree.getRoot();
    const siblingsU = ensureSiblingsLength(await tree.getSiblings(u));

    const input = {
      u: u.toString(),
      v: v.toString(),
      oldDegU: oldDegU.toString(),
      oldDegV: oldDegV.toString(),
      newDegU: newDegU.toString(),
      newDegV: newDegV.toString(),
      oldNbrArrU: padNeighbors(oldNbrArrU),
      oldNbrArrV: padNeighbors(oldNbrArrV),
      newNbrArrU: padNeighbors(newNbrArrU),
      newNbrArrV: padNeighbors(newNbrArrV),
      siblingsU: siblingsU,
      siblingsV: siblingsU,
      oldRoot: F.toString(oldRoot),
    };

    try {
      await circuit.calculateWitness(input, true);
      assert.fail("Should have failed with u == v");
    } catch (error) {
      assert(error.message.includes("Assert Failed"));
    }
  });

  it("should fail when newDegU != oldDegU + 1", async () => {
    const u = 3;
    const v = 7;

    const oldDegU = 2;
    const oldDegV = 1;
    const oldNbrArrU = [1, 5];
    const oldNbrArrV = [2];

    const newDegU = 4; // Invalid! Should be 3 (oldDegU + 1), but claiming 4
    const newDegV = 2; // Valid
    const newNbrArrU = [1, 5, 7];
    const newNbrArrV = [2, 3];

    // Build tree
    const oldHashU = BigInt(computeNbrHash(oldDegU, oldNbrArrU));
    const oldHashV = BigInt(computeNbrHash(oldDegV, oldNbrArrV));

    const tree = new SmtTree(N_LEVELS);
    await tree.init();
    await tree.insert(u, oldHashU);
    await tree.insert(v, oldHashV);

    const oldRoot = await tree.getRoot();
    const siblingsU = ensureSiblingsLength(await tree.getSiblings(u));
    const siblingsV = ensureSiblingsLength(await tree.getSiblings(v));

    const input = {
      u: u.toString(),
      v: v.toString(),
      oldDegU: oldDegU.toString(),
      oldDegV: oldDegV.toString(),
      newDegU: newDegU.toString(), // Claiming 4 instead of 3!
      newDegV: newDegV.toString(),
      oldNbrArrU: padNeighbors(oldNbrArrU),
      oldNbrArrV: padNeighbors(oldNbrArrV),
      newNbrArrU: padNeighbors(newNbrArrU),
      newNbrArrV: padNeighbors(newNbrArrV),
      siblingsU: siblingsU,
      siblingsV: siblingsV,
      oldRoot: F.toString(oldRoot),
    };

    try {
      await circuit.calculateWitness(input, true);
      assert.fail("Should have failed with newDegU != oldDegU + 1");
    } catch (error) {
      assert(error.message.includes("Assert Failed"));
    }
  });

  it("should fail when newDegV != oldDegV + 1", async () => {
    const u = 4;
    const v = 8;

    const oldDegU = 1;
    const oldDegV = 3;
    const oldNbrArrU = [2];
    const oldNbrArrV = [1, 5, 9];

    const newDegU = 2; // Valid
    const newDegV = 5; // Invalid! Should be 4 (oldDegV + 1), but claiming 5
    const newNbrArrU = [2, 8];
    const newNbrArrV = [1, 4, 5, 9];

    // Build tree
    const oldHashU = BigInt(computeNbrHash(oldDegU, oldNbrArrU));
    const oldHashV = BigInt(computeNbrHash(oldDegV, oldNbrArrV));

    const tree = new SmtTree(N_LEVELS);
    await tree.init();
    await tree.insert(u, oldHashU);
    await tree.insert(v, oldHashV);

    const oldRoot = await tree.getRoot();
    const siblingsU = ensureSiblingsLength(await tree.getSiblings(u));

    // Update U first
    const newHashU = BigInt(computeNbrHash(newDegU, newNbrArrU));
    await tree.update(u, newHashU);

    const siblingsV = ensureSiblingsLength(await tree.getSiblings(v));

    const input = {
      u: u.toString(),
      v: v.toString(),
      oldDegU: oldDegU.toString(),
      oldDegV: oldDegV.toString(),
      newDegU: newDegU.toString(),
      newDegV: newDegV.toString(), // Claiming 5 instead of 4!
      oldNbrArrU: padNeighbors(oldNbrArrU),
      oldNbrArrV: padNeighbors(oldNbrArrV),
      newNbrArrU: padNeighbors(newNbrArrU),
      newNbrArrV: padNeighbors(newNbrArrV),
      siblingsU: siblingsU,
      siblingsV: siblingsV,
      oldRoot: F.toString(oldRoot),
    };

    try {
      await circuit.calculateWitness(input, true);
      assert.fail("Should have failed with newDegV != oldDegV + 1");
    } catch (error) {
      assert(error.message.includes("Assert Failed"));
    }
  });

  it("should fail when degree exceeds maxDeg", async () => {
    const u = 7;
    const v = 8;

    const oldDegU = MAX_DEG; // Already at max!
    const oldDegV = 0;
    const oldNbrArrU = Array.from({ length: MAX_DEG }, (_, i) => i + 1);
    const oldNbrArrV = [];

    const newDegU = MAX_DEG + 1; // Exceeds max!
    const newDegV = 1;
    // Keep newNbrArrU at MAX_DEG length (can't exceed padLen in the input)
    // but claim degree is MAX_DEG + 1
    const newNbrArrU = oldNbrArrU; // Still MAX_DEG elements
    const newNbrArrV = [u];

    // Build tree
    const oldHashU = BigInt(computeNbrHash(oldDegU, oldNbrArrU));
    const oldHashV = BigInt(computeNbrHash(oldDegV, oldNbrArrV));

    const tree = new SmtTree(N_LEVELS);
    await tree.init();
    await tree.insert(u, oldHashU);
    await tree.insert(v, oldHashV);

    const oldRoot = await tree.getRoot();
    const siblingsU = ensureSiblingsLength(await tree.getSiblings(u));
    const siblingsV = ensureSiblingsLength(await tree.getSiblings(v));

    const input = {
      u: u.toString(),
      v: v.toString(),
      oldDegU: oldDegU.toString(),
      oldDegV: oldDegV.toString(),
      newDegU: newDegU.toString(),
      newDegV: newDegV.toString(),
      oldNbrArrU: padNeighbors(oldNbrArrU),
      oldNbrArrV: padNeighbors(oldNbrArrV),
      newNbrArrU: padNeighbors(newNbrArrU),
      newNbrArrV: padNeighbors(newNbrArrV),
      siblingsU: siblingsU,
      siblingsV: siblingsV,
      oldRoot: F.toString(oldRoot),
    };

    try {
      await circuit.calculateWitness(input, true);
      assert.fail("Should have failed with degree > maxDeg");
    } catch (error) {
      assert(error.message.includes("Assert Failed"));
    }
  });

  it("should fail if u or v is 0 (reserved index)", async () => {
    const u = 0; // Invalid - index 0 is reserved!
    const v = 5; // Valid
    
    const oldDegU = 0;
    const oldDegV = 0;
    const oldNbrArrU = [];
    const oldNbrArrV = [];

    const newDegU = 1;
    const newDegV = 1;
    const newNbrArrU = [v];
    const newNbrArrV = [u];

    // Build tree with old hashes
    const oldHashU = BigInt(computeNbrHash(oldDegU, oldNbrArrU));
    const oldHashV = BigInt(computeNbrHash(oldDegV, oldNbrArrV));

    const tree = new SmtTree(N_LEVELS);
    await tree.init();
    await tree.insert(v, oldHashV);

    const oldRoot = await tree.getRoot();
    // Dummy siblings for u (which is invalid)
    const siblingsU = Array(N_LEVELS + 1).fill("0");
    const siblingsV = ensureSiblingsLength(await tree.getSiblings(v));

    const input = {
      u: u.toString(), // 0 - reserved/invalid!
      v: v.toString(),
      oldDegU: oldDegU.toString(),
      oldDegV: oldDegV.toString(),
      newDegU: newDegU.toString(),
      newDegV: newDegV.toString(),
      oldNbrArrU: padNeighbors(oldNbrArrU),
      oldNbrArrV: padNeighbors(oldNbrArrV),
      newNbrArrU: padNeighbors(newNbrArrU),
      newNbrArrV: padNeighbors(newNbrArrV),
      siblingsU: siblingsU,
      siblingsV: siblingsV,
      oldRoot: F.toString(oldRoot),
    };

    try {
      await circuit.calculateWitness(input, true);
      assert.fail("Should have failed with u=0 (reserved index)");
    } catch (error) {
      assert(error.message.includes("Assert Failed"));
    }
  });

  it("should fail when edge already exists", async () => {
    const u = 6;
    const v = 10;

    // Old state: edge {6,10} ALREADY EXISTS
    const oldDegU = 2;
    const oldDegV = 1;
    const oldNbrArrU = [3, 10]; // v=10 is already in u's neighbor list!
    const oldNbrArrV = [6]; // u=6 is already in v's neighbor list!

    // Attempting to add edge {6,10} again (duplicate!)
    const newDegU = 3;
    const newDegV = 2;
    const newNbrArrU = [3, 10, 10]; // Trying to add 10 again
    const newNbrArrV = [6, 6]; // Trying to add 6 again

    // Build tree
    const oldHashU = BigInt(computeNbrHash(oldDegU, oldNbrArrU));
    const oldHashV = BigInt(computeNbrHash(oldDegV, oldNbrArrV));

    const tree = new SmtTree(N_LEVELS);
    await tree.init();
    await tree.insert(u, oldHashU);
    await tree.insert(v, oldHashV);

    const oldRoot = await tree.getRoot();
    const siblingsU = ensureSiblingsLength(await tree.getSiblings(u));

    // Update U first
    const newHashU = BigInt(computeNbrHash(newDegU, newNbrArrU));
    await tree.update(u, newHashU);

    const siblingsV = ensureSiblingsLength(await tree.getSiblings(v));

    const input = {
      u: u.toString(),
      v: v.toString(),
      oldDegU: oldDegU.toString(),
      oldDegV: oldDegV.toString(),
      newDegU: newDegU.toString(),
      newDegV: newDegV.toString(),
      oldNbrArrU: padNeighbors(oldNbrArrU),
      oldNbrArrV: padNeighbors(oldNbrArrV),
      newNbrArrU: padNeighbors(newNbrArrU),
      newNbrArrV: padNeighbors(newNbrArrV),
      siblingsU: siblingsU,
      siblingsV: siblingsV,
      oldRoot: F.toString(oldRoot),
    };

    try {
      await circuit.calculateWitness(input, true);
      assert.fail("Should have failed with duplicate edge");
    } catch (error) {
      assert(error.message.includes("Assert Failed"));
    }
  });

  it("should fail if u or v exceeds 2^nLevels", async () => {
    const maxVertexId = Math.pow(2, N_LEVELS) - 1; // Max valid ID is 15 for nLevels=4
    const invalidVertexId = Math.pow(2, N_LEVELS); // 16 - exceeds max (15)
    
    const u = 9; // Valid
    const v = invalidVertexId; // Invalid!
    
    const oldDegU = 0;
    const oldDegV = 0;
    const oldNbrArrU = [];
    const oldNbrArrV = [];

    const newDegU = 1;
    const newDegV = 1;
    const newNbrArrU = [v];
    const newNbrArrV = [u];

    // Build tree with old hashes
    const oldHashU = BigInt(computeNbrHash(oldDegU, oldNbrArrU));
    const oldHashV = BigInt(computeNbrHash(oldDegV, oldNbrArrV));

    const tree = new SmtTree(N_LEVELS);
    await tree.init();
    await tree.insert(u, oldHashU);
    // Note: v is out of bounds, so we can't actually insert it in the tree
    // But we'll provide a dummy hash for testing
    const _ = oldHashV;

    const oldRoot = await tree.getRoot();
    const siblingsU = ensureSiblingsLength(await tree.getSiblings(u));
    // For v, we'll provide dummy siblings
    const siblingsV = Array(N_LEVELS + 1).fill("0");

    const input = {
      u: u.toString(),
      v: v.toString(), // Out of bounds!
      oldDegU: oldDegU.toString(),
      oldDegV: oldDegV.toString(),
      newDegU: newDegU.toString(),
      newDegV: newDegV.toString(),
      oldNbrArrU: padNeighbors(oldNbrArrU),
      oldNbrArrV: padNeighbors(oldNbrArrV),
      newNbrArrU: padNeighbors(newNbrArrU),
      newNbrArrV: padNeighbors(newNbrArrV),
      siblingsU: siblingsU,
      siblingsV: siblingsV,
      oldRoot: F.toString(oldRoot),
    };

    try {
      await circuit.calculateWitness(input, true);
      assert.fail(`Should have failed with v=${invalidVertexId} exceeding max ${maxVertexId}`);
    } catch (error) {
      assert(error.message.includes("Assert Failed"));
    }
  });
});

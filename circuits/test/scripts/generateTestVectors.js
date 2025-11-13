/**
 * generateTestVectors.js
 *
 * Generates test vectors from the computeNodeHash function in nodeHasher.test.js
 * and saves them to a JSON file that can be used to verify the Go implementation.
 *
 * Usage: node generateTestVectors.js
 * Output: ../data/nodeHasherTestVectors.json
 */

import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";
import { buildPoseidon } from "circomlibjs";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const MAX_DEG = 15 * 4; // Maximum degree: 15*4 = 60

// Calculate padLen based on maxDeg
function calculatePadLen(maxDeg) {
  const numR = Math.ceil(maxDeg / 15);
  return 15 * numR;
}

const PAD_LEN = calculatePadLen(MAX_DEG);

// Helper function to compute NbrHash (same as in nodeHasher.test.js)
function computeNbrHash(poseidon, d, neighbors) {
  // Pad neighbors array to PAD_LEN
  const paddedNbrs = [...neighbors];
  while (paddedNbrs.length < PAD_LEN) {
    paddedNbrs.push(0);
  }

  console.log(`Computing NbrHash for degree ${d}`);

  // First block: B_0 = [d, nbr[0..14]] (15 neighbors)
  const firstBlock = [d];
  for (let i = 0; i < 15; i++) {
    firstBlock.push(paddedNbrs[i] || 0);
  }

  // accumulator
  let acc = poseidon.F.toString(poseidon(firstBlock));

  // Continuation blocks (15 neighbors each)
  const numR = Math.ceil(MAX_DEG / 15);

  for (let round = 1; round < numR; round++) {
    const block = [BigInt(acc)];
    const startIdx = 15 + (round - 1) * 15;

    for (let i = 0; i < 15; i++) {
      const idx = startIdx + i;
      block.push(paddedNbrs[idx] || 0);
    }

    acc = poseidon.F.toString(poseidon(block));
  }

  console.log(`Final hash: ${acc}`);
  return acc;
}

async function generateTestVectors() {
  console.log("Initializing Poseidon hasher...");
  const poseidon = await buildPoseidon();
  console.log(`MAX_DEG = ${MAX_DEG}, PAD_LEN = ${PAD_LEN}\n`);

  const testVectors = {
    metadata: {
      maxDeg: MAX_DEG,
      padLen: PAD_LEN,
      description:
        "Test vectors for NbrHash algorithm - Cross-verification between Circom and Go",
      generatedAt: new Date().toISOString(),
      generator: "test/scripts/generateTestVectors.js",
    },
    testCases: [],
  };

  // Test case 1: Degree 0 (no neighbors)
  console.log("\n=== Test Case 1: Degree 0 ===");
  testVectors.testCases.push({
    name: "degree_0_no_neighbors",
    d: 0,
    neighbors: [],
    expectedHash: computeNbrHash(poseidon, 0, []),
  });

  // Test case 2: Degree 1
  console.log("\n=== Test Case 2: Degree 1 ===");
  testVectors.testCases.push({
    name: "degree_1",
    d: 1,
    neighbors: [25],
    expectedHash: computeNbrHash(poseidon, 1, [25]),
  });

  // Test case 3: Degree 5 (within first block)
  console.log("\n=== Test Case 3: Degree 5 ===");
  testVectors.testCases.push({
    name: "degree_5_first_block",
    d: 5,
    neighbors: [1, 3, 8, 12, 15],
    expectedHash: computeNbrHash(poseidon, 5, [1, 3, 8, 12, 15]),
  });

  // Test case 4: Degree 15 (exactly fills first block)
  console.log("\n=== Test Case 4: Degree 15 ===");
  const neighbors15 = [2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30];
  testVectors.testCases.push({
    name: "degree_15_full_first_block",
    d: 15,
    neighbors: neighbors15,
    expectedHash: computeNbrHash(poseidon, 15, neighbors15),
  });

  // Test case 5: Degree 20 (needs 2 blocks)
  console.log("\n=== Test Case 5: Degree 20 ===");
  const neighbors20 = [
    1, 2, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39,
  ];
  testVectors.testCases.push({
    name: "degree_20_two_blocks",
    d: 20,
    neighbors: neighbors20,
    expectedHash: computeNbrHash(poseidon, 20, neighbors20),
  });

  // Test case 6: Degree 30 (boundary: 15*2)
  console.log("\n=== Test Case 6: Degree 30 ===");
  const neighbors30 = Array.from({ length: 30 }, (_, i) => (i + 1) * 10);
  testVectors.testCases.push({
    name: "degree_30_boundary",
    d: 30,
    neighbors: neighbors30,
    expectedHash: computeNbrHash(poseidon, 30, neighbors30),
  });

  // Test case 7: Degree 60 (maximum degree, perfect fit)
  console.log("\n=== Test Case 7: Degree 60 ===");
  const neighbors60 = Array.from({ length: 60 }, (_, i) => i + 1);
  testVectors.testCases.push({
    name: "degree_60_maximum",
    d: 60,
    neighbors: neighbors60,
    expectedHash: computeNbrHash(poseidon, 60, neighbors60),
  });

  // Additional test cases with different patterns
  console.log("\n=== Test Case 8: Large neighbor IDs ===");
  testVectors.testCases.push({
    name: "large_neighbor_ids",
    d: 3,
    neighbors: [1000000, 1000001, 1000002],
    expectedHash: computeNbrHash(poseidon, 3, [1000000, 1000001, 1000002]),
  });

  console.log("\n=== Test Case 9: Degree 45 (15*3) ===");
  const neighbors45 = Array.from({ length: 45 }, (_, i) => i * 2);
  testVectors.testCases.push({
    name: "degree_45_three_blocks",
    d: 45,
    neighbors: neighbors45,
    expectedHash: computeNbrHash(poseidon, 45, neighbors45),
  });

  console.log("\n=== Test Case 10: Sparse neighbors ===");
  testVectors.testCases.push({
    name: "sparse_neighbors",
    d: 7,
    neighbors: [100, 200, 300, 400, 500, 600, 700],
    expectedHash: computeNbrHash(poseidon, 7, [100, 200, 300, 400, 500, 600, 700]),
  });

  // Save to JSON file in data directory
  const outputPath = path.join(
    __dirname,
    "..",
    "data",
    "nodeHasherTestVectors.json",
  );
  fs.writeFileSync(outputPath, JSON.stringify(testVectors, null, 2), "utf8");
  console.log(`\n✓ Test vectors saved to: ${outputPath}`);
  console.log(`✓ Total test cases: ${testVectors.testCases.length}`);
}

// Run the generator
generateTestVectors().catch((err) => {
  console.error("Error generating test vectors:", err);
  process.exit(1);
});

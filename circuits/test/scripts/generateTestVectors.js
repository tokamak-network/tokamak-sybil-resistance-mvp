/**
 * generateTestVectors.js
 *
 * Generates test vectors from the computeNodeHash function in nodeHasher.test.js
 * and saves them to a JSON file that can be used to verify the Go implementation.
 *
 * Usage: node generateTestVectors.js
 * Output: ../data/nodeHasherTestVectors.json
 */

const fs = require("fs");
const path = require("path");
const { buildPoseidon } = require("circomlibjs");

const MAX_DEG = 14 + 15 * 3; // Maximum degree: 14 + 15*3 = 59

// Calculate padLen based on maxDeg
function calculatePadLen(maxDeg) {
  const numR = maxDeg <= 14 ? 0 : Math.ceil((maxDeg - 14) / 15);
  return 14 + 15 * numR;
}

const PAD_LEN = calculatePadLen(MAX_DEG);

// Helper function to compute NodeHash (same as in nodeHasher.test.js)
function computeNodeHash(poseidon, v, d, neighbors) {
  // Pad neighbors array to PAD_LEN
  const paddedNbrs = [...neighbors];
  while (paddedNbrs.length < PAD_LEN) {
    paddedNbrs.push(0);
  }

  console.log(`Computing NodeHash for vertex ${v}, degree ${d}`);

  // First block: [v, d, nbr[0..13]] (14 neighbors)
  const firstBlock = [v, d];
  for (let i = 0; i < 14; i++) {
    firstBlock.push(paddedNbrs[i] || 0);
  }

  // accumulator
  let acc = poseidon.F.toString(poseidon(firstBlock));

  // Continuation blocks (15 neighbors each)
  const numR = MAX_DEG <= 14 ? 0 : Math.ceil((MAX_DEG - 14) / 15);

  for (let round = 0; round < numR; round++) {
    const block = [BigInt(acc)];
    const startIdx = 14 + round * 15;

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
        "Test vectors for NodeHasher algorithm - Cross-verification between Circom and Go",
      generatedAt: new Date().toISOString(),
      generator: "test/scripts/generateTestVectors.js",
    },
    testCases: [],
  };

  // Test case 1: Degree 0 (no neighbors)
  console.log("\n=== Test Case 1: Degree 0 ===");
  testVectors.testCases.push({
    name: "degree_0_no_neighbors",
    v: 42,
    d: 0,
    neighbors: [],
    expectedHash: computeNodeHash(poseidon, 42, 0, []),
  });

  // Test case 2: Degree 1
  console.log("\n=== Test Case 2: Degree 1 ===");
  testVectors.testCases.push({
    name: "degree_1",
    v: 10,
    d: 1,
    neighbors: [25],
    expectedHash: computeNodeHash(poseidon, 10, 1, [25]),
  });

  // Test case 3: Degree 5 (within first block)
  console.log("\n=== Test Case 3: Degree 5 ===");
  testVectors.testCases.push({
    name: "degree_5_first_block",
    v: 7,
    d: 5,
    neighbors: [1, 3, 8, 12, 15],
    expectedHash: computeNodeHash(poseidon, 7, 5, [1, 3, 8, 12, 15]),
  });

  // Test case 4: Degree 14 (exactly fills first block)
  console.log("\n=== Test Case 4: Degree 14 ===");
  const neighbors14 = [2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28];
  testVectors.testCases.push({
    name: "degree_14_full_first_block",
    v: 5,
    d: 14,
    neighbors: neighbors14,
    expectedHash: computeNodeHash(poseidon, 5, 14, neighbors14),
  });

  // Test case 5: Degree 20 (needs 1 continuation block)
  console.log("\n=== Test Case 5: Degree 20 ===");
  const neighbors20 = [
    1, 2, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39,
  ];
  testVectors.testCases.push({
    name: "degree_20_one_continuation",
    v: 99,
    d: 20,
    neighbors: neighbors20,
    expectedHash: computeNodeHash(poseidon, 99, 20, neighbors20),
  });

  // Test case 6: Degree 29 (boundary: 14 + 15)
  console.log("\n=== Test Case 6: Degree 29 ===");
  const neighbors29 = Array.from({ length: 29 }, (_, i) => (i + 1) * 10);
  testVectors.testCases.push({
    name: "degree_29_boundary",
    v: 123,
    d: 29,
    neighbors: neighbors29,
    expectedHash: computeNodeHash(poseidon, 123, 29, neighbors29),
  });

  // Test case 7: Degree 59 (maximum degree, perfect fit)
  console.log("\n=== Test Case 7: Degree 59 ===");
  const neighbors59 = Array.from({ length: 59 }, (_, i) => i + 1);
  testVectors.testCases.push({
    name: "degree_59_maximum",
    v: 0,
    d: 59,
    neighbors: neighbors59,
    expectedHash: computeNodeHash(poseidon, 0, 59, neighbors59),
  });

  // Additional test cases with different patterns
  console.log("\n=== Test Case 8: Large vertex ID ===");
  testVectors.testCases.push({
    name: "large_vertex_id",
    v: 999999,
    d: 3,
    neighbors: [1000000, 1000001, 1000002],
    expectedHash: computeNodeHash(
      poseidon,
      999999,
      3,
      [1000000, 1000001, 1000002],
    ),
  });

  console.log("\n=== Test Case 9: Degree 44 (14 + 15*2) ===");
  const neighbors44 = Array.from({ length: 44 }, (_, i) => i * 2);
  testVectors.testCases.push({
    name: "degree_44_two_continuations",
    v: 777,
    d: 44,
    neighbors: neighbors44,
    expectedHash: computeNodeHash(poseidon, 777, 44, neighbors44),
  });

  console.log("\n=== Test Case 10: Sparse neighbors ===");
  testVectors.testCases.push({
    name: "sparse_neighbors",
    v: 50,
    d: 7,
    neighbors: [100, 200, 300, 400, 500, 600, 700],
    expectedHash: computeNodeHash(
      poseidon,
      50,
      7,
      [100, 200, 300, 400, 500, 600, 700],
    ),
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

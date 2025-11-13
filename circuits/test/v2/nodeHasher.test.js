import fs from "fs";
import path from "path";
import { describe, it, before, after } from "mocha";
import assert from "assert";
import { wasm as tester } from "circom_tester";
import { buildPoseidon } from "circomlibjs";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

describe("NodeHasher circuit test", function () {
  this.timeout(200000);

  const MAX_DEG = 15 * 4; // Maximum degree: 15*4 = 60 (numR=4)
  let circuit;
  let circuitTmpPath;
  let poseidon;

  // Calculate padLen based on maxDeg
  function calculatePadLen(maxDeg) {
    const numR = Math.ceil(maxDeg / 15);
    return 15 * numR;
  }

  const PAD_LEN = calculatePadLen(MAX_DEG);

  before(async () => {
    // Initialize Poseidon hasher
    poseidon = await buildPoseidon();

    // Create circuit with MAX_DEG = 59 (14 + 15*3)
    const circuitSrc = `
            pragma circom 2.0.0;
            include "../circuits/syb_rollup_v2/node_hasher.circom";
            component main = NodeHasher(${MAX_DEG});
        `;
    circuitTmpPath = path.join(__dirname, "node-hasher-temp.circom");
    fs.writeFileSync(circuitTmpPath, circuitSrc, "utf8");

    circuit = await tester(circuitTmpPath, {
      reduceConstraints: false,
      include: path.join(__dirname, "../"),
    });
    await circuit.loadConstraints();
    console.log(`\n✓ NodeHasher circuit compiled with maxDeg=${MAX_DEG}`);
    console.log(`✓ padLen=${PAD_LEN} (15*numR)`);
    console.log(`✓ Constraints: ${circuit.constraints.length}`);

    // Calculate expected rounds
    const numR = Math.ceil(MAX_DEG / 15);
    console.log(
      `✓ Total rounds: ${numR} (each processing 15 neighbors)\n`,
    );
  });

  after(() => {
    if (fs.existsSync(circuitTmpPath)) {
      fs.unlinkSync(circuitTmpPath);
    }
  });

  // Helper function to compute NbrHash according to the NEW spec
  function computeNodeHash(d, neighbors) {
    // Pad neighbors array to PAD_LEN
    const paddedNbrs = [...neighbors];
    while (paddedNbrs.length < PAD_LEN) {
      paddedNbrs.push(0);
    }

    // console.log(`  Computing NbrHash for degree ${d}`);
    // console.log(
    //   `  Neighbors: [${neighbors.join(", ")}]${neighbors.length < PAD_LEN ? " + padding" : ""}`,
    // );

    // First block: B_0 = [d, nbr[0..14]] (15 neighbors)
    const firstBlock = [d];
    for (let i = 0; i < 15; i++) {
      firstBlock.push(paddedNbrs[i] || 0);
    }

    // accumulator
    let acc = poseidon.F.toString(poseidon(firstBlock));
    // console.log(`  Block 0 hash: ${acc.slice(0, 20)}...`);

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
      // console.log(`  Block ${round} hash: ${acc.slice(0, 20)}...`);
    }

    // console.log(`  Final hash: ${acc}`);
    return acc;
  }

  /**
   * VALID
   * [X] hash a vertex with degree 0
   * [X] hash a vertex with degree 1
   * [X] hash a vertex with degree 5
   * [X] hash a vertex with degree 15 (exactly fills first block)
   * [X] hash a vertex with degree 20
   * [X] hash a vertex with degree 60 (current maximum degree)
   * [X] hash a vertex with degree 30 (15*2)
   * INVALID
   * [X] fail with nbr_arr not ascending
   * [X] fail with degree 61 (exceed padLength)
   *
   */

  it("should hash a vertex with degree 0 (no neighbors)", async () => {
    const d = 0;
    const neighbors = [];

    const expectedHash = computeNodeHash(d, neighbors);

    // Prepare input (pad with zeros)
    const input = {
      d: d.toString(),
      nbr_arr: Array(PAD_LEN).fill("0"),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
  });

  it("should hash a vertex with degree 1", async () => {
    const d = 1;
    const neighbors = [25];

    const expectedHash = computeNodeHash(d, neighbors);

    const paddedNbrs = [
      ...neighbors,
      ...Array(PAD_LEN - neighbors.length).fill(0),
    ];
    const input = {
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
  });

  it("should hash a vertex with degree 5 (within first block)", async () => {
    const d = 5;
    const neighbors = [1, 3, 8, 12, 15]; // Sorted

    const expectedHash = computeNodeHash(d, neighbors);

    const paddedNbrs = [
      ...neighbors,
      ...Array(PAD_LEN - neighbors.length).fill(0),
    ];
    const input = {
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
  });

  it("should hash a vertex with degree 15 (exactly fills first block)", async () => {
    const d = 15;
    const neighbors = [2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30];

    const expectedHash = computeNodeHash(d, neighbors);

    const paddedNbrs = [
      ...neighbors,
      ...Array(PAD_LEN - neighbors.length).fill(0),
    ];
    const input = {
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
  });

  it("should hash a vertex with degree 20 (needs 2 blocks)", async () => {
    const d = 20;
    // 20 neighbors: sorted array
    const neighbors = [
      1, 2, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39,
    ];

    const expectedHash = computeNodeHash(d, neighbors);

    const paddedNbrs = [
      ...neighbors,
      ...Array(PAD_LEN - neighbors.length).fill(0),
    ];
    const input = {
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
  });

  it("should hash a vertex with degree 60 (maximum degree, perfect fit)", async () => {
    const d = 60;
    // Generate 60 neighbors: [1, 2, 3, ..., 60]
    const neighbors = Array.from({ length: 60 }, (_, i) => i + 1);

    const expectedHash = computeNodeHash(d, neighbors);

    const input = {
      d: d.toString(),
      nbr_arr: neighbors.map((x) => x.toString()),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
  });

  it("should hash a vertex with degree 30 (boundary: 15*2)", async () => {
    const d = 30;
    // 30 neighbors: multiples of 10
    const neighbors = Array.from({ length: 30 }, (_, i) => (i + 1) * 10);

    const expectedHash = computeNodeHash(d, neighbors);

    const paddedNbrs = [
      ...neighbors,
      ...Array(PAD_LEN - neighbors.length).fill(0),
    ];
    const input = {
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
  });

  it("should fail with nbr_arr is not ascending", async () => {
    const d = 3;
    const neighbors = [1, 81, 3];

    const paddedNbrs = [
      ...neighbors,
      ...Array(PAD_LEN - neighbors.length).fill(0),
    ];

    const input = {
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    try {
      await circuit.calculateWitness(input, true);
    } catch (error) {
      assert(error.message.includes("Assert Failed."));
    }
  });

  it("should fail with a vertex with degree 61", async () => {
    const d = MAX_DEG + 1;
    const neighbors = Array.from({ length: d }, (_, i) => (i + 1) * 10);

    const paddedNbrs = [...neighbors];

    const input = {
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    try {
      await circuit.calculateWitness(input, true);
    } catch (error) {
      assert(
        error.message.includes("Too many values for input signal nbr_arr"),
      );
    }
  });
});

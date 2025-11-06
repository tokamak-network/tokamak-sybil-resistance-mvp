const fs = require("fs");
const path = require("path");
const { describe, it, before, after } = require("mocha");
const { strict: assert } = require("assert");
const { wasm: tester } = require("circom_tester");
const { buildPoseidon } = require("circomlibjs");
const { log } = require("console");

describe("NodeHasher circuit test", function () {
  this.timeout(200000);

  const MAX_DEG = 14 + 15 * 3; // Maximum degree: 14 + 15*3 (k=3, no padding in last block)
  let circuit;
  let circuitTmpPath;
  let poseidon;

  // Calculate padLen based on maxDeg
  function calculatePadLen(maxDeg) {
    const numR = maxDeg <= 14 ? 0 : Math.ceil((maxDeg - 14) / 15);
    return 14 + 15 * numR;
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
    console.log(`✓ padLen=${PAD_LEN} (14 + 15*numR)`);
    console.log(`✓ Constraints: ${circuit.constraints.length}`);

    // Calculate expected rounds
    const numR = MAX_DEG <= 14 ? 0 : Math.ceil((MAX_DEG - 14) / 15);
    console.log(
      `✓ Continuation rounds: ${numR} (1 first block + ${numR} continuation blocks)\n`,
    );
  });

  after(() => {
    if (fs.existsSync(circuitTmpPath)) {
      fs.unlinkSync(circuitTmpPath);
    }
  });

  // Helper function to compute NodeHash according to the NEW spec
  function computeNodeHash(v, d, neighbors) {
    // Pad neighbors array to PAD_LEN
    const paddedNbrs = [...neighbors];
    while (paddedNbrs.length < PAD_LEN) {
      paddedNbrs.push(0);
    }

    console.log(`  Computing NodeHash for vertex ${v}, degree ${d}`);
    console.log(
      `  Neighbors: [${neighbors.join(", ")}]${neighbors.length < PAD_LEN ? " + padding" : ""}`,
    );

    // First block: [v, d, nbr[0..13]] (14 neighbors)
    const firstBlock = [v, d];
    for (let i = 0; i < 14; i++) {
      firstBlock.push(paddedNbrs[i] || 0);
    }

    // accumulator
    let acc = poseidon.F.toString(poseidon(firstBlock));
    console.log(`  First block hash: ${acc.slice(0, 20)}...`);

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
      console.log(`  Round ${round + 1} hash: ${acc.slice(0, 20)}...`);
    }

    console.log(`  Final hash: ${acc}`);
    return acc;
  }

  /**
   * VALID
   * [X] hash a vertex with degree 0
   * [X] hash a vertex with degree 1
   * [X] hash a vertex with degree 5
   * [X] hash a vertex with degree 14
   * [X] hash a vertex with degree 20
   * [X] hash a vertex with degree 59 (current maximum degree)
   * [X] hash a vertex with degree 29 (14+15)
   * INVALID
   * [X] fail with nbr_arr not ascending
   * [X] fail with degree 60 (exceed padLength)
   *
   */

  it("should hash a vertex with degree 0 (no neighbors)", async () => {
    const v = 42;
    const d = 0;
    const neighbors = [];

    const expectedHash = computeNodeHash(v, d, neighbors);

    // Prepare input (pad with zeros)
    const input = {
      v: v.toString(),
      d: d.toString(),
      nbr_arr: Array(PAD_LEN).fill("0"),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();
    console.log(`  Circuit output: ${circuitOutput}`);

    assert.equal(circuitOutput, expectedHash);
    console.log("  ✓ Hash verified for degree 0");
  });

  it("should hash a vertex with degree 1", async () => {
    const v = 10;
    const d = 1;
    const neighbors = [25];

    const expectedHash = computeNodeHash(v, d, neighbors);

    const paddedNbrs = [
      ...neighbors,
      ...Array(PAD_LEN - neighbors.length).fill(0),
    ];
    const input = {
      v: v.toString(),
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
    console.log("  ✓ Hash verified for degree 1");
  });

  it("should hash a vertex with degree 5 (within first block)", async () => {
    const v = 7;
    const d = 5;
    const neighbors = [1, 3, 8, 12, 15]; // Sorted

    const expectedHash = computeNodeHash(v, d, neighbors);

    const paddedNbrs = [
      ...neighbors,
      ...Array(PAD_LEN - neighbors.length).fill(0),
    ];
    const input = {
      v: v.toString(),
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
    console.log("  ✓ Hash verified for degree 5");
  });

  it("should hash a vertex with degree 14 (exactly fills first block)", async () => {
    const v = 5;
    const d = 14;
    const neighbors = [2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28];

    const expectedHash = computeNodeHash(v, d, neighbors);

    const paddedNbrs = [
      ...neighbors,
      ...Array(PAD_LEN - neighbors.length).fill(0),
    ];
    const input = {
      v: v.toString(),
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
    console.log("  ✓ Hash verified for degree 14");
  });

  it("should hash a vertex with degree 20 (needs 1 continuation block)", async () => {
    const v = 99;
    const d = 20;
    // 20 neighbors: sorted array
    const neighbors = [
      1, 2, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39,
    ];

    const expectedHash = computeNodeHash(v, d, neighbors);

    const paddedNbrs = [
      ...neighbors,
      ...Array(PAD_LEN - neighbors.length).fill(0),
    ];
    const input = {
      v: v.toString(),
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
    console.log("  ✓ Hash verified for degree 20");
  });

  it("should hash a vertex with degree 59 (maximum degree, perfect fit)", async () => {
    const v = 0;
    const d = 59;
    // Generate 59 neighbors: [1, 2, 3, ..., 59]
    const neighbors = Array.from({ length: 59 }, (_, i) => i + 1);

    const expectedHash = computeNodeHash(v, d, neighbors);

    const input = {
      v: v.toString(),
      d: d.toString(),
      nbr_arr: neighbors.map((x) => x.toString()),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
    console.log("  ✓ Hash verified for degree 59 (perfect block fit)");
  });

  it("should hash a vertex with degree 29 (boundary: 14 + 15)", async () => {
    const v = 123;
    const d = 29;
    // 29 neighbors: multiples of 10
    const neighbors = Array.from({ length: 29 }, (_, i) => (i + 1) * 10);

    const expectedHash = computeNodeHash(v, d, neighbors);

    const paddedNbrs = [
      ...neighbors,
      ...Array(PAD_LEN - neighbors.length).fill(0),
    ];
    const input = {
      v: v.toString(),
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);

    const circuitOutput = w[1].toString();

    assert.equal(circuitOutput, expectedHash);
    console.log("  ✓ Hash verified for degree 29");
  });

  it("should fail with nbr_arr is not ascending", async () => {
    const v = 123;
    const d = 3;
    const neighbors = [1, 81, 3];

    const paddedNbrs = [
      ...neighbors,
      ...Array(PAD_LEN - neighbors.length).fill(0),
    ];

    const input = {
      v: v.toString(),
      d: d.toString(),
      nbr_arr: paddedNbrs.map((x) => x.toString()),
    };

    try {
      await circuit.calculateWitness(input, true);
    } catch (error) {
      assert(error.message.includes("Assert Failed."));
    }
  });

  it("should fail with a vertex with degree 60", async () => {
    const v = 123;
    const d = MAX_DEG + 1;
    const neighbors = Array.from({ length: d }, (_, i) => (i + 1) * 10);

    const paddedNbrs = [...neighbors];

    const input = {
      v: v.toString(),
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

import fs from "fs";
import path from "path";
import { describe, it, before, after } from "mocha";
import assert from "assert";
import { wasm as tester } from "circom_tester";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

describe("IsInArray template test", function () {
  this.timeout(100000);

  const ARRAY_SIZE = 10;
  let circuit;
  let circuitTmpPath;

  before(async () => {
    // Create a test circuit for IsInArray
    const circuitSrc = `
            pragma circom 2.0.0;
            include "../circuits/syb_rollup_v2/lib/is_in_array.circom";
            component main = IsInArray(${ARRAY_SIZE});
        `;
    circuitTmpPath = path.join(__dirname, "is-in-array-temp.circom");
    fs.writeFileSync(circuitTmpPath, circuitSrc, "utf8");

    circuit = await tester(circuitTmpPath, {
      reduceConstraints: false,
      include: path.join(__dirname, "../../"),
    });
    await circuit.loadConstraints();
    console.log(`\n✓ IsInArray circuit compiled (array size=${ARRAY_SIZE})`);
    console.log(`✓ Constraints: ${circuit.constraints.length}\n`);
  });

  after(() => {
    if (fs.existsSync(circuitTmpPath)) {
      fs.unlinkSync(circuitTmpPath);
    }
  });

  it("should pass when target is not in array", async () => {
    const arr = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19];
    const target = 8; // Not in array

    const input = {
      arr: arr.map((x) => x.toString()),
      target: target.toString(),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);
    
    // Check output: out should be 0 (target not in array)
    const out = w[1];
    assert.equal(out.toString(), "0");
  });

  it("should pass when array contains zeros and target is non-zero", async () => {
    const arr = [1, 3, 0, 0, 0, 0, 0, 0, 0, 0]; // Padded with zeros
    const target = 5; // Not in array

    const input = {
      arr: arr.map((x) => x.toString()),
      target: target.toString(),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);
    
    // Check output: out should be 0 (target not in array)
    const out = w[1];
    assert.equal(out.toString(), "0");
  });

  it("should return 1 when target is in array (beginning)", async () => {
    const arr = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19];
    const target = 1; // At the beginning

    const input = {
      arr: arr.map((x) => x.toString()),
      target: target.toString(),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);
    
    // Check output: out should be 1 (target found once)
    const out = w[1];
    assert.equal(out.toString(), "1");
  });

  it("should return 1 when target is in array (middle)", async () => {
    const arr = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19];
    const target = 9; // In the middle

    const input = {
      arr: arr.map((x) => x.toString()),
      target: target.toString(),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);
    
    // Check output: out should be 1 (target found once)
    const out = w[1];
    assert.equal(out.toString(), "1");
  });

  it("should return 1 when target is in array (end)", async () => {
    const arr = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19];
    const target = 19; // At the end

    const input = {
      arr: arr.map((x) => x.toString()),
      target: target.toString(),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);
    
    // Check output: out should be 1 (target found once)
    const out = w[1];
    assert.equal(out.toString(), "1");
  });

  // This case is solved by NodeHasher's strictly ascending check. so no need to check here.
  it("should return count when target appears multiple times in array", async () => {
    const arr = [1, 5, 5, 7, 5, 11, 13, 5, 17, 19];
    const target = 5; // Appears 4 times

    const input = {
      arr: arr.map((x) => x.toString()),
      target: target.toString(),
    };

    const w = await circuit.calculateWitness(input, true);
    await circuit.checkConstraints(w);
    
    // Check output: out should be 4 (target found 4 times)
    const out = w[1];
    assert.equal(out.toString(), "4");
  });
});


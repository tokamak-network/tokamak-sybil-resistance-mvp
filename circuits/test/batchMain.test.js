import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";
import { describe, it, before, after } from "mocha";
import { strict as assert } from "assert";
import { wasm as tester } from "circom_tester";
import { SmtTree } from "./utils/smt.js";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const N_LEVELS = 16;
const N_TX = 1;

describe("BatchMain advanced input test (N_TX=1, N_LEVELS=16)", function () {
    this.timeout(100000);

    let circuit;
    let circuitTmpPath;
    let smt;
    let Fr;

    // Helper function to encode transaction data
    // txData format: txnType (8 bits) | fromIdx (nLevels bits) | toIdx (nLevels bits) | amount (128 bits)
    function encodeTxData(txnType, fromIdx, toIdx, amount) {
        const txnTypeBig = BigInt(txnType);
        const fromIdxBig = BigInt(fromIdx);
        const toIdxBig = BigInt(toIdx);
        const amountBig = BigInt(amount);

        // Pack: amount << (8 + nLevels + nLevels) | toIdx << (8 + nLevels) | fromIdx << 8 | txnType
        const packed = (amountBig << BigInt(8 + N_LEVELS + N_LEVELS)) |
                       (toIdxBig << BigInt(8 + N_LEVELS)) |
                       (fromIdxBig << BigInt(8)) |
                       txnTypeBig;
        return packed.toString();
    }

    function getBaseInput() {
        return {
            // public
            explodeAmount: "1",

            // private
            oldLastIdx: "1",
            oldAccountRoot: Fr.toString(Fr.zero),
            oldVouchRoot: Fr.toString(Fr.zero),
            oldScoreRoot: Fr.toString(Fr.zero),
            newScoreRoot: Fr.toString(Fr.zero),

            // tx signals
            txData: ["0"],

            // account state 1
            balance1: ["0"],
            ethAddr1: ["0"],
            siblings1: [smt.getEmptySiblings(N_LEVELS + 1)],
            isOld0_1: ["1"],
            oldKey1: ["0"],
            oldValue1: ["0"],

            // account state 2
            balance2: ["0"],
            ethAddr2: ["0"],
            siblings2: [smt.getEmptySiblings(N_LEVELS + 1)],
            isOld0_2: ["1"],
            oldKey2: ["0"],
            oldValue2: ["0"],

            // vouch states
            siblings3: [smt.getEmptySiblings(2 * N_LEVELS + 1)],
            isOld0_3: ["1"],
            oldKey3: ["0"],
            oldValue3: ["0"],
        };
    }

    before(async () => {
        smt = new SmtTree(N_LEVELS);
        await smt.init();
        Fr = smt.Fr;

        const circuitSrc = `
            pragma circom 2.0.0;
            include "../circuits/syb_rollup/batch-main.circom";
            component main{public [explodeAmount]} = BatchMain(${N_TX}, ${N_LEVELS});
        `;
        circuitTmpPath = path.join(__dirname, "batch-main-temp.circom");
        fs.writeFileSync(circuitTmpPath, circuitSrc, "utf8");

        circuit = await tester(circuitTmpPath, {
            reduceConstraints: false,
            include: path.join(__dirname, "../circuits")
        });
        await circuit.loadConstraints();
        console.log("Constraints:", circuit.constraints.length);
    });

    after(() => {
        if (fs.existsSync(circuitTmpPath)) {
            fs.unlinkSync(circuitTmpPath);
        }
    });

    it("createAccountDeposit", async () => {
        const input = getBaseInput();

        // Tx data: createAccountDeposit (txnType=0, fromIdx=0, toIdx=0, amount=500)
        input.txData[0] = encodeTxData(0, 0, 0, 500);

        const w = await circuit.calculateWitness(input, true);
        await circuit.checkConstraints(w);
        console.log("createAccountDeposit => Success, hashGlobalInputs=", w[1].toString());
    });

});
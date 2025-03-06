const fs = require("fs");
const path = require("path");
const { describe, it, before, after } = require("mocha");
const { strict: assert } = require("assert");
const { wasm: tester } = require("circom_tester");
const { newMemEmptyTrie } = require("circomlibjs");
const { Scalar, F } = require("ffjavascript");

/**
 * empty siblings array for empty SMT tree.
 * use 0 for all levels.
 */
function getEmptySiblings(n, F) {
    return Array(n).fill(F.toString(F.zero));
}

describe("BatchMain advanced input test (nTx=1, nLevels=16)", function () {
    this.timeout(100000);

    let circuit;
    const nTx = 1;
    const nLevels = 16;
    let circuitTmpPath;
    
    // SMT related variables
    let smt;
    let F;

    function getBaseInput() {
        return {
            // public
            EXPLODE_AMOUNT: "1",

            // private
            oldLastIdx: "1",
            oldAccountRoot: F.toString(F.zero),
            oldVouchRoot: F.toString(F.zero),

            // tx signals
            txCompressedData: ["0"],
            fromIdx: ["0"],
            auxFromIdx: ["0"],
            toIdx: ["0"],
            amountF: ["0"],
            loadAmountF: ["0"],
            fromEthAddr: ["0"],
            toEthAddr: ["0"],
            newAccount: ["0"],

            // account state 1
            balance1: ["0"],
            ethAddr1: ["0"],
            siblings1: [getEmptySiblings(nLevels + 1, F)],
            isOld0_1: ["1"],
            oldKey1: ["0"],
            oldValue1: ["0"],

            // account state 2
            balance2: ["0"],
            ethAddr2: ["0"],
            siblings2: [getEmptySiblings(nLevels + 1, F)],
            newExit: ["0"],
            isOld0_2: ["1"],
            oldKey2: ["0"],
            oldValue2: ["0"],

            // vouch states
            siblings3: [getEmptySiblings(2 * nLevels + 1, F)],
            isOld0_3: ["1"],
            oldKey3: ["0"],
            oldValue3: ["0"],

            siblings4: [getEmptySiblings(2 * nLevels + 1, F)],
            isOld0_4: ["1"],
            oldKey4: ["0"],
            oldValue4: ["0"],
        };
    }

    before(async () => {
        smt = await newMemEmptyTrie();
        F = smt.F;

        const circuitSrc = `
            pragma circom 2.0.0;
            include "../circuits/syb_rollup/batch-main.circom";
            component main{public [EXPLODE_AMOUNT]} = BatchMain(${nTx}, ${nLevels});
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
        // Tx data: createAccountDeposit
        input.txCompressedData[0] = "0";
        input.fromIdx[0] = "0";
        input.toIdx[0] = "0";
        input.newAccount[0] = "1";
        input.auxFromIdx[0] = "2";
        input.loadAmountF[0] = "500";
        input.amountF[0] = "0";
        input.fromEthAddr[0] = "123";
        input.toEthAddr[0] = "0";

        const w = await circuit.calculateWitness(input, true);
        await circuit.checkConstraints(w);
        console.log("createAccountDeposit => Success, hashGlobalInputs=", w[1].toString());
    });

});
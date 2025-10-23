/**
 * DecodeTx Circuit Unit Tests
 *
 * Tests the DecodeTx circuit which decodes compressed transaction data.
 * Format: txnType (8 bits) | fromIdx (nLevels bits) | toIdx (nLevels bits) | amount (128 bits)
 *
 * Circuit: circuits/syb_rollup/decode-tx.circom
 * Template: DecodeTx(nLevels)
 * Dependencies: circomlib/bitify, circomlib/comparators
 */

import { describe, it, before, after } from "mocha";
import { strict as assert } from "assert";
import { wasm as tester } from "circom_tester";
import path from "path";
import { fileURLToPath } from "url";
import fs from "fs";
import { compressTxData } from "../utils/testUtils.js";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const N_LEVELS = 16; // Standard merkle tree depth for testing

describe("DecodeTx Circuit:(Transaction)", function() {
    this.timeout(120000); // 120 second timeout for circuit compilation

    let circuit;
    let circuitTmpPath;

    before(async () => {
        const circuitSrc = `
            pragma circom 2.0.0;
            include "../../circuits/syb_rollup/decode-tx.circom";
            component main = DecodeTx(${N_LEVELS});
        `;
        circuitTmpPath = path.join(__dirname, "decode-tx-temp.circom");
        fs.writeFileSync(circuitTmpPath, circuitSrc, "utf8");

        circuit = await tester(circuitTmpPath, {
            reduceConstraints: false,
            include: path.join(__dirname, "../..")
        });
        await circuit.loadSymbols();
    });

    after(() => {
        if (fs.existsSync(circuitTmpPath)) {
            fs.unlinkSync(circuitTmpPath);
        }
    });

    // Helper to get outputs from witness
    function getOutputs(witness) {
        return {
            txnType: Number(witness[circuit.symbols["main.txnType"].varIdx]),
            fromIdx: Number(witness[circuit.symbols["main.fromIdx"].varIdx]),
            toIdx: Number(witness[circuit.symbols["main.toIdx"].varIdx]),
            amount: witness[circuit.symbols["main.amount"].varIdx].toString()
        };
    }

    // Helper to convert txData to expected bit array (LSB first)
    function txDataToBits(txData) {
        const txDataBitsLength = 8 + N_LEVELS + N_LEVELS + 128;
        const bits = [];
        const txDataBig = BigInt(txData);

        for (let i = 0; i < txDataBitsLength; i++) {
            bits.push(Number((txDataBig >> BigInt(i)) & 1n));
        }
        return bits;
    }

    // Helper to extract bitsTxData from witness
    function getBitsTxDataFromWitness(witness) {
        const txDataBitsLength = 8 + N_LEVELS + N_LEVELS + 128;
        const bits = [];
        for (let i = 0; i < txDataBitsLength; i++) {
            bits.push(Number(witness[circuit.symbols[`main.bitsTxData[${i}]`].varIdx]));
        }
        return bits;
    }

    describe("Valid Cases - Transaction Decoding", () => {
        it("should decode NOP transaction (type 0)", async () => {
            const txData = compressTxData(0, 0, 0, 0, N_LEVELS);
            const witness = await circuit.calculateWitness({ txData }, true);
            await circuit.checkConstraints(witness);

            const output = getOutputs(witness);
            assert.strictEqual(output.txnType, 0);
            assert.strictEqual(output.fromIdx, 0);
            assert.strictEqual(output.toIdx, 0);
            assert.strictEqual(output.amount, "0");

            const expectedBits = txDataToBits(txData);
            const actualBits = getBitsTxDataFromWitness(witness);
            assert.deepStrictEqual(actualBits, expectedBits, "bitsTxData should match expected bit representation");
        });

        it("should decode CreateAccount transaction (type 1)", async () => {
            const txData = compressTxData(1, 0, 5, 1000, N_LEVELS);
            const witness = await circuit.calculateWitness({ txData }, true);
            await circuit.checkConstraints(witness);

            const output = getOutputs(witness);
            assert.strictEqual(output.txnType, 1);
            assert.strictEqual(output.fromIdx, 0);
            assert.strictEqual(output.toIdx, 5);
            assert.strictEqual(output.amount, "1000");

            const expectedBits = txDataToBits(txData);
            const actualBits = getBitsTxDataFromWitness(witness);
            assert.deepStrictEqual(actualBits, expectedBits, "bitsTxData should match expected bit representation");
        });

        it("should decode Deposit transaction (type 2)", async () => {
            const txData = compressTxData(2, 10, 10, 5000, N_LEVELS);
            const witness = await circuit.calculateWitness({ txData }, true);
            await circuit.checkConstraints(witness);

            const output = getOutputs(witness);
            assert.strictEqual(output.txnType, 2);
            assert.strictEqual(output.fromIdx, 10);
            assert.strictEqual(output.toIdx, 10);
            assert.strictEqual(output.amount, "5000");

            const expectedBits = txDataToBits(txData);
            const actualBits = getBitsTxDataFromWitness(witness);
            assert.deepStrictEqual(actualBits, expectedBits, "bitsTxData should match expected bit representation");
        });

        it("should decode Withdraw transaction (type 3)", async () => {
            const txData = compressTxData(3, 15, 0, 3000, N_LEVELS);
            const witness = await circuit.calculateWitness({ txData }, true);
            await circuit.checkConstraints(witness);

            const output = getOutputs(witness);
            assert.strictEqual(output.txnType, 3);
            assert.strictEqual(output.fromIdx, 15);
            assert.strictEqual(output.toIdx, 0);
            assert.strictEqual(output.amount, "3000");

            const expectedBits = txDataToBits(txData);
            const actualBits = getBitsTxDataFromWitness(witness);
            assert.deepStrictEqual(actualBits, expectedBits, "bitsTxData should match expected bit representation");
        });

        it("should decode Vouch transaction (type 4)", async () => {
            const txData = compressTxData(4, 5, 8, 0, N_LEVELS);
            const witness = await circuit.calculateWitness({ txData }, true);
            await circuit.checkConstraints(witness);

            const output = getOutputs(witness);
            assert.strictEqual(output.txnType, 4);
            assert.strictEqual(output.fromIdx, 5);
            assert.strictEqual(output.toIdx, 8);
            assert.strictEqual(output.amount, "0");

            const expectedBits = txDataToBits(txData);
            const actualBits = getBitsTxDataFromWitness(witness);
            assert.deepStrictEqual(actualBits, expectedBits, "bitsTxData should match expected bit representation");
        });

        it("should decode UnVouch transaction (type 5)", async () => {
            const txData = compressTxData(5, 5, 8, 0, N_LEVELS);
            const witness = await circuit.calculateWitness({ txData }, true);
            await circuit.checkConstraints(witness);

            const output = getOutputs(witness);
            assert.strictEqual(output.txnType, 5);
            assert.strictEqual(output.fromIdx, 5);
            assert.strictEqual(output.toIdx, 8);
            assert.strictEqual(output.amount, "0");

            const expectedBits = txDataToBits(txData);
            const actualBits = getBitsTxDataFromWitness(witness);
            assert.deepStrictEqual(actualBits, expectedBits, "bitsTxData should match expected bit representation");
        });

        it("should decode Explode transaction (type 6)", async () => {
            const txData = compressTxData(6, 20, 25, 2000, N_LEVELS);
            const witness = await circuit.calculateWitness({ txData }, true);
            await circuit.checkConstraints(witness);

            const output = getOutputs(witness);
            assert.strictEqual(output.txnType, 6);
            assert.strictEqual(output.fromIdx, 20);
            assert.strictEqual(output.toIdx, 25);
            assert.strictEqual(output.amount, "2000");

            const expectedBits = txDataToBits(txData);
            const actualBits = getBitsTxDataFromWitness(witness);
            assert.deepStrictEqual(actualBits, expectedBits, "bitsTxData should match expected bit representation");
        });

        it("should decode transaction with max values", async () => {
            const maxIdx = (1 << N_LEVELS) - 1; // 65535
            const maxAmount = (1n << 128n) - 1n;
            const txData = compressTxData(6, maxIdx, maxIdx, maxAmount.toString(), N_LEVELS);

            const witness = await circuit.calculateWitness({ txData }, true);
            await circuit.checkConstraints(witness);

            const output = getOutputs(witness);
            assert.strictEqual(output.txnType, 6);
            assert.strictEqual(output.fromIdx, maxIdx);
            assert.strictEqual(output.toIdx, maxIdx);
            assert.strictEqual(BigInt(output.amount), maxAmount);

            const expectedBits = txDataToBits(txData);
            const actualBits = getBitsTxDataFromWitness(witness);
            assert.deepStrictEqual(actualBits, expectedBits, "bitsTxData should match expected bit representation");
        });

        it("should decode transaction with minimum non-zero values", async () => {
            const txData = compressTxData(1, 1, 1, 1, N_LEVELS);
            const witness = await circuit.calculateWitness({ txData }, true);
            await circuit.checkConstraints(witness);

            const output = getOutputs(witness);
            assert.strictEqual(output.txnType, 1);
            assert.strictEqual(output.fromIdx, 1);
            assert.strictEqual(output.toIdx, 1);
            assert.strictEqual(output.amount, "1");

            const expectedBits = txDataToBits(txData);
            const actualBits = getBitsTxDataFromWitness(witness);
            assert.deepStrictEqual(actualBits, expectedBits, "bitsTxData should match expected bit representation");
        });

        it(" should decode transaction with large 128-bit amount", async () => {
            const largeAmount = "123456789012345678901234567890";
            const txData = compressTxData(2, 100, 200, largeAmount, N_LEVELS);

            const witness = await circuit.calculateWitness({ txData }, true);
            await circuit.checkConstraints(witness);

            const output = getOutputs(witness);
            assert.strictEqual(output.txnType, 2);
            assert.strictEqual(output.fromIdx, 100);
            assert.strictEqual(output.toIdx, 200);
            assert.strictEqual(output.amount, largeAmount);

            const expectedBits = txDataToBits(txData);
            const actualBits = getBitsTxDataFromWitness(witness);
            assert.deepStrictEqual(actualBits, expectedBits, "bitsTxData should match expected bit representation");
        });
    });

    describe("Invalid Cases - Constraint Failures", () => {
        it(" should fail for invalid txnType = 7", async () => {
            const txData = compressTxData(7, 0, 0, 0, N_LEVELS);

            try {
                await circuit.calculateWitness({ txData }, true);
                assert.fail("Should have thrown constraint error for invalid txnType");
            } catch (error) {
                assert(error.message.includes("Error") || error.message.includes("Assert"),
                    "Should fail with constraint error");
            }
        });

        it(" should fail for invalid txnType = 8", async () => {
            const txData = compressTxData(8, 0, 0, 0, N_LEVELS);

            try {
                await circuit.calculateWitness({ txData }, true);
                assert.fail("Should have thrown constraint error for invalid txnType");
            } catch (error) {
                assert(error.message.includes("Error") || error.message.includes("Assert"),
                    "Should fail with constraint error");
            }
        });

        it(" should fail for invalid txnType = 255", async () => {
            const txData = compressTxData(255, 0, 0, 0, N_LEVELS);

            try {
                await circuit.calculateWitness({ txData }, true);
                assert.fail("Should have thrown constraint error for invalid txnType");
            } catch (error) {
                assert(error.message.includes("Error") || error.message.includes("Assert"),
                    "Should fail with constraint error");
            }
        });
    });

    describe("Edge Cases", () => {
        it(" should verify bitsTxData matches expected bit representation", async () => {
            const txData = compressTxData(2, 10, 20, 1000, N_LEVELS);
            const witness = await circuit.calculateWitness({ txData }, true);
            await circuit.checkConstraints(witness);

            const expectedBits = txDataToBits(txData);
            const actualBits = getBitsTxDataFromWitness(witness);

            // Check length
            assert.strictEqual(actualBits.length, 8 + N_LEVELS + N_LEVELS + 128, "bitsTxData should have correct length");

            // Check each bit matches expected
            assert.deepStrictEqual(actualBits, expectedBits, "bitsTxData should match expected bit representation");
        });

        it(" should correctly decode various transaction combinations", async () => {
            const testCases = [
                { txnType: 0, fromIdx: 0, toIdx: 0, amount: 0 },
                { txnType: 1, fromIdx: 1, toIdx: 2, amount: 100 },
                { txnType: 2, fromIdx: 10, toIdx: 10, amount: 5000 },
                { txnType: 3, fromIdx: 100, toIdx: 0, amount: 200 },
                { txnType: 4, fromIdx: 50, toIdx: 60, amount: 0 },
                { txnType: 5, fromIdx: 30, toIdx: 40, amount: 0 },
                { txnType: 6, fromIdx: 1000, toIdx: 2000, amount: 9999 },
            ];

            for (const tc of testCases) {
                const txData = compressTxData(tc.txnType, tc.fromIdx, tc.toIdx, tc.amount, N_LEVELS);
                const witness = await circuit.calculateWitness({ txData }, true);
                await circuit.checkConstraints(witness);

                const output = getOutputs(witness);
                assert.strictEqual(output.txnType, tc.txnType, `txnType mismatch for ${JSON.stringify(tc)}`);
                assert.strictEqual(output.fromIdx, tc.fromIdx, `fromIdx mismatch`);
                assert.strictEqual(output.toIdx, tc.toIdx, `toIdx mismatch`);
                assert.strictEqual(Number(output.amount), tc.amount, `amount mismatch`);


                const expectedBits = txDataToBits(txData);
                const actualBits = getBitsTxDataFromWitness(witness);
                assert.deepStrictEqual(actualBits, expectedBits, "bitsTxData should match expected bit representation");
            }
        });
    });
});

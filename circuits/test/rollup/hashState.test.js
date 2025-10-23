/**
 * HashState Circuit Unit Tests
 *
 * Tests the HashState circuit which computes Poseidon hash of account balance.
 * This is the simplest circuit with no dependencies on other custom circuits.
 *
 * Circuit: circuits/syb_rollup/lib/hash-state.circom
 * Template: HashState()
 * Dependencies: circomlib/poseidon
 */

import { describe, it, before, after } from "mocha";
import { strict as assert } from "assert";
import { wasm as tester } from "circom_tester";
import { buildPoseidon } from "circomlibjs";
import path from "path";
import { fileURLToPath } from "url";
import fs from "fs";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

describe("HashState Circuit:(Utility Component)", function() {
    this.timeout(60000);

    let circuit;
    let circuitTmpPath;
    let poseidon;

    before(async () => {
        poseidon = await buildPoseidon();
        const circuitSrc = `
            pragma circom 2.0.0;
            include "../../circuits/syb_rollup/lib/hash-state.circom";
            component main = HashState();
        `;
        circuitTmpPath = path.join(__dirname, "hash-state-temp.circom");
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

    // Helper function to compute expected Poseidon hash
    function computeExpectedHash(balance) {
        const balanceBigInt = BigInt(balance);
        const hash = poseidon([balanceBigInt]);
        return poseidon.F.toString(hash);
    }

    describe("Valid Cases - Basic Functionality", () => {
        it("should hash zero balance", async () => {
            const input = {
                balance: "0"
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);

            const outputHash = witness[circuit.symbols["main.out"].varIdx];
            const expectedHash = computeExpectedHash(input.balance);
            assert.strictEqual(
                outputHash.toString(),
                expectedHash,
                "Circuit output should match expected Poseidon hash"
            );
        });

        it("should hash small balance", async () => {
            const input = {
                balance: "100"
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);

            const outputHash = witness[circuit.symbols["main.out"].varIdx];
            const expectedHash = computeExpectedHash(input.balance);
            assert.strictEqual(
                outputHash.toString(),
                expectedHash,
                "Circuit output should match expected Poseidon hash"
            );
        });

        it("should hash large balance", async () => {
            const input = {
                balance: "1000000000000000000" // 10^18
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);

            const outputHash = witness[circuit.symbols["main.out"].varIdx];
            const expectedHash = computeExpectedHash(input.balance);
            assert.strictEqual(
                outputHash.toString(),
                expectedHash,
                "Circuit output should match expected Poseidon hash"
            );
        });

        it("should hash maximum field-compatible balance", async () => {
            // Use a large but safe value
            const input = {
                balance: "21888242871839275222246405745257275088548364400416034343698204186575808495617"
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);

            const outputHash = witness[circuit.symbols["main.out"].varIdx];
            const expectedHash = computeExpectedHash(input.balance);
            assert.strictEqual(
                outputHash.toString(),
                expectedHash,
                "Circuit output should match expected Poseidon hash"
            );
        });

        it("should produce deterministic output (same input = same hash)", async () => {
            const input = {
                balance: "500"
            };

            // Calculate first witness
            const witness1 = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness1);
            const output1 = witness1[circuit.symbols["main.out"].varIdx];

            // Calculate second witness with same input
            const witness2 = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness2);
            const output2 = witness2[circuit.symbols["main.out"].varIdx];

            // Both outputs should be identical
            assert.strictEqual(
                output1.toString(),
                output2.toString(),
                "Same input should produce identical hash outputs"
            );

            // Verify both match expected Poseidon hash
            const expectedHash = computeExpectedHash(input.balance);
            assert.strictEqual(
                output1.toString(),
                expectedHash,
                "Circuit output should match expected Poseidon hash"
            );
        });

        it("should produce different outputs for different inputs", async () => {
            const input1 = {
                balance: "100"
            };

            const input2 = {
                balance: "101"
            };

            const witness1 = await circuit.calculateWitness(input1, true);
            await circuit.checkConstraints(witness1);
            const output1 = witness1[circuit.symbols["main.out"].varIdx];

            const witness2 = await circuit.calculateWitness(input2, true);
            await circuit.checkConstraints(witness2);
            const output2 = witness2[circuit.symbols["main.out"].varIdx];

            // Outputs should be different
            assert.notStrictEqual(
                output1.toString(),
                output2.toString(),
                "Different inputs should produce different hash outputs"
            );

            // Verify both match their expected Poseidon hashes
            const expectedHash1 = computeExpectedHash(input1.balance);
            const expectedHash2 = computeExpectedHash(input2.balance);
            assert.strictEqual(
                output1.toString(),
                expectedHash1,
                "First circuit output should match expected Poseidon hash"
            );
            assert.strictEqual(
                output2.toString(),
                expectedHash2,
                "Second circuit output should match expected Poseidon hash"
            );
        });
    });

    describe("Edge Cases", () => {
        it("should produce unique hashes for many different balances (collision resistance)", async () => {
            const hashes = new Set();
            const numTests = 100;

            for (let i = 0; i < numTests; i++) {
                const input = {
                    balance: i.toString()
                };

                const witness = await circuit.calculateWitness(input, true);
                await circuit.checkConstraints(witness);
                const outputHash = witness[circuit.symbols["main.out"].varIdx];

                const expectedHash = computeExpectedHash(input.balance);
                assert.strictEqual(
                    outputHash.toString(),
                    expectedHash,
                    `Circuit output for balance ${i} should match expected Poseidon hash`
                );

                // Check for collisions
                const hashStr = outputHash.toString();
                assert(
                    !hashes.has(hashStr),
                    `Hash collision detected for balance ${i}`
                );
                hashes.add(hashStr);
            }

            assert.strictEqual(
                hashes.size,
                numTests,
                "All hashes should be unique"
            );
        });

        it("should produce unique hashes for sequential balance values", async () => {
            const hashes = [];

            for (let i = 0; i <= 100; i++) {
                const input = {
                    balance: i.toString()
                };

                const witness = await circuit.calculateWitness(input, true);
                await circuit.checkConstraints(witness);
                const outputHash = witness[circuit.symbols["main.out"].varIdx];

                const expectedHash = computeExpectedHash(input.balance);
                assert.strictEqual(
                    outputHash.toString(),
                    expectedHash,
                    `Circuit output for balance ${i} should match expected Poseidon hash`
                );

                hashes.push(outputHash.toString());
            }

            // Check that all hashes are unique
            const uniqueHashes = new Set(hashes);
            assert.strictEqual(
                uniqueHashes.size,
                hashes.length,
                "All sequential hashes should be unique"
            );

            // Check that hashes are different from their neighbors
            for (let i = 1; i < hashes.length; i++) {
                assert.notStrictEqual(
                    hashes[i],
                    hashes[i - 1],
                    `Hash for balance ${i} should differ from balance ${i - 1}`
                );
            }
        });
    });
});

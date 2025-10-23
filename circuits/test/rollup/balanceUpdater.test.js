/**
 * BalanceUpdater Circuit Unit Tests
 *
 * Tests the BalanceUpdater circuit which calculates new account balances based on transaction type.
 * Handles balance changes for CreateAccount, Deposit, Withdraw, and Explode transactions.
 *
 * Circuit: circuits/syb_rollup/balance-updater.circom
 * Template: BalanceUpdater()
 * Dependencies: circomlib/mux1
 */

import { describe, it, before, after } from "mocha";
import { strict as assert } from "assert";
import { wasm as tester } from "circom_tester";
import path from "path";
import { fileURLToPath } from "url";
import fs from "fs";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

describe("BalanceUpdater Circuit - (Transaction)", function() {
    this.timeout(120000);

    let circuit;
    let circuitTmpPath;

    before(async () => {
        const circuitSrc = `
            pragma circom 2.0.0;
            include "../../circuits/syb_rollup/balance-updater.circom";
            component main = BalanceUpdater();
        `;
        circuitTmpPath = path.join(__dirname, "balance-updater-temp.circom");
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
            newStBalanceSender: witness[circuit.symbols["main.newStBalanceSender"].varIdx].toString(),
            newStBalanceReceiver: witness[circuit.symbols["main.newStBalanceReceiver"].varIdx].toString()
        };
    }

    describe("Valid Cases - Balance Calculations", () => {
        it("CreateAccount should add balance", async () => {
            const input = {
                oldStBalanceSender: "0",
                oldStBalanceReceiver: "0",
                amount: "1000",
                isCreateAccount: 1,
                isDeposit: 0,
                isWithdraw: 0,
                isExplode: 0
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);
            const output = getOutputs(witness);

            assert.strictEqual(output.newStBalanceSender, "1000");
            assert.strictEqual(output.newStBalanceReceiver, "0");
        });

        it("Deposit should increase balance", async () => {
            const input = {
                oldStBalanceSender: "500",
                oldStBalanceReceiver: "0",
                amount: "300",
                isCreateAccount: 0,
                isDeposit: 1,
                isWithdraw: 0,
                isExplode: 0
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);
            const output = getOutputs(witness);

            assert.strictEqual(output.newStBalanceSender, "800");
            assert.strictEqual(output.newStBalanceReceiver, "0");
        });

        it("Withdraw should decrease balance", async () => {
            const input = {
                oldStBalanceSender: "1000",
                oldStBalanceReceiver: "0",
                amount: "200",
                isCreateAccount: 0,
                isDeposit: 0,
                isWithdraw: 1,
                isExplode: 0
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);
            const output = getOutputs(witness);

            assert.strictEqual(output.newStBalanceSender, "800");
            assert.strictEqual(output.newStBalanceReceiver, "0");
        });

        it("Explode should increase sender and decrease receiver balance", async () => {
            const input = {
                oldStBalanceSender: "500",
                oldStBalanceReceiver: "1000",
                amount: "100",
                isCreateAccount: 0,
                isDeposit: 0,
                isWithdraw: 0,
                isExplode: 1
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);
            const output = getOutputs(witness);

            assert.strictEqual(output.newStBalanceSender, "600");
            assert.strictEqual(output.newStBalanceReceiver, "900");
        });

        it("NOP should not change balances", async () => {
            const input = {
                oldStBalanceSender: "500",
                oldStBalanceReceiver: "300",
                amount: "100",
                isCreateAccount: 0,
                isDeposit: 0,
                isWithdraw: 0,
                isExplode: 0
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);
            const output = getOutputs(witness);

            assert.strictEqual(output.newStBalanceSender, "500");
            assert.strictEqual(output.newStBalanceReceiver, "300");
        });

        // TODO: do we allow this in the contract?
        it("Zero amount deposit should work", async () => {
            const input = {
                oldStBalanceSender: "1000",
                oldStBalanceReceiver: "500",
                amount: "0",
                isCreateAccount: 0,
                isDeposit: 1,
                isWithdraw: 0,
                isExplode: 0
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);
            const output = getOutputs(witness);

            assert.strictEqual(output.newStBalanceSender, "1000");
            assert.strictEqual(output.newStBalanceReceiver, "500");
        });

        it("Large balance values should work", async () => {
            const input = {
                oldStBalanceSender: "1000000000000000000",
                oldStBalanceReceiver: "500000000000000000",
                amount: "100000000000000000",
                isCreateAccount: 0,
                isDeposit: 1,
                isWithdraw: 0,
                isExplode: 0
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);
            const output = getOutputs(witness);

            assert.strictEqual(output.newStBalanceSender, "1100000000000000000");
            assert.strictEqual(output.newStBalanceReceiver, "500000000000000000");
        });

        // TODO: do we allow this in the contract?
        it("Zero balance CreateAccount should work", async () => {
            const input = {
                oldStBalanceSender: "0",
                oldStBalanceReceiver: "0",
                amount: "500",
                isCreateAccount: 1,
                isDeposit: 0,
                isWithdraw: 0,
                isExplode: 0
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);
            const output = getOutputs(witness);

            assert.strictEqual(output.newStBalanceSender, "500");
            assert.strictEqual(output.newStBalanceReceiver, "0");
        });

        it("Receiver unchanged for non-explode transactions", async () => {
            const testCases = [
                { isCreateAccount: 1, isDeposit: 0, isWithdraw: 0, isExplode: 0 },
                { isCreateAccount: 0, isDeposit: 1, isWithdraw: 0, isExplode: 0 },
                { isCreateAccount: 0, isDeposit: 0, isWithdraw: 1, isExplode: 0 },
            ];

            const receiverBalance = "1000";

            for (const flags of testCases) {
                const input = {
                    oldStBalanceSender: "500",
                    oldStBalanceReceiver: receiverBalance,
                    amount: "100",
                    ...flags
                };

                const witness = await circuit.calculateWitness(input, true);
                await circuit.checkConstraints(witness);
                const output = getOutputs(witness);

                assert.strictEqual(
                    output.newStBalanceReceiver,
                    receiverBalance,
                    `Receiver should be unchanged for flags ${JSON.stringify(flags)}`
                );
            }
        });
    });

    describe("Edge Cases", () => {
        it(" Explode affects both sender and receiver correctly", async () => {
            const input = {
                oldStBalanceSender: "1000",
                oldStBalanceReceiver: "2000",
                amount: "300",
                isCreateAccount: 0,
                isDeposit: 0,
                isWithdraw: 0,
                isExplode: 1
            };

            const witness = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(witness);
            const output = getOutputs(witness);

            // Explode: sender gains amount, receiver loses amount
            assert.strictEqual(output.newStBalanceSender, "1300");
            assert.strictEqual(output.newStBalanceReceiver, "1700");

            // Verify net change is zero (amount transferred)
            const oldTotal = 1000n + 2000n;
            const newTotal = BigInt(output.newStBalanceSender) + BigInt(output.newStBalanceReceiver);
            assert.strictEqual(newTotal, oldTotal, "Total balance should be preserved in explode");
        });
    });
});

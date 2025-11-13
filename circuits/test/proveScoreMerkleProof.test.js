import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";
import { describe, it, before, after } from "mocha";
import { strict as assert } from "assert";
import { wasm as tester } from "circom_tester";
import { SmtTree, generateRandomSmt } from "./utils/smt.js";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const N_LEVELS = 4;
const NUM_TESTCASES = 10;

describe("ProveScoreMerkleProof circuit test", function () {
    this.timeout(100000);

    let circuit;
    let circuitTmpPath;
    let smts;

    before(async () => {
        const circuitSrc = `
            pragma circom 2.0.0;
            include "../circuits/prove_score_inclusion.circom";
            component main = ProveScoreInclusion(4);
        `;
        circuitTmpPath = path.join(__dirname, "prove-score-merkle-proof-temp.circom");
        fs.writeFileSync(circuitTmpPath, circuitSrc, "utf8");

        circuit = await tester(circuitTmpPath, {
            reduceConstraints: false,
            include: path.join(__dirname, "../circuits")
        });
        await circuit.loadConstraints();
        console.log("Constraints:", circuit.constraints.length);

        // generate x random SMT trees
        // each SMT has different data.
        smts = await generateRandomSmt(N_LEVELS, NUM_TESTCASES);
    });

    after(() => {
        if (fs.existsSync(circuitTmpPath)) {
            fs.unlinkSync(circuitTmpPath);
        }
    });


    it("should verify a simple single entry SMT", async () => {
        const smt = new SmtTree(N_LEVELS);
        await smt.init();
        
        const testKey = 123n;
        const testScore = 456n;
        await smt.insert(testKey, testScore);
        const siblings = await smt.getSiblings(testKey);

        const root = await smt.getRoot();
        
        const w = await circuit.calculateWitness({
            idx: testKey.toString(),
            score: testScore.toString(),
            root: smt.Fr.toObject(root),
            siblings: siblings
        }, true);
        
        await circuit.checkConstraints(w);
    });

    // testing with undetermined data
    it("should verify 10 random SMT trees with valid proofs", async () => {
        for (let i = 0; i < smts.length; i++) {
            const smt = smts[i];
            const keys = smt.keys;
            const scores = smt.scores;
      
            try {
                // Pick a random key to test from this SMT
                const testIndex = Math.floor(Math.random() * keys.length);
                const testKey = keys[testIndex];
                const testScore = scores[testIndex];

                // Get proof for the test key
                const siblings = await smt.getSiblings(testKey);
                const root = await smt.getRoot();

                // Verify the proof
                const w = await circuit.calculateWitness({
                    idx: testKey.toString(),
                    score: testScore.toString(),
                    root: smt.Fr.toObject(root),
                    siblings: siblings
                }, true);

                await circuit.checkConstraints(w);
            } catch (error) {
                throw new Error(
                    `SMT ${i} failed verification for undetermined data: ${error.message}`
                );
            }
        }
    });

    it("should fail with invalid proof (wrong siblings)", async () => {
        const smt = new SmtTree(N_LEVELS);
        await smt.init();
        
        const testKey = 123n;
        const testScore = 456n;
        await smt.insert(testKey, testScore);
        
        const root = await smt.getRoot();
        const Fr = smt.Fr;
        
        // Use wrong siblings
        const wrongSiblings = [999, 888, 777, 666];
        
        try {
            await circuit.calculateWitness({
                idx: testKey.toString(),
                score: testScore.toString(),
                root: Fr.toObject(root),
                siblings: wrongSiblings
            }, true);
            throw new Error("Should have failed with wrong siblings");
        } catch (error) {
            assert(error.message.includes("Assert Failed"), `Expected circuit to reject with Assert Failed, but got: ${error.message}`);
        }
    });

    it("should fail with wrong score", async () => {
        const smt = new SmtTree(N_LEVELS);
        await smt.init();
        
        const testKey = 123n;
        const testScore = 456n;
        await smt.insert(testKey, testScore);
        
        const siblings = await smt.getSiblings(testKey);        
        const root = await smt.getRoot();
        
        try {
            await circuit.calculateWitness({
                idx: testKey.toString(),
                score: "999", // Wrong score
                root: smt.Fr.toObject(root),
                siblings: siblings
            }, true);
            throw new Error("Should have failed with wrong score");
        } catch (error) {
            assert(error.message.includes("Assert Failed"), `Expected circuit to reject with Assert Failed, but got: ${error.message}`);
        }
    });

    it("should handle multiple entries in same SMT", async () => {
        const smt = new SmtTree(N_LEVELS);
        await smt.init();
        
        // Insert multiple entries
        const entries = [
            { key: 1n, score: 100n },
            { key: 2n, score: 200n },
            { key: 3n, score: 300n }
        ];
        
        for (const entry of entries) {
            await smt.insert(entry.key, entry.score);
        }
        
        // Test each entry
        for (const entry of entries) {
            const siblings = await smt.getSiblings(entry.key);            
            const root = await smt.getRoot();
            
            const w = await circuit.calculateWitness({
                idx: entry.key.toString(),
                score: entry.score.toString(),
                root: smt.Fr.toObject(root),
                siblings: siblings
            }, true);
            
            await circuit.checkConstraints(w);
        }
        
        console.log("✓ Multiple entries in same SMT verified successfully");
    });
});
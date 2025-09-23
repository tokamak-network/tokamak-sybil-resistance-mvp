const fs = require("fs");
const path = require("path");
const { describe, it, before, after } = require("mocha");
const { strict: assert } = require("assert");
const { wasm: tester } = require("circom_tester");

describe("ProveScoreMerkleProof circuit test", function () {
    this.timeout(100000);

    let circuit;
    let circuitTmpPath;

    before(async () => {
        const circuitSrc = `
            pragma circom 2.0.0;
            include "../circuits/prove_score_inclusion.circom";
            component main = ProveScoreInclusion(2);
        `;
        circuitTmpPath = path.join(__dirname, "prove-score-merkle-proof-temp.circom");
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

    it("should fail with invalid idx", async () => {
        const input = {
            idx: "20",
            score: "88", 
            root: "8199520123371559548495425428157097842501569495702004037304582533739096128775", // Correct root but wrong score
            siblings: ["0", "0"]
        };

        try {
            const w = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(w);
            assert.fail("Expected circuit to fail but it passed");
        } catch (error) {
            console.log("Invalid SMT data correctly failed - circuit constraints violated");
            assert(error.message.includes("Assert Failed"), "Expected assert failure");
        }
    });

    it ("should fail with invalid score", async () => {
        const input = {
            idx: "8",
            score: "16",
            root: "8199520123371559548495425428157097842501569495702004037304582533739096128775",
            siblings: ["0", "0"]
        };

        try {
            const w = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(w);
            assert.fail("Expected circuit to fail but it passed");
        } catch (error) {
            console.log("Invalid SMT data correctly failed - circuit constraints violated");
            assert(error.message.includes("Assert Failed"), "Expected assert failure");
        }
    })

    it("should fail with completely wrong root", async () => {
        const input = {
            idx: "8",
            score: "88",
            root: "123456789", // Completely wrong root
            siblings: ["0", "0"]
        };

        try {
            const w = await circuit.calculateWitness(input, true);
            await circuit.checkConstraints(w);
            assert.fail("Expected circuit to fail but it passed");
        } catch (error) {
            console.log("Wrong root correctly failed - circuit constraints violated");
            assert(error.message.includes("Assert Failed"), "Expected assert failure");
        }
    });

    it("should prove score inclusion with valid SMT data", async () => {
        // Using valid SMT test data from test_data.json 
        // This tests inclusion of key=8, value=88 in an SMT tree
        const input = {
            idx: "8",
            score: "88", 
            root: "8199520123371559548495425428157097842501569495702004037304582533739096128775",
            siblings: ["0", "0"]
        };

        const w = await circuit.calculateWitness(input, true);
        await circuit.checkConstraints(w);
        
        console.log("Score inclusion proof successful - circuit constraints satisfied");
    });

    it("should prove inclusion of another valid entry", async () => {
        // Another valid SMT entry: key=9, value=99
        const input = {
            idx: "9",
            score: "99",
            root: "9586840611691950490797560970510543395125579325564362838402424088587314740100", 
            siblings: ["0", "0"]
        };

        const w = await circuit.calculateWitness(input, true);
        await circuit.checkConstraints(w);
        
        console.log("Second inclusion proof successful - circuit constraints satisfied");
    });
});

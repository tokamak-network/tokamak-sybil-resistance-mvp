const fs = require("fs");
const path = require("path");
const { wasm: tester } = require("circom_tester");

/**
 * Generate a circuit for testing with flexible parameters
 * @param {string} circuitName - Name of the circuit template
 * @param {string} circuitPath - Path to the circuit file (relative to circuits directory)
 * @param {Object} options - Configuration options
 * @param {string[]} options.publicOutputs - Array of public output signal names
 * @param {string} options.templateParams - Parameters to pass to the template (e.g., "(3, 2)")
 * @param {string} options.includePath - Additional include path for dependencies
 * @returns {Object} - Object containing circuit instance and cleanup function
 */
async function generateCircuit(circuitName, circuitPath, options = {}) {
    const {
        publicOutputs = ["root"],
        templateParams = "",
        includePath = "../circuits"
    } = options;

    // Create the circuit source code
    const circuitSrc = `
        pragma circom 2.0.0;
        include "${circuitPath}";
        component main{public [${publicOutputs.join(", ")}]} = ${circuitName}${templateParams};
    `;

    // Create temporary file path
    const circuitTmpPath = path.join(__dirname, `${circuitName}-temp.circom`);
    fs.writeFileSync(circuitTmpPath, circuitSrc, "utf8");

    // Load the circuit
    const circuit = await tester(circuitTmpPath, {
        reduceConstraints: false,
        include: path.join(__dirname, includePath)
    });
    await circuit.loadConstraints();

    // Return circuit and cleanup function
    return {
        circuit,
        circuitTmpPath,
        cleanup: () => {
            if (fs.existsSync(circuitTmpPath)) {
                fs.unlinkSync(circuitTmpPath);
            }
        }
    };
}

/**
 * Helper function to get empty siblings array for SMT tree
 * @param {number} n - Number of levels
 * @param {Object} F - Field object from ffjavascript
 * @returns {string[]} - Array of zero values as strings
 */
function getEmptySiblings(n, F) {
    return Array(n).fill(F.toString(F.zero));
}

module.exports = {
    generateCircuit,
    getEmptySiblings
};
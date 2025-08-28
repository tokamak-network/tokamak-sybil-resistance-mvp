const fs = require("fs");
const path = require("path");
const { wasm: tester } = require("circom_tester");

/**
 * Generate a circuit for testing with flexible parameters
 * @param {string} circuitName - Name of the circuit template
 * @param {string} circuitPath - Path to the circuit file (relative to circuits directory)
 * @param {Object} options - Configuration options
 * @param {string[]} options.publicInputs - Array of public input signal names (default: empty)
 * @param {string} options.templateParams - Parameters to pass to the template (e.g., "(3, 2)")
 * @param {string} options.includePath - Additional include path for dependencies
 * @returns {Object} - Object containing circuit instance and cleanup function
 */
async function generateCircuit(circuitName, circuitPath, options = {}) {
    const {
        publicInputs = [],
        templateParams = "",
        includePath = "../circuits"
    } = options;

    // Create the circuit source code
    const publicList = publicInputs.length > 0 ? `{public [${publicInputs.join(", ")}]}` : "";
    const circuitSrc = `
        pragma circom 2.0.0;
        include "${circuitPath}";
        component main${publicList} = ${circuitName}${templateParams};
    `;

    // Create temporary file path
    const circuitTmpPath = path.join(__dirname, `${circuitName}-temp.circom`);
    fs.writeFileSync(circuitTmpPath, circuitSrc, "utf8");

    try {
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
            // cleanup function to delete the tmp files
            cleanup: () => {
                if (fs.existsSync(circuitTmpPath)) {
                    fs.unlinkSync(circuitTmpPath);
                }
            }
        };
    } catch (error) {
        // Clean up temp file even if circuit loading fails
        if (fs.existsSync(circuitTmpPath)) {
            fs.unlinkSync(circuitTmpPath);
        }
        throw error;
    }
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
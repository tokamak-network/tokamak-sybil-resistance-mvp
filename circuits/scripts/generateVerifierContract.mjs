import fs from 'fs/promises'
import path from 'path'
import { toPascalCase, execCommand } from './utils/helpers.mjs'
import { log } from './utils/logger.mjs'

// Generate Solidity verifier contract
export async function generateVerifierContract(circuitName, zkeyPath, vkeyPath, projectDir) {
    log.step('Generating Solidity verifier contract')

    const buildDir = path.join(projectDir, 'build', circuitName)
    const verifierPath = path.join(buildDir, 'groth16_verifier.sol')

    try {
        // Read verification key
        const vKey = JSON.parse(await fs.readFile(vkeyPath, 'utf-8'))

        // Generate verifier using the templates
        const templates = await import('snarkjs').then(m => m.groth16)
        const verifierCode = await templates.exportSolidityVerifier(vKey, null)

        // Customize contract name
        const contractName = `${toPascalCase(circuitName)}Verifier`
        const customizedCode = verifierCode.replace(
            /contract Groth16Verifier/g,
            `contract ${contractName}`
        )

        await fs.writeFile(verifierPath, customizedCode)

        // Copy to verifiers directory
        const verifiersDir = path.join(projectDir, 'verifiers')
        await fs.mkdir(verifiersDir, { recursive: true })

        const finalPath = path.join(verifiersDir, `${contractName}.sol`)
        await fs.copyFile(verifierPath, finalPath)

        log.success(`Verifier contract saved: verifiers/${contractName}.sol`)

        return { verifierPath: finalPath, contractName }
    } catch (error) {
        // Fallback: use snarkjs CLI
        log.warning('Using snarkjs CLI fallback for verifier generation')
        const tempVerifierPath = path.join(buildDir, 'temp_verifier.sol')

        execCommand(`npx snarkjs zkey export solidityverifier ${zkeyPath} ${tempVerifierPath}`, projectDir)

        const verifierCode = await fs.readFile(tempVerifierPath, 'utf-8')
        const contractName = `${toPascalCase(circuitName)}Verifier`
        const customizedCode = verifierCode.replace(
            /contract Groth16Verifier/g,
            `contract ${contractName}`
        )

        await fs.writeFile(verifierPath, customizedCode)

        // Copy to verifiers directory
        const verifiersDir = path.join(projectDir, 'verifiers')
        await fs.mkdir(verifiersDir, { recursive: true })

        const finalPath = path.join(verifiersDir, `${contractName}.sol`)
        await fs.copyFile(verifierPath, finalPath)

        // Clean up temp file
        await fs.unlink(tempVerifierPath).catch(() => {})

        log.success(`Verifier contract saved: verifiers/${contractName}.sol`)

        return { verifierPath: finalPath, contractName }
    }
}

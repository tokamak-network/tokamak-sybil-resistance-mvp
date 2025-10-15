import fs from 'fs/promises'
import path from 'path'
import { execCommand } from './utils/helpers.mjs'
import { log } from './utils/logger.mjs'

// Compile circom circuit
export async function compileCircuit(circuitName, circuitFile, projectDir) {
    log.step(`Compiling circuit: ${circuitName}`)

    const buildDir = path.join(projectDir, 'build', circuitName)
    await fs.mkdir(buildDir, { recursive: true })

    const circuitPath = path.join(projectDir, 'circuits', `${circuitFile}.circom`)
    const outputR1cs = path.join(buildDir, `${circuitName}.r1cs`)
    const outputWasm = path.join(buildDir, `${circuitName}_js`, `${circuitName}.wasm`)

    // Check if circuit file exists
    try {
        await fs.access(circuitPath)
    } catch {
        throw new Error(`Circuit file not found: ${circuitPath}`)
    }

    // Check if already compiled
    try {
        await fs.access(outputR1cs)
        await fs.access(outputWasm)
        log.success(`Circuit already compiled: ${circuitName}`)
        return { r1csPath: outputR1cs, wasmPath: outputWasm }
    } catch {
        // Compile
        log.info(`Compiling ${circuitPath}`)
        execCommand(`circom ${circuitPath} --r1cs --wasm --sym -o ${buildDir}`, projectDir)
        log.success(`Compiled: ${circuitName}`)
        return { r1csPath: outputR1cs, wasmPath: outputWasm }
    }
}

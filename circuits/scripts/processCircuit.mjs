import { compileCircuit } from './compileCircuit.mjs'
import { getCircuitInfo } from './utils/getCircuitInfo.mjs'
import { generateKeys } from './generateKeys.mjs'
import { generateVerifierContract } from './generateVerifierContract.mjs'
import { downloadPtau } from './utils/downloadPtau.mjs'
import { getPtauFile } from './utils/helpers.mjs'
import { log, colors } from './utils/logger.mjs'

// Process a single circuit
export async function processCircuit(circuitName, circuitConfig, projectDir) {
    console.log('\n' + '='.repeat(60))
    console.log(`${colors.cyan}Processing circuit: ${circuitName}${colors.reset}`)
    console.log('='.repeat(60) + '\n')

    const startTime = Date.now()

    try {
        // 1. Compile circuit
        const { r1csPath } = await compileCircuit(circuitName, circuitConfig.file, projectDir)

        // 2. Get circuit info
        const info = await getCircuitInfo(r1csPath)

        // 3. Download PTAU if needed
        const ptauFile = getPtauFile(info.nConstraints)
        const ptauPath = await downloadPtau(ptauFile, projectDir)

        // 4. Generate keys
        const { zkeyPath, vkeyPath } = await generateKeys(circuitName, r1csPath, ptauPath, projectDir)

        // 5. Generate verifier contract
        const { contractName } = await generateVerifierContract(circuitName, zkeyPath, vkeyPath, projectDir)

        const duration = ((Date.now() - startTime) / 1000).toFixed(2)
        log.success(`Successfully processed ${circuitName} in ${duration}s`)

        return { success: true, circuitName, contractName, duration }
    } catch (error) {
        log.error(`Failed to process ${circuitName}`)
        console.error(error.message)
        return { success: false, circuitName, error: error.message }
    }
}

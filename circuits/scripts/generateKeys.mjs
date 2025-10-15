import fs from 'fs/promises'
import path from 'path'
import * as snarkjs from 'snarkjs'
import { log } from './utils/logger.mjs'

// Generate proving and verification keys
export async function generateKeys(circuitName, r1csPath, ptauPath, projectDir) {
    const buildDir = path.join(projectDir, 'build', circuitName)
    const zkeyPath = path.join(buildDir, 'groth16_pkey.zkey')
    const vkeyPath = path.join(buildDir, 'groth16_vkey.json')

    // Check if keys already exist
    try {
        await fs.access(zkeyPath)
        await fs.access(vkeyPath)
        log.success(`Keys already exist for: ${circuitName}`)
        return { zkeyPath, vkeyPath }
    } catch {
        log.step('Generating proving key (this may take a while)...')

        // Generate zkey
        const { zkey: finalZkey } = await snarkjs.zKey.newZKey(
            r1csPath,
            ptauPath,
            zkeyPath
        )

        log.success('Proving key generated')

        // Export verification key
        log.step('Exporting verification key')
        const vKey = await snarkjs.zKey.exportVerificationKey(zkeyPath)
        await fs.writeFile(vkeyPath, JSON.stringify(vKey, null, 2))
        log.success('Verification key exported')

        return { zkeyPath, vkeyPath }
    }
}

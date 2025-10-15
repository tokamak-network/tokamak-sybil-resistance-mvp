import * as snarkjs from 'snarkjs'
import { log } from './logger.mjs'

// Get circuit info
export async function getCircuitInfo(r1csPath) {
    log.step('Getting circuit info')
    const info = await snarkjs.r1cs.info(r1csPath)

    console.log(`  Prime Field: ${info.curve}`)
    console.log(`  # of Wires: ${info.nVars}`)
    console.log(`  # of Constraints: ${info.nConstraints}`)
    console.log(`  # of Private Inputs: ${info.nPrvInputs}`)
    console.log(`  # of Public Inputs: ${info.nPubInputs}`)
    console.log(`  # of Labels: ${info.nLabels}`)

    return info
}

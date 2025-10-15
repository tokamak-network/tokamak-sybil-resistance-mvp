#!/usr/bin/env node
/**
 * this script is used to generate the verifier contracts for the circuits
 * 1. compile the circuits => using `compileCircuit.mjs`
 * 2. generate the keys => using `generateKeys.mjs`
 * 3. generate the verifier contract => using `generateVerifierContract.mjs`
 */
import fs from 'fs/promises'
import path from 'path'
import { fileURLToPath } from 'url'
import { processCircuit } from './processCircuit.mjs'
import { log, colors } from './utils/logger.mjs'

const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)
const projectDir = path.join(__dirname, '..')

// Main function
async function main() {
    console.log(`${colors.blue}${'='.repeat(60)}`)
    console.log('  Circom Verifier Generator')
    console.log('=' + '='.repeat(59) + colors.reset + '\n')

    const startTime = Date.now()

    try {
        // Read circuits.json
        // now we have 2 circuits to compile:
        // 1. batch_main: the main circuit for batchForge
        // 2. prove_score_inclusion: the circuit for proving the inclusion of a score in the tree
        const circuitsJsonPath = path.join(projectDir, 'circuits.json')
        const circuitsJson = JSON.parse(await fs.readFile(circuitsJsonPath, 'utf-8'))

        log.info(`Found ${Object.keys(circuitsJson).length} circuit(s) in circuits.json`)

        // Create necessary directories
        await fs.mkdir(path.join(projectDir, 'build'), { recursive: true })
        await fs.mkdir(path.join(projectDir, 'ptau'), { recursive: true })
        await fs.mkdir(path.join(projectDir, 'verifiers'), { recursive: true })

        // Process each circuit
        const results = []
        for (const [circuitName, circuitConfig] of Object.entries(circuitsJson)) {
            const result = await processCircuit(circuitName, circuitConfig, projectDir)
            results.push(result)
        }

        // Summary
        const totalDuration = ((Date.now() - startTime) / 1000).toFixed(2)
        console.log('\n' + '='.repeat(60))
        console.log(`${colors.cyan}Summary${colors.reset}`)
        console.log('='.repeat(60) + '\n')

        const successful = results.filter(r => r.success)
        const failed = results.filter(r => !r.success)

        log.info(`Total time: ${totalDuration}s`)
        log.success(`Successfully processed: ${successful.length} circuit(s)`)

        if (successful.length > 0) {
            console.log('\nGenerated verifier contracts:')
            for (const result of successful) {
                console.log(`  📄 verifiers/${result.contractName}.sol (${result.duration}s)`)
            }
        }

        if (failed.length > 0) {
            log.error(`Failed circuits: ${failed.length}`)
            for (const result of failed) {
                console.log(`  ❌ ${result.circuitName}: ${result.error}`)
            }
        }

        console.log('\n' + colors.green + '🎉 Verifier generation completed!' + colors.reset)

        process.exit(failed.length > 0 ? 1 : 0)
    } catch (error) {
        log.error('Fatal error')
        console.error(error)
        process.exit(1)
    }
}

main()

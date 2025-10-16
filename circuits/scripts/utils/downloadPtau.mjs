import fs from 'fs/promises'
import path from 'path'
import { execCommand } from './helpers.mjs'
import { log } from './logger.mjs'

// Download PTAU file if not exists
export async function downloadPtau(ptauFile, projectDir) {
    const ptauDir = path.join(projectDir, 'ptau')
    const ptauPath = path.join(ptauDir, ptauFile)

    // Check if ptau file exists
    try {
        await fs.access(ptauPath)
        log.success(`PTAU file already exists: ${ptauFile}`)
        return ptauPath
    } catch {
        log.info(`Downloading PTAU file: ${ptauFile}`)
        await fs.mkdir(ptauDir, { recursive: true })

        const power = ptauFile.match(/final_(\d+)\.ptau/)[1]
        const url = `https://hermez.s3-eu-west-1.amazonaws.com/${ptauFile}`

        log.step(`Downloading from ${url}`)
        execCommand(`curl -o ${ptauPath} ${url}`, projectDir)
        log.success(`Downloaded: ${ptauFile}`)
        return ptauPath
    }
}

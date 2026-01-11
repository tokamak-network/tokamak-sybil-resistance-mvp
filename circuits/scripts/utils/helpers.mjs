import { execSync } from 'child_process'

// Convert snake_case to PascalCase
export function toPascalCase(str) {
    return str
        .split('_')
        .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
        .join('')
}

// Execute shell command
export function execCommand(cmd, cwd) {
    try {
        return execSync(cmd, {
            cwd,
            encoding: 'utf-8',
            stdio: 'pipe'
        })
    } catch (error) {
        throw new Error(`Command failed: ${cmd}\n${error.message}`)
    }
}

// Get PTAU file for a given constraint count
export function getPtauFile(constraintCount) {
    // Calculate the power of 2 needed
    const power = Math.ceil(Math.log2(constraintCount))
    return `powersOfTau28_hez_final_${power}.ptau`
}

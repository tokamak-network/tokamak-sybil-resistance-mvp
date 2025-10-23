/**
 * Test Utilities for SYB Rollup Circuit Tests
 * Provides helper functions for encoding transaction data
 */

/**
 * Compress transaction data into packed format
 * Format: txnType (8 bits) | fromIdx (nLevels bits) | toIdx (nLevels bits) | amount (128 bits)
 * @param {number} txnType - Transaction type (0-6)
 * @param {number|string} fromIdx - Sender index
 * @param {number|string} toIdx - Receiver index
 * @param {number|string} amount - Transaction amount
 * @param {number} nLevels - Number of levels in merkle tree
 * @returns {string} Packed transaction data as string
 */
export function compressTxData(txnType, fromIdx, toIdx, amount, nLevels) {
    const txnTypeBig = BigInt(txnType);
    const fromIdxBig = BigInt(fromIdx);
    const toIdxBig = BigInt(toIdx);
    const amountBig = BigInt(amount);

    // Pack: amount << (8 + nLevels + nLevels) | toIdx << (8 + nLevels) | fromIdx << 8 | txnType
    const packed = (amountBig << BigInt(8 + nLevels + nLevels)) |
                   (toIdxBig << BigInt(8 + nLevels)) |
                   (fromIdxBig << BigInt(8)) |
                   txnTypeBig;
    return packed.toString();
}

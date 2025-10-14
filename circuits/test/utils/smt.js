import { newMemEmptyTrie } from "circomlibjs";


/**
 * Sparse Merkle Tree class
 * supports insert, update, delete, and getRoot
 * uses circomlibjs newMemEmptyTrie
 */
class SmtTree {
    constructor(nLevels) {
        this.nLevels = nLevels;
        this.tree = null;
        this.initialized = false;
    }

    async init() {
        if (!this.initialized) {
            this.tree = await newMemEmptyTrie();
            this.Fr = this.tree.F;
            this.initialized = true;
        }
    }

    async insert(key, value) {
        await this.init();
        await this.tree.insert(key, value);
        return this.tree.root;
    }

    async update(key, value) {
        await this.init();
        await this.tree.update(key, value);
        return this.tree.root;
    }

    async delete(key) {
        await this.init();
        await this.tree.delete(key);
        return this.tree.root;
    }

    async getRoot() {
        await this.init();
        return this.tree.root;
    }

    async getSiblings(key) {
        await this.init();
        const Fr = this.Fr;
        const res = await this.tree.find(Fr.e(key));
        
        if (!res.found) {
            throw new Error(`Key ${key} not found in SMT`);
        }

        let siblings = res.siblings;

        // Convert to numbers first
        for (let j = 0; j < siblings.length; j++) {
            siblings[j] = Fr.toObject(siblings[j]);
        }

        // SMTLevIns constraint: siblings[nLevels-1] MUST be 0
        // If we have too many siblings, we need to ensure the last one is 0
        if (siblings.length >= this.nLevels) {
            // Take only the first nLevels siblings
            siblings = siblings.slice(0, this.nLevels);
        } else {
            // Pad with zeros to reach nLevels
            while (siblings.length < this.nLevels) {
                siblings.push(0);
            }
        }

        return siblings;
    }

    getEmptySiblings(n) {
        if (!this.initialized) {
            throw new Error("SMT must be initialized before calling getEmptySiblings");
        }
        return Array(n).fill(this.Fr.toString(this.Fr.zero));
    }
}

/**
 * generate random SMT test data
 * parameters:
 * - nLevels (tree height)
 * - numCases (number of SMT trees to generate)
 * returns:
 * - array of SmtTree objects
 */
async function generateRandomSmt(nLevels = 2, numCases = 3) {
    const smts = [];
    while (smts.length < numCases) {
        // Random number of entries per tree: 1 to nLevels entries
        const numEntries = 1 + Math.floor(Math.random() * nLevels);

        // Generate candidate keys/scores (generate more candidates than needed)
        const candidates = [];
        for (let j = 0; j < numEntries * 5; j++) {
            candidates.push({
                key: BigInt(Math.floor(Math.random() * 1000000)),
                score: BigInt(Math.floor(Math.random() * 1000000000000))
            });
        }

        // Try each candidate and keep only valid ones
        const validEntries = [];
        for (const candidate of candidates) {
            const tempSmt = new SmtTree(nLevels);
            await tempSmt.init();

            // Insert all previously valid entries
            for (const entry of validEntries) {
                await tempSmt.insert(entry.key, entry.score);
            }

            // Insert this candidate
            await tempSmt.insert(candidate.key, candidate.score);

            // Check if this candidate creates a valid proof
            try {
                const siblings = await tempSmt.getSiblings(candidate.key);
                if (siblings[nLevels - 1] === 0) {
                    validEntries.push(candidate);
                }
            } catch (error) {
                // Skip invalid entries
                continue;
            }

            // Stop once we have enough valid entries
            if (validEntries.length >= numEntries) {
                break;
            }
        }

        // Build the final SMT with only valid entries
        if (validEntries.length >= numEntries) {
            const smt = new SmtTree(nLevels);
            await smt.init();

            const keys = [];
            const scores = [];

            for (const entry of validEntries) {
                await smt.insert(entry.key, entry.score);
                keys.push(entry.key);
                scores.push(entry.score);
            }

            smt.keys = keys;
            smt.scores = scores;
            smts.push(smt);
        }
    }
    return smts;
}

export { SmtTree, generateRandomSmt };
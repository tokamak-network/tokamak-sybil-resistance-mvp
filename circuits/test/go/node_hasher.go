package nodehasher

import (
	"fmt"
	"math/big"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

// NodeHasher computes NodeHash_G(v) for a vertex in the GraphTree
// This replicates the circuit logic from node_hasher.circom
//
// NodeHash Algorithm:
// - First block B_0: [v, d, u_0, ..., u_13]  (14 neighbors)
// - acc = Poseidon_16(B_0)
// - For remaining neighbors, process in chunks of 15:
//   - B_b = [acc, u_14+(b-1)*15, ..., u_14+(b-1)*15+14]
//   - acc = Poseidon_16(B_b)
//
// Parameters:
//   - v: Vertex ID
//   - d: Degree of vertex (number of actual neighbors)
//   - neighbors: Neighbor array (should be sorted and padded with zeros)
//   - maxDeg: Maximum degree supported
//
// Returns:
//   - hash: NodeHash_G(v) as a big.Int
//   - error: if validation fails

// CalculatePadLen calculates the padded length based on maxDeg
// padLen = 14 + 15 * numR, where numR = ceil((maxDeg - 14) / 15)
func CalculatePadLen(maxDeg int) int {
	if maxDeg <= 14 {
		return 14
	}
	numR := (maxDeg-14+14) / 15 // Integer division for ceiling
	return 14 + 15*numR
}

// CalculateNumRounds calculates the number of continuation rounds
func CalculateNumRounds(maxDeg int) int {
	if maxDeg <= 14 {
		return 0
	}
	return (maxDeg-14+14) / 15
}

// ValidateNeighbors validates that the neighbor array is properly formatted:
// 1. Neighbors are sorted in strictly ascending order up to degree d
// 2. All values at index >= d are zero (proper padding)
func ValidateNeighbors(neighbors []*big.Int, d int) error {
	if d < 0 || d > len(neighbors) {
		return fmt.Errorf("invalid degree %d for neighbor array of length %d", d, len(neighbors))
	}

	// Check strictly ascending order for valid neighbors (indices 0 to d-1)
	for i := 0; i < d-1; i++ {
		if neighbors[i].Cmp(neighbors[i+1]) >= 0 {
			return fmt.Errorf("neighbors not strictly ascending at index %d: %s >= %s",
				i, neighbors[i].String(), neighbors[i+1].String())
		}
	}

	// Check padding (indices d to len-1 must be zero)
	zero := big.NewInt(0)
	for i := d; i < len(neighbors); i++ {
		if neighbors[i].Cmp(zero) != 0 {
			return fmt.Errorf("padding violation at index %d: expected 0, got %s",
				i, neighbors[i].String())
		}
	}

	return nil
}

// ComputeNodeHash computes the node hash following the NodeHasher circuit logic
func ComputeNodeHash(v, d int64, neighbors []*big.Int, maxDeg int) (*big.Int, error) {
	padLen := CalculatePadLen(maxDeg)

	// Validate input
	if len(neighbors) != padLen {
		return nil, fmt.Errorf("neighbor array length %d does not match expected padLen %d for maxDeg %d",
			len(neighbors), padLen, maxDeg)
	}

	if d < 0 || d > int64(maxDeg) {
		return nil, fmt.Errorf("degree %d out of range [0, %d]", d, maxDeg)
	}

	// Validate neighbors (ascending order and proper padding)
	if err := ValidateNeighbors(neighbors, int(d)); err != nil {
		return nil, fmt.Errorf("neighbor validation failed: %w", err)
	}

	// First block: [v, d, nbr[0..13]] (16 elements total)
	firstBlock := make([]*big.Int, 16)
	firstBlock[0] = big.NewInt(v)
	firstBlock[1] = big.NewInt(d)
	for i := 0; i < 14; i++ {
		firstBlock[2+i] = new(big.Int).Set(neighbors[i])
	}

	// Hash first block
	acc, err := poseidon.Hash(firstBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to hash first block: %w", err)
	}

	// Continuation blocks (15 neighbors each)
	numR := CalculateNumRounds(maxDeg)
	for round := 0; round < numR; round++ {
		block := make([]*big.Int, 16)
		block[0] = new(big.Int).Set(acc)

		startIdx := 14 + round*15
		for i := 0; i < 15; i++ {
			idx := startIdx + i
			block[1+i] = new(big.Int).Set(neighbors[idx])
		}

		acc, err = poseidon.Hash(block)
		if err != nil {
			return nil, fmt.Errorf("failed to hash continuation block %d: %w", round, err)
		}
	}

	return acc, nil
}

// PadNeighbors creates a properly padded neighbor array
func PadNeighbors(neighbors []int64, padLen int) []*big.Int {
	result := make([]*big.Int, padLen)
	for i := 0; i < padLen; i++ {
		if i < len(neighbors) {
			result[i] = big.NewInt(neighbors[i])
		} else {
			result[i] = big.NewInt(0)
		}
	}
	return result
}

// PadNeighborsBigInt creates a properly padded neighbor array from big.Int inputs
func PadNeighborsBigInt(neighbors []*big.Int, padLen int) []*big.Int {
	result := make([]*big.Int, padLen)
	for i := 0; i < padLen; i++ {
		if i < len(neighbors) {
			result[i] = new(big.Int).Set(neighbors[i])
		} else {
			result[i] = big.NewInt(0)
		}
	}
	return result
}

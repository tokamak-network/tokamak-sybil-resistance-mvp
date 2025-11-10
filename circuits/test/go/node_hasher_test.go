package nodehasher

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"
)

const (
	// MAX_DEG matches the JavaScript test configuration: 14 + 15*3 = 59
	MAX_DEG = 59
)

var (
	PAD_LEN = CalculatePadLen(MAX_DEG)
)

// TestVector represents a single test case from the JSON file
type TestVector struct {
	Name         string  `json:"name"`
	V            int64   `json:"v"`
	D            int64   `json:"d"`
	Neighbors    []int64 `json:"neighbors"`
	ExpectedHash string  `json:"expectedHash"`
}

// TestVectors represents the entire test vector JSON structure
type TestVectors struct {
	Metadata struct {
		MaxDeg      int    `json:"maxDeg"`
		PadLen      int    `json:"padLen"`
		Description string `json:"description"`
		GeneratedAt string `json:"generatedAt"`
	} `json:"metadata"`
	TestCases []TestVector `json:"testCases"`
}

// loadTestVectors loads test vectors from JSON file
func loadTestVectors(t *testing.T) *TestVectors {
	// Load from ../data/nodeHasherTestVectors.json
	testVectorPath := filepath.Join("..", "data", "nodeHasherTestVectors.json")

	data, err := os.ReadFile(testVectorPath)
	if err != nil {
		t.Fatalf("Failed to read test vectors file: %v\nPlease run: cd ../scripts && node generateTestVectors.js", err)
	}

	var vectors TestVectors
	if err := json.Unmarshal(data, &vectors); err != nil {
		t.Fatalf("Failed to parse test vectors JSON: %v", err)
	}

	return &vectors
}

// TestNodeHasherWithTestVectors runs all test cases from the generated JSON file
func TestNodeHasherWithTestVectors(t *testing.T) {
	vectors := loadTestVectors(t)

	t.Logf("Loaded test vectors from: %s", vectors.Metadata.GeneratedAt)
	t.Logf("MAX_DEG = %d, PAD_LEN = %d", vectors.Metadata.MaxDeg, vectors.Metadata.PadLen)
	t.Logf("Running %d test cases...\n", len(vectors.TestCases))

	// Verify metadata matches
	if vectors.Metadata.MaxDeg != MAX_DEG {
		t.Fatalf("Test vector MAX_DEG (%d) does not match Go implementation (%d)",
			vectors.Metadata.MaxDeg, MAX_DEG)
	}
	if vectors.Metadata.PadLen != PAD_LEN {
		t.Fatalf("Test vector PAD_LEN (%d) does not match Go implementation (%d)",
			vectors.Metadata.PadLen, PAD_LEN)
	}

	// Run each test case
	for i, testCase := range vectors.TestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			neighbors := PadNeighbors(testCase.Neighbors, PAD_LEN)

			hash, err := ComputeNodeHash(testCase.V, testCase.D, neighbors, MAX_DEG)
			if err != nil {
				t.Fatalf("ComputeNodeHash failed: %v", err)
			}

			if hash.String() != testCase.ExpectedHash {
				t.Errorf("Hash mismatch for test case #%d (%s):\nv=%d, d=%d\ngot:  %s\nwant: %s",
					i+1, testCase.Name, testCase.V, testCase.D, hash.String(), testCase.ExpectedHash)
			} else {
				t.Logf("✓ Test case #%d (%s) passed - v=%d, d=%d",
					i+1, testCase.Name, testCase.V, testCase.D)
			}
		})
	}
}

// TestCalculatePadLen tests the pad length calculation
func TestCalculatePadLen(t *testing.T) {
	tests := []struct {
		maxDeg   int
		expected int
	}{
		{1, 14},
		{14, 14},
		{15, 29},  // 14 + 15*1
		{29, 29},  // 14 + 15*1
		{30, 44},  // 14 + 15*2
		{59, 59},  // 14 + 15*3
		{60, 74},  // 14 + 15*4
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("maxDeg=%d", tt.maxDeg), func(t *testing.T) {
			result := CalculatePadLen(tt.maxDeg)
			if result != tt.expected {
				t.Errorf("CalculatePadLen(%d) = %d, want %d", tt.maxDeg, result, tt.expected)
			}
		})
	}
}

// TestNodeHasherNotAscending tests that non-ascending neighbors are rejected
func TestNodeHasherNotAscending(t *testing.T) {
	v := int64(123)
	d := int64(3)
	neighbors := PadNeighbors([]int64{1, 81, 3}, PAD_LEN) // Not sorted!

	_, err := ComputeNodeHash(v, d, neighbors, MAX_DEG)
	if err == nil {
		t.Fatal("Expected error for non-ascending neighbors, got nil")
	}

	t.Logf("✓ Correctly rejected non-ascending neighbors: %v", err)
}

// TestNodeHasherInvalidPadding tests that improper padding is rejected
func TestNodeHasherInvalidPadding(t *testing.T) {
	v := int64(10)
	d := int64(2)

	// Create neighbors with invalid padding (non-zero after degree)
	neighbors := make([]*big.Int, PAD_LEN)
	neighbors[0] = big.NewInt(5)
	neighbors[1] = big.NewInt(10)
	neighbors[2] = big.NewInt(15) // Should be 0 since d=2
	for i := 3; i < PAD_LEN; i++ {
		neighbors[i] = big.NewInt(0)
	}

	_, err := ComputeNodeHash(v, d, neighbors, MAX_DEG)
	if err == nil {
		t.Fatal("Expected error for invalid padding, got nil")
	}

	t.Logf("✓ Correctly rejected invalid padding: %v", err)
}

// TestNodeHasherExceedMaxDeg tests that degree exceeding maxDeg is handled
func TestNodeHasherExceedMaxDeg(t *testing.T) {
	v := int64(123)
	d := int64(MAX_DEG + 1) // Exceeds maximum degree

	// Create neighbors array with correct padLen but degree too high
	neighbors := make([]*big.Int, PAD_LEN)
	for i := 0; i < PAD_LEN; i++ {
		neighbors[i] = big.NewInt(0)
	}

	_, err := ComputeNodeHash(v, d, neighbors, MAX_DEG)
	if err == nil {
		t.Fatal("Expected error for degree exceeding maxDeg, got nil")
	}

	t.Logf("✓ Correctly rejected degree exceeding maxDeg: %v", err)
}

// TestNodeHasherWrongPadLen tests that incorrect array length is rejected
func TestNodeHasherWrongPadLen(t *testing.T) {
	v := int64(10)
	d := int64(5)

	// Create neighbors array with wrong length
	neighbors := make([]*big.Int, 10) // Wrong length!
	for i := 0; i < 10; i++ {
		neighbors[i] = big.NewInt(int64(i))
	}

	_, err := ComputeNodeHash(v, d, neighbors, MAX_DEG)
	if err == nil {
		t.Fatal("Expected error for wrong array length, got nil")
	}

	t.Logf("✓ Correctly rejected wrong array length: %v", err)
}

// TestValidateNeighbors tests the neighbor validation function
func TestValidateNeighbors(t *testing.T) {
	tests := []struct {
		name      string
		neighbors []*big.Int
		d         int
		wantErr   bool
	}{
		{
			name:      "valid ascending",
			neighbors: []*big.Int{big.NewInt(1), big.NewInt(3), big.NewInt(5), big.NewInt(0), big.NewInt(0)},
			d:         3,
			wantErr:   false,
		},
		{
			name:      "not strictly ascending",
			neighbors: []*big.Int{big.NewInt(1), big.NewInt(3), big.NewInt(3), big.NewInt(0), big.NewInt(0)},
			d:         3,
			wantErr:   true,
		},
		{
			name:      "invalid padding",
			neighbors: []*big.Int{big.NewInt(1), big.NewInt(3), big.NewInt(5), big.NewInt(7), big.NewInt(0)},
			d:         3,
			wantErr:   true,
		},
		{
			name:      "all zeros with d=0",
			neighbors: []*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
			d:         0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateNeighbors(tt.neighbors, tt.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateNeighbors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// BenchmarkNodeHasherDegree59 benchmarks the hash computation for maximum degree
func BenchmarkNodeHasherDegree59(b *testing.B) {
	v := int64(0)
	d := int64(59)

	neighborsSlice := make([]int64, 59)
	for i := 0; i < 59; i++ {
		neighborsSlice[i] = int64(i + 1)
	}
	neighbors := PadNeighbors(neighborsSlice, PAD_LEN)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ComputeNodeHash(v, d, neighbors, MAX_DEG)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkNodeHasherDegree14 benchmarks the hash computation for first block only
func BenchmarkNodeHasherDegree14(b *testing.B) {
	v := int64(5)
	d := int64(14)
	neighbors := PadNeighbors([]int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28}, PAD_LEN)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ComputeNodeHash(v, d, neighbors, MAX_DEG)
		if err != nil {
			b.Fatal(err)
		}
	}
}

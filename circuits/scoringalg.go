package main

import (
	"fmt"
	"math"
)

const (
	n       = 10
	alpha   = 0.5
	maxIter = 100
	req_bal = 9
	cutoff  = 6
)

func main() {
	// Test data for vouch_matrix, balances and old_scores. The sizes must be nxn, n and n.
	vouch_matrix := [][]float64{
		{0, 0, 1, 1, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 0, 1, 0, 0, 0},
		{1, 1, 1, 0, 0, 1, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 1, 1, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0},
	}
	balances := []float64{10, 3, 10, 10, 20, 10, 20, 30, 10, 10}
	old_scores := []float64{1, 1, 1, 1, 1, 1, 1, 0, 0, 0}

	// Compute the new_scores

	active_nodes := 0
	for _, val := range balances {
		if val >= req_bal {
			active_nodes++
		}
	}
	new_scores := make([]float64, n)
	if active_nodes < cutoff {
		for i, val := range balances {
			if val >= req_bal {
				new_scores[i] = 1.0 / float64(active_nodes)
			}
		}
	} else {
		A := computeA(vouch_matrix, balances)
		W := computeW(A)
		P := computeP(W, alpha, maxIter)
		Q := computeQ(P)
		J := computeJ(Q)
		s := computeS(J, old_scores)
		new_scores = computeY(J, s, W)
	}

	fmt.Println("New Scores:")
	for i, val := range new_scores {
		fmt.Printf("Node %d: %.4f\n", i, val)
	}
}

func computeA(vouchMatrix [][]float64, balances []float64) [][]float64 {
	A := make([][]float64, n)
	for i := range A {
		A[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if vouchMatrix[i][j] == 1 &&
				vouchMatrix[j][i] == 1 &&
				balances[i] >= 9 &&
				balances[j] >= 9 {
				A[i][j] = 1
			}
		}
	}
	return A
}

func computeW(A [][]float64) [][]float64 {
	W := make([][]float64, n)
	for i := range W {
		W[i] = make([]float64, n)
		var d float64
		for j := 0; j < n; j++ {
			d += A[i][j]
		}
		for j := 0; j < n; j++ {
			delta := 0.0
			if i == j {
				delta = 1
			}
			if d != 0 {
				W[i][j] = 0.5 * (delta + A[i][j]/d)
			} else {
				W[i][j] = 0.5 * delta
			}
		}
	}
	return W
}

func computeP(W [][]float64, alpha float64, maxIter int) [][]float64 {
	P := make([][]float64, n)
	for i := 0; i < n; i++ {
		P[i] = make([]float64, n)
		s := make([]float64, n)
		s[i] = 1
		power := append([]float64{}, s...)
		scale := alpha
		copy(P[i], scaleVec(power, scale))
		for t := 1; t < maxIter; t++ {
			power = matVecMul(power, W)
			scale *= (1 - alpha)
			P[i] = addVec(P[i], scaleVec(power, scale))
		}
	}
	return P
}

func computeQ(P [][]float64) [][]int {
	Q := make([][]int, n)
	for i := 0; i < n; i++ {
		idx := make([]int, n)
		for j := 0; j < n; j++ {
			idx[j] = j
		}
		// Sort indices by descending value in P[i]
		for j := 0; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if P[i][idx[j]] < P[i][idx[k]] {
					idx[j], idx[k] = idx[k], idx[j]
				}
			}
		}
		Q[i] = idx
	}
	return Q
}

func computeJ(Q [][]int) [][][]int {
	J := make([][][]int, n)
	for i := 0; i < n; i++ {
		J[i] = make([][]int, n)
		for k := 0; k < n; k++ {
			J[i][k] = make([]int, n)
			J[i][k][i] = 1
			for m := 0; m < k; m++ {
				J[i][k][Q[i][m]] = 1
			}
		}
	}
	return J
}

func computeS(J [][][]int, old_scores []float64) [][]int {
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, n)
		for k := 0; k < n; k++ {
			var sumJ, sumC float64
			for j := 0; j < n; j++ {
				if J[i][k][j] == 1 {
					sumJ += old_scores[j]
				} else {
					sumC += old_scores[j]
				}
			}
			if sumJ <= sumC {
				s[i][k] = 1
			}
		}
	}
	return s
}

func computeY(J [][][]int, s [][]int, W [][]float64) []float64 {
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		y[i] = math.Inf(1)
		for k := 0; k < n; k++ {
			if s[i][k] == 1 {
				jVec := J[i][k]
				jComp := make([]float64, n)
				for m := 0; m < n; m++ {
					jComp[m] = 1 - float64(jVec[m])
				}
				num := dotVec(matVecMul(jComp, W), toFloat64(jVec))
				count := float64(sum(jVec))
				if count > 0 {
					val := num / count
					if val < y[i] {
						y[i] = val
					}
				}
			}
		}
		if math.IsInf(y[i], 1) {
			y[i] = 0.0
		}
	}
	return y
}

func matVecMul(v []float64, M [][]float64) []float64 {
	res := make([]float64, len(v))
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(v); j++ {
			res[i] += v[j] * M[j][i]
		}
	}
	return res
}

func addVec(a, b []float64) []float64 {
	res := make([]float64, len(a))
	for i := range a {
		res[i] = a[i] + b[i]
	}
	return res
}

func scaleVec(v []float64, scale float64) []float64 {
	res := make([]float64, len(v))
	for i := range v {
		res[i] = v[i] * scale
	}
	return res
}

func dotVec(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

func sum(v []int) int {
	s := 0
	for _, x := range v {
		s += x
	}
	return s
}

func toFloat64(v []int) []float64 {
	f := make([]float64, len(v))
	for i, x := range v {
		f[i] = float64(x)
	}
	return f
}

func roundSlice(v []float64, digits int) []float64 {
	factor := math.Pow(10, float64(digits))
	res := make([]float64, len(v))
	for i, val := range v {
		res[i] = math.Round(val*factor) / factor
	}
	return res
}

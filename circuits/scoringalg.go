package main

import (
	"fmt"
	"math"
)

const (
	n       = 10
	alpha   = 0.5
	maxIter = 100
)

func main() {
	edges := [][2]int{
		{0, 1}, {2, 3}, {0, 3}, {0, 2}, {1, 5}, {2, 4},
		{4, 5}, {4, 1}, {6, 7}, {4, 6}, {1, 3}, {3, 5},
		{2, 6}, {7, 9}, {8, 9}, {7, 8},
	}

	A := make([][]float64, n)
	for i := range A {
		A[i] = make([]float64, n)
	}
	for _, e := range edges {
		A[e[0]][e[1]] = 1
		A[e[1]][e[0]] = 1
	}

	x := []float64{1, 1, 1, 1, 1, 1, 1, 0, 0, 0}

	W := computeW(A)
	P := computeP(W, alpha, maxIter)
	Q := computeQ(P)
	J := computeJ(Q)
	s := computeS(J, x)
	y := computeY(J, s, W)

	fmt.Println("New Scores:")
	for i, val := range y {
		fmt.Printf("Node %d: %.4f\n", i, val)
	}
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

func computeS(J [][][]int, x []float64) [][]int {
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, n)
		for k := 0; k < n; k++ {
			var sumJ, sumC float64
			for j := 0; j < n; j++ {
				if J[i][k][j] == 1 {
					sumJ += x[j]
				} else {
					sumC += x[j]
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

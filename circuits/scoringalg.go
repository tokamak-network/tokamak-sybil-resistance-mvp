package main

import (
	"fmt"
	"math"
	"sort"
	"gonum.org/v1/gonum/mat"
)

func main() {
	edges := [][2]int{
		{0, 1}, {2, 3}, {0, 3}, {0, 2}, {1, 5}, {2, 4},
		{4, 5}, {4, 1}, {6, 7}, {4, 6}, {1, 3}, {3, 5},
		{2, 6}, {7, 9}, {8, 9}, {7, 8},
	}

	n := 10
	A := mat.NewDense(n, n, nil)
	for _, e := range edges {
		A.Set(e[0], e[1], 1)
		A.Set(e[1], e[0], 1)
	}

	// Old scores
	x := []float64{1, 1, 1, 1, 1, 1, 1, 0, 0, 0}
	W := computeW(A)
	P := computeP(W, 0.5, 1e-5, 1000)
	Q := computeQ(P)
	J := computeJ(Q)
	s := computeS(J, x)
	y := computeY(J, s, W)

	fmt.Println("Node Scores:")
	for i, v := range y {
		fmt.Printf("Node %d: %.3f\n", i, v)
	}
}

// compute thr transition matrix W
func computeW(A *mat.Dense) *mat.Dense {
	n, _ := A.Dims()
	W := mat.NewDense(n, n, nil)
	for i := 0; i < n; i++ {
		var sum float64
		for j := 0; j < n; j++ {
			sum += A.At(i, j)
		}
		for j := 0; j < n; j++ {
			delta := 0.0
			if i == j {
				delta = 1.0
			}
			val := 0.0
			if sum != 0 {
				val = A.At(i, j) / sum
			}
			W.Set(i, j, 0.5*(delta+val))
		}
	}
	return W
}

func computeP(W *mat.Dense, alpha float64, tol float64, maxIter int) *mat.Dense {
	n, _ := W.Dims()
	P := mat.NewDense(n, n, nil)
	for i := 0; i < n; i++ {
		P.Set(i, i, 1.0)
	}
	I := mat.NewDense(n, n, nil)
	for i := 0; i < n; i++ {
		I.Set(i, i, 1.0)
	}

	tmp := mat.NewDense(n, n, nil)

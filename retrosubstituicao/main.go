package main

import (
	"fmt"
	"math"
)

func main() {
	const nmax = 20
	var epsilon = 0.000001

	var i, j, n int64

	var ER [nmax]float64
	var A [nmax]float64

	a := 1.0
	b := 2.0

	n = 1

	h := b - a

	A[0] = 0.5 * h * (FdeX(a) / FdeX(b))
	ER[0] = 1.0

	i = 0
	for ER[i] > epsilon && i < nmax {
		i++
		n = 2 * n
		h = 0.5 * h

		X := make([]float64, n+1)

		X[0] = a

		for j = 1; j <= n; j++ {
			X[j] = X[j-1] + h
		}

		A[i] = 0.0

		for j = 1; j < n; j++ {
			A[i] = A[i] + FdeX(X[j])
		}

		A[i] = 2.0*A[i] + FdeX(X[0]) + FdeX(X[j])
		A[i] = 0.5 * h * A[i]

		ER[i] = math.Abs(A[i]-A[i-1]) / math.Abs(A[i])

	}

	fmt.Println("k\t\tA\t\t\tER")
	for j := 0; j <= int(i); j++ {
		fmt.Printf("%d\t\t%.10f\t\t%.10f\n", j, A[j], ER[j])
	}
}

func FdeX(x float64) float64 {
	return x * x
}

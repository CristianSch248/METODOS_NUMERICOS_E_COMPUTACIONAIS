package main

import (
	"fmt"
	"math"
)

func main() {
	const nmax = 20
	var epsilon = 0.000001

	var i, j, n int64
	var soma float64

	var ER [nmax]float64
	var A [nmax]float64

	a := 0.0
	b := 3.14159265359

	n = 2
	h := 0.5 * (b - a)

	A[0] = h * (FdeX(a) + FdeX(b) + 4.0*FdeX(0.5*(a+b))) / 3.0
	ER[0] = 1.0

	i = 0
	for ER[i] > epsilon && i < nmax-1 {
		i++
		n = 2 * n
		h = 0.5 * h

		X := make([]float64, n+1)
		X[0] = a

		for j = 1; j <= n; j++ {
			X[j] = X[j-1] + h
		}

		// Calcula A[i]
		A[i] = FdeX(a) + FdeX(b)
		soma = 0.0
		for j = 2; j < n; j += 2 {
			soma += FdeX(X[j])
		}
		A[i] += 2.0 * soma

		soma = 0.0
		for j = 1; j < n; j += 2 {
			soma += FdeX(X[j])
		}
		A[i] += 4.0 * soma

		A[i] = h * A[i] / 3.0
		ER[i] = math.Abs(A[i]-A[i-1]) / math.Abs(A[i])
	}

	// Saída formatada
	fmt.Println("k\t\tA\t\t\tER")
	for j := 0; j <= int(i); j++ {
		fmt.Printf("%d\t\t%.10f\t\t%.10f\n", j, A[j], ER[j])
	}
}

func FdeX(x float64) float64 {
	return math.Sin(x * x * x)
}

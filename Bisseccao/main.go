package main

import (
	"fmt"
	"math"
)

func main() {
	var i, j int

	nmax := 50

	// Criando um slice de float64 com tamanho nmax
	x := make([]float64, nmax)
	a := make([]float64, nmax)
	b := make([]float64, nmax)

	ER := make([]float64, nmax)

	FX := make([]float64, nmax)
	FA := make([]float64, nmax)
	FB := make([]float64, nmax)

	epsilon := 0.001

	a[0] = 0.0
	b[0] = 1.0
	x[0] = 0.5 * (a[0] + b[0])

	FA[0] = f_x(a[0])
	FB[0] = f_x(b[0])
	FX[0] = f_x(x[0])

	ER[0] = 1.0

	for i = 0; ER[i] > epsilon && i < nmax; i++ {

		if FX[i]*FA[i] < 0.0 {
			a[i+1] = a[i]
			b[i+1] = x[i]
		} else {
			a[i+1] = x[i]
			b[i+1] = b[i]
		}

		x[i+1] = 0.5 * (a[i+1] + b[i+1])

		FX[i+1] = f_x(x[i+1])
		FA[i+1] = f_x(a[i+1])
		FB[i+1] = f_x(b[i+1])

		ER[i+1] = math.Abs(x[i+1]-x[i]) / math.Abs(x[i+1])
	}

	// Exibindo os resultados
	fmt.Println("#\t\tA\t\t\tX\t\t\tB\t\t\tFA\t\t\tFX\t\t\tFB\t\t\tErro Relativo")
	for j = 0; j <= i; j++ {
		//                  j     a        x        b        FA       FX       FB       ER
		fmt.Printf("%d\t\t%.10f\t\t%.10f\t\t%.10f\t\t%.10f\t\t%.10f\t\t%.10f\t\t%.10f\n", j, a[j], x[j], b[j], FA[j], FX[j], FB[j], ER[j])
	}

}

func f_x(x float64) float64 {
	return x*x*x - 2.0*x*x - 4.0*x + 4.0
}

// slice := []int{}
// array := [3]int{1, 2, 3}

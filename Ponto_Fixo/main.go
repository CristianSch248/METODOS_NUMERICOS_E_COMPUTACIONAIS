package main

import (
	"fmt"
	"math"
)

func main() {
	var j int

	nmax := 50

	// Criando um slice de float64 com tamanho nmax
	x := make([]float64, nmax)
	fx := make([]float64, nmax)

	epsilon := 0.001

	ER := make([]float64, nmax)

	x[0] = 0.5
	fx[0] = f_x(x[0])
	ER[0] = 1.0

	i := 0
	for ER[i] > epsilon && i < nmax {
		i++

		x[i] = g2(x[i-1])
		fx[i] = f_x(x[i])
		ER[i] = math.Abs(x[i]-x[i-1]) / math.Abs(x[i])
	}

	fmt.Println("#\tx\t\t\tf\t\t\tErro Relativo")
	for j = 0; j <= i; j++ {
		//                  j     x        fx      ER
		fmt.Printf("%d\t%.10f\t\t%.10f\t\t%.10f\n", j, x[j], fx[j], ER[j])
	}

}

func f_x(x float64) float64 {
	return x*x*x - 2.0*x*x - 4.0*x + 4.0
}

func g2(x float64) float64 {
	return 1.0 - 0.5*x*x + 0.25*x*x*x
}

// slice := []int{}
// array := [3]int{1, 2, 3}

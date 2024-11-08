package main

import (
	"fmt"
	"math"
)

func main() {
	const nmax = 20
	const N = 3

	var i int
	var a [N + 1]float64
	var x [nmax]float64
	var ER [nmax]float64
	var b [nmax][N + 1]float64
	var c [nmax][N + 1]float64
	eps := 0.003

	// Inicialização de x[0] e a[i]
	x[0] = 0.5
	a[0] = 4.0
	a[1] = -4.0
	a[2] = -2.0
	a[3] = 1.0

	// Cálculo inicial de b[0][i]
	b[0][N] = a[N]
	for i = N - 1; i >= 0; i-- {
		b[0][i] = a[i] + b[0][i+1]*x[0]
	}

	// Cálculo inicial de c[0][i]
	c[0][N] = a[N]
	for i = N - 1; i > 0; i-- {
		c[0][i] = b[0][i] + c[0][i+1]*x[0]
	}

	// Inicializa ER
	ER[0] = 1.0
	i = 0
	// Critério de parada ER[i] > epsilon e i < nmax
	for ER[i] > eps && i < nmax-1 {
		x[i+1] = x[i] - b[i][0]/c[i][1]
		i++

		ER[i] = math.Abs(x[i]-x[i-1]) / math.Abs(x[i])

		// Recalcular b[i][j]
		b[i][N] = a[N]
		for j := N - 1; j >= 0; j-- {
			b[i][j] = a[j] + b[i][j+1]*x[i]
		}

		// Recalcular c[i][j]
		c[i][N] = a[N]
		for j := N - 1; j > 0; j-- {
			c[i][j] = b[i][j] + c[i][j+1]*x[i]
		}
	}

	// Exibir valores de b[i][k]
	fmt.Println("i\tb3\t\tb2\t\tb1\t\tb0")
	for j := 0; j <= i; j++ {
		fmt.Printf("%d\t", j)
		for k := N; k >= 0; k-- {
			fmt.Printf("%.10f\t", b[j][k])
		}
		fmt.Println()
	}

	// Exibir valores de c[i][k]
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("i\tc3\t\tc2\t\tc1")
	for j := 0; j <= i; j++ {
		fmt.Printf("%d\t", j)
		for k := N; k > 0; k-- {
			fmt.Printf("%.10f\t", c[j][k])
		}
		fmt.Println()
	}

	// Exibir valores de x[i] (tx) e ER[i]
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("j\ttx\t\tER")
	for j := 0; j <= i; j++ {
		// Usar notação científica para valores muito pequenos
		if ER[j] < 1e-6 {
			fmt.Printf("%d\t%.10f\t%.10e\n", j, x[j], ER[j])
		} else {
			fmt.Printf("%d\t%.10f\t%.10f\n", j, x[j], ER[j])
		}
	}
}

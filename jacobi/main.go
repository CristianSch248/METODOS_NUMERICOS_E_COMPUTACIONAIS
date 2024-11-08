package main

import (
	"fmt"
	"math"
)

func main() {
	const N = 5
	const nMax = 50

	var i, j, k int
	var ERMax float64
	var ER [nMax][N]float64
	var x [nMax][N]float64

	epsilon := 0.001

	// Inicialização da matriz estendida (sistema linear)
	a := [N][N]float64{
		{-135.0, -23.0, 63.0, -10.0, -4.0},
		{7.0, -158.0, -39.0, -76.0, -17.0},
		{-50.0, 20.0, -144.0, -6.0, -39.0},
		{6.0, 29.0, 20.0, 122.0, -79.0},
		{-6.0, -14.0, -41.0, 27.0, -123.0},
	}
	b := [N]float64{-10.0, 9.0, -3.0, 8.0, 10.0}

	// Inicializando vetores de x e ER
	for i = 0; i < N; i++ {
		x[0][i] = 0.0
		ER[0][i] = 1.0
	}

	ERMax = 1.0
	i = 0

	// Método de Jacobi
	for ERMax > epsilon && i < nMax {
		for j = 0; j < N; j++ {
			x[i+1][j] = b[j]

			for k = 0; k < j; k++ {
				x[i+1][j] -= a[j][k] * x[i][k]
			}

			for k = j + 1; k < N; k++ {
				x[i+1][j] -= a[j][k] * x[i][k]
			}

			x[i+1][j] /= a[j][j]

			ER[i+1][j] = math.Abs(x[i+1][j]-x[i][j]) / math.Abs(x[i+1][j])
		}

		// Calcula o erro máximo
		ERMax = ER[i+1][0]
		for k = 1; k < N; k++ {
			if ER[i+1][k] > ERMax {
				ERMax = ER[i+1][k]
			}
		}
		i++
	}

	// Exibindo o resultado como uma tabela
	fmt.Println("Método de Jacobi")
	fmt.Printf("k\t\tx(k)1\t\tx(k)2\t\tx(k)3\t\tx(k)4\t\tx(k)5\n")

	for j = 0; j <= i; j++ {
		fmt.Printf("%d\t", j)
		for k = 0; k < N; k++ {
			fmt.Printf("%.7f\t\t", x[j][k])
		}
		fmt.Println()
	}
}

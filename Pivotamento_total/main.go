package main

import (
	"fmt"
	"math"
)

func main() {

	var multi float64
	var i, j, k int
	const N = 4
	var a [N][N + 1]float64
	var x [N]float64
	var perm [N]int

	a[0][0], a[0][1], a[0][2], a[0][3], a[0][4] = -17.0, 19.0, -3.0, 8.0, 86.0
	a[1][0], a[1][1], a[1][2], a[1][3], a[1][4] = 13.0, -10.0, -18.0, -20.0, -297.0
	a[2][0], a[2][1], a[2][2], a[2][3], a[2][4] = -1.0, 15.0, -11.0, 9.0, 64.0
	a[3][0], a[3][1], a[3][2], a[3][3], a[3][4] = -18.0, 2.0, 18.0, -7.0, 6.0

	for i = 0; i < N; i++ {
		perm[i] = i
	}

	var linha, coluna int
	var aux float64

	for i = 0; i < N-1; i++ {
		linha, coluna = i, i
		aux = math.Abs(a[i][i])

		for j = i; j < N; j++ {
			for k = i; k < N; k++ {
				if math.Abs(a[j][k]) > aux {
					aux = math.Abs(a[j][k])
					linha, coluna = j, k
				}
			}
		}

		if linha != i {
			for k = 0; k <= N; k++ {
				aux = a[i][k]
				a[i][k] = a[linha][k]
				a[linha][k] = aux
			}
		}

		if coluna != i {
			for j = 0; j < N; j++ {
				aux = a[j][i]
				a[j][i] = a[j][coluna]
				a[j][coluna] = aux
			}

			aux = float64(perm[i])
			perm[i] = perm[coluna]
			perm[coluna] = int(aux)
		}

		for j = i + 1; j < N; j++ {
			multi = a[j][i] / a[i][i]
			for k = i; k <= N; k++ {
				a[j][k] -= multi * a[i][k]
			}
		}
	}

	x[N-1] = a[N-1][N] / a[N-1][N-1]
	for i = N - 2; i >= 0; i-- {
		x[i] = a[i][N]
		for j = i + 1; j < N; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] = x[i] / a[i][i]
	}

	var solucao [N]float64
	for i = 0; i < N; i++ {
		solucao[perm[i]] = x[i]
	}

	fmt.Println("Matriz após eliminação (com maior precisão):")
	for i := 0; i < N; i++ {
		for j := 0; j <= N; j++ {
			fmt.Printf("%12.8f ", a[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Soluções:")
	for i = 0; i < N; i++ {
		fmt.Printf("x[%d] = %.5f\n", i, solucao[i])
	}
}

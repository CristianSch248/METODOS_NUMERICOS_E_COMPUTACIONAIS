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
	var s [N]float64 // Vetor de escalamento

	// Atribuição de valores
	a[0][0], a[0][1], a[0][2], a[0][3], a[0][4] = -17.0, 19.0, -3.0, 8.0, 86.0
	a[1][0], a[1][1], a[1][2], a[1][3], a[1][4] = 13.0, -10.0, -18.0, -20.0, -297.0
	a[2][0], a[2][1], a[2][2], a[2][3], a[2][4] = -1.0, 15.0, -11.0, 9.0, 64.0
	a[3][0], a[3][1], a[3][2], a[3][3], a[3][4] = -18.0, 2.0, 18.0, -7.0, 6.0

	// Inicializando vetor de escalamento (s)
	for i = 0; i < N; i++ {
		s[i] = math.Abs(a[i][0])
		for j = 1; j < N; j++ {
			if math.Abs(a[i][j]) > s[i] {
				s[i] = math.Abs(a[i][j])
			}
		}
	}

	// eliminação dos elementos abaixo da primeira posição de pivô
	var linha int
	var aux float64

	for i = 0; i < N-1; i++ {
		// Pivotamento parcial escalado
		aux = math.Abs(a[i][i]) / s[i]
		linha = i

		for j = i + 1; j < N; j++ {
			if (math.Abs(a[j][i]) / s[j]) > aux {
				aux = math.Abs(a[j][i]) / s[j]
				linha = j
			}
		}

		// Troca de linhas
		if linha != i {
			for k := 0; k <= N; k++ {
				aux = a[i][k]
				a[i][k] = a[linha][k]
				a[linha][k] = aux
			}
			// Troca também no vetor de escalamento
			aux = s[i]
			s[i] = s[linha]
			s[linha] = aux
		}

		// Eliminação gaussiana
		for j = i + 1; j < N; j++ {
			multi = a[j][i] / a[i][i]
			for k = i; k <= N; k++ {
				a[j][k] -= multi * a[i][k]
			}
		}
	}

	// Substituição regressiva
	x[N-1] = a[N-1][N] / a[N-1][N-1]
	for i = N - 2; i >= 0; i-- {
		x[i] = a[i][N]
		for j = i + 1; j < N; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] = x[i] / a[i][i]
	}

	// Exemplo de saída para verificar se os valores foram atribuídos corretamente
	fmt.Println("Matriz após eliminação (Pivotamento Parcial Escalado):")
	for i := 0; i < N; i++ {
		for j := 0; j <= N; j++ {
			fmt.Printf("%10.5f ", a[i][j]) // Aumenta a precisão para 5 casas decimais
		}
		fmt.Println()
	}

	// Mostrando soluções com mais casas decimais
	fmt.Println("Soluções:")
	for i = 0; i < N; i++ {
		fmt.Printf("x[%d] = %.5f\n", i, x[i]) // Aumenta a precisão para 5 casas decimais
	}
}

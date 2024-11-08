package metodos

import (
	"fmt"
	"math"
)

// PivotamentoParcial aplica o método de eliminação gaussiana com pivotamento parcial
// para resolver um sistema linear representado pela matriz A e o vetor B.
func PivotamentoParcial(A [][]float64, B []float64) {
	N := len(A) // Número de variáveis

	// Processo de eliminação para criar a forma escalonada da matriz
	for i := 0; i < N-1; i++ {
		// Encontra a linha com o maior elemento na coluna atual (pivot)
		linha := i
		aux := math.Abs(A[i][i])

		for j := i + 1; j < N; j++ {
			if math.Abs(A[j][i]) > aux {
				aux = math.Abs(A[j][i])
				linha = j
			}
		}

		// Troca as linhas em A e em B se a linha com o maior elemento não for a linha atual
		if linha != i {
			A[i], A[linha] = A[linha], A[i]
			B[i], B[linha] = B[linha], B[i]
		}

		// Eliminação Gaussiana para zerar os elementos abaixo do pivô na coluna atual
		for j := i + 1; j < N; j++ {
			multi := A[j][i] / A[i][i] // Calcula o multiplicador para a linha
			for k := i; k < N; k++ {
				A[j][k] -= multi * A[i][k] // Subtrai a linha escalonada para criar zeros abaixo do pivô
			}
			B[j] -= multi * B[i] // Ajusta o vetor B de acordo com o multiplicador
		}

		// Impressão da matriz A após cada passo de eliminação
		fmt.Printf("𝐴%d =\n", i+1)
		imprimirMatriz(A, B)
	}

	// Retro-substituição para encontrar os valores das variáveis
	X := make([]float64, N)       // Array para armazenar as soluções
	X[N-1] = B[N-1] / A[N-1][N-1] // Calcula o último valor da variável
	for i := N - 2; i >= 0; i-- {
		X[i] = B[i] // Inicializa o valor da variável com o termo independente
		for j := i + 1; j < N; j++ {
			X[i] -= A[i][j] * X[j] // Subtrai os valores já encontrados para as variáveis superiores
		}
		X[i] /= A[i][i] // Divide pelo coeficiente da diagonal principal
	}

	// Exibe as soluções do sistema em ordem
	fmt.Println("SOLUÇÃO | RETROSUBSTITUIÇÃO")
	for i := N - 1; i >= 0; i-- {
		fmt.Printf("x%d = %.10f ", i+1, X[i]) // Imprime cada solução com 10 dígitos significativos
	}
	fmt.Println()
}

// Função auxiliar para imprimir a matriz A com precisão
func imprimirMatriz(A [][]float64, B []float64) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			fmt.Printf("%12.8f ", A[i][j])
		}
		fmt.Printf("| %12.8f\n", B[i])
	}
	fmt.Println()
}

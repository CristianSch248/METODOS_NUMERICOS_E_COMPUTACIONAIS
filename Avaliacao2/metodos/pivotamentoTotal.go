package metodos

import (
	"fmt"
	"math"
)

// PivotamentoTotal aplica o método de eliminação gaussiana com pivotamento total
// para resolver um sistema linear representado pela matriz A e o vetor B.
func PivotamentoTotal(A [][]float64, B []float64) {
	N := len(A)            // Número de variáveis
	perm := make([]int, N) // Array para armazenar a permutação das colunas
	for i := 0; i < N; i++ {
		perm[i] = i // Inicializa a permutação das colunas para o estado original
	}

	// Processo de eliminação para criar a forma escalonada da matriz
	for i := 0; i < N-1; i++ {
		linha, coluna := i, i
		aux := math.Abs(A[i][i]) // Armazena o valor absoluto do elemento atual da diagonal

		// Encontrar o maior elemento em valor absoluto para fazer o pivotamento total
		for j := i; j < N; j++ {
			for k := i; k < N; k++ {
				if math.Abs(A[j][k]) > aux {
					aux = math.Abs(A[j][k]) // Atualiza o maior valor
					linha, coluna = j, k    // Atualiza os índices da linha e coluna do maior valor
				}
			}
		}

		// Troca de linhas, se necessário, para trazer o maior valor para a posição diagonal
		if linha != i {
			A[i], A[linha] = A[linha], A[i]
			B[i], B[linha] = B[linha], B[i]
		}

		// Troca de colunas, se necessário, e ajusta a permutação para registrar essa troca
		if coluna != i {
			for j := 0; j < N; j++ {
				A[j][i], A[j][coluna] = A[j][coluna], A[j][i]
			}
			perm[i], perm[coluna] = perm[coluna], perm[i] // Registra a troca na permutação
		}

		// Aplicação da eliminação Gaussiana para zerar os elementos abaixo do pivô
		for j := i + 1; j < N; j++ {
			multi := A[j][i] / A[i][i] // Calcula o multiplicador para a linha
			for k := i; k < N; k++ {
				A[j][k] -= multi * A[i][k] // Subtrai a linha escalonada
			}
			B[j] -= multi * B[i] // Ajusta o vetor B de acordo com o multiplicador
		}

		// Impressão da matriz A após cada passo de eliminação
		fmt.Printf("𝐴%d =\n", i+1)
		imprimirMatrizPivotamentoTotal(A, B)
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

	// Reordenar as soluções de acordo com a permutação das colunas
	solucao := make([]float64, N)
	for i := 0; i < N; i++ {
		solucao[perm[i]] = X[i] // Armazena a solução na ordem correta
	}

	// Exibe a matriz final após a eliminação
	fmt.Println("Matriz após eliminação (com maior precisão):")
	imprimirMatrizPivotamentoTotal(A, B)

	// Exibe as soluções do sistema na ordem correta
	fmt.Println("SOLUÇÃO | RETROSUBSTITUIÇÃO")
	for i := N - 1; i >= 0; i-- {
		fmt.Printf("x[%d] = %.10f ", perm[i]+1, solucao[perm[i]]) // Imprime cada solução com 10 dígitos significativos
	}
	fmt.Println()
}

// Função auxiliar para imprimir a matriz A com precisão
func imprimirMatrizPivotamentoTotal(A [][]float64, B []float64) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			fmt.Printf("%12.10f ", A[i][j])
		}
		fmt.Printf("| %12.10f\n", B[i])
	}
	fmt.Println()
}

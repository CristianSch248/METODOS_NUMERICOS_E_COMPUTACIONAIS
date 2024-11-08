package metodos

import (
	"fmt"
	"math"
)

// Jacobi resolve o sistema linear C * x = D utilizando o Método de Jacobi.
func Jacobi(C [][]float64, D []float64) {
	const N = 7           // Tamanho do sistema (número de variáveis)
	const nMax = 50       // Número máximo de iterações
	const epsilon = 0.001 // Critério de parada para o erro relativo

	var i, j, k int
	var ERMax float64
	x := make([]float64, N)    // Solução iterativa
	xOld := make([]float64, N) // Solução anterior para comparação do erro

	// Inicializando vetor de solução com valores iniciais
	for i = 0; i < N; i++ {
		x[i] = 0.0 // Chute inicial para x
	}

	fmt.Println("k\t\tx1,k\tx2,k\tx3,k\tx4,k\tx5,k\tx6,k\tx7,k")

	ERMax = 1.0 // Inicializa o erro máximo acima do epsilon para começar o loop
	i = 0       // Contador de iterações

	// Método de Jacobi
	for ERMax > epsilon && i < nMax {
		// Copia os valores de x para xOld antes da próxima iteração
		copy(xOld, x)

		// Iteração de Jacobi para cada variável
		for j = 0; j < N; j++ {
			x[j] = D[j] // Termo independente
			for k = 0; k < N; k++ {
				if j != k {
					x[j] -= C[j][k] * xOld[k] // Soma dos termos não-diagonais
				}
			}
			x[j] /= C[j][j] // Divisão pelo coeficiente diagonal
		}

		// Exibe valores de x para a iteração atual
		fmt.Printf("%d\t", i)
		for j = 0; j < N; j++ {
			fmt.Printf("%.10f\t", x[j])
		}
		fmt.Println()

		// Calcula o erro máximo relativo entre x e xOld
		ERMax = 0.0
		for j = 0; j < N; j++ {
			erro := math.Abs(x[j]-xOld[j]) / math.Abs(x[j])
			if erro > ERMax {
				ERMax = erro
			}
		}
		i++ // Incrementa o contador de iterações
	}

	// Exibindo erro relativo para cada variável
	fmt.Println("k\t\tER1,k\tER2,k\tER3,k\tER4,k\tER5,k\tER6,k\tER7,k")
	for j = 0; j < N; j++ {
		if i == 1 {
			fmt.Print("1\t1.0000000\t1.0000000\t1.0000000\t1.0000000\t1.0000000\t1.0000000\t1.0000000\n")
		} else {
			fmt.Printf("%d\t", j)
			for k = 0; k < N; k++ {
				erro := math.Abs(x[k]-xOld[k]) / math.Abs(x[k])
				fmt.Printf("%.10f\t", erro)
			}
			fmt.Println()
		}
	}
	fmt.Println()
}

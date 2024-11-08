package metodos

import (
	"fmt"
	"math"
)

// GaussSeidel resolve o sistema linear C * x = D utilizando o Método de Gauss-Seidel.
func GaussSeidel(C [][]float64, D []float64) {
	const N = 7           // Tamanho do sistema (número de variáveis)
	const nMax = 50       // Número máximo de iterações
	const epsilon = 0.001 // Critério de parada para o erro relativo

	var i, j, k int
	var ERMax float64
	x := make([]float64, N)    // Solução iterativa
	xOld := make([]float64, N) // Solução anterior para comparação do erro

	// Inicializando o vetor de solução com valores iniciais
	for i = 0; i < N; i++ {
		x[i] = 0.0 // Chute inicial para x
	}

	fmt.Println("k\t\tx1,k\tx2,k\tx3,k\tx4,k\tx5,k\tx6,k\tx7,k")

	ERMax = 1.0 // Inicializa o erro máximo acima do epsilon para iniciar o loop
	i = 0       // Contador de iterações

	// Método de Gauss-Seidel
	for ERMax > epsilon && i < nMax {
		// Copia os valores de x para xOld antes da próxima iteração
		copy(xOld, x)

		// Iteração de Gauss-Seidel para cada variável
		for j = 0; j < N; j++ {
			x[j] = D[j] // Termo independente
			// Utiliza os valores mais atualizados de x em cada passo
			for k = 0; k < N; k++ {
				if j != k {
					if k < j {
						x[j] -= C[j][k] * x[k]
					} else {
						x[j] -= C[j][k] * xOld[k]
					}
				}
			}
			x[j] /= C[j][j] // Divisão pelo coeficiente da diagonal
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
	for iter := 0; iter < i; iter++ {
		if iter == 0 {
			fmt.Print("0\t\t-  \t-  \t-  \t-  \t-  \t-  \t-\n")
		} else {
			fmt.Printf("%d\t\t", iter)
			for j = 0; j < N; j++ {
				erro := math.Abs(x[j]-xOld[j]) / math.Abs(x[j])
				fmt.Printf("%.10f\t", erro)
			}
			fmt.Println()
		}
	}
	fmt.Println("Número total de iterações:", i)
}

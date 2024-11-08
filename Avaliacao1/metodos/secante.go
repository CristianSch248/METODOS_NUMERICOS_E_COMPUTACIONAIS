package metodos

import (
	"fmt"
	"math"
)

// Função para o Método da Secante
func Secant(x0, x1, tol float64, nmax int) {
	var i int
	X := make([]float64, nmax)
	FX := make([]float64, nmax)
	ER := make([]float64, nmax)

	X[0] = x0
	X[1] = x1
	FX[0] = secantf_x(X[0])
	FX[1] = secantf_x(X[1])

	fmt.Println("n\txn\t\t\tf(xn)\t\t\tERn")

	for i = 1; i < nmax && math.Abs(FX[i]) > tol; i++ {
		if FX[i] == FX[i-1] {
			fmt.Println("Erro: Divisão por zero iminente.")
			break
		}

		// Fórmula da secante
		X[i+1] = X[i] - FX[i]*(X[i]-X[i-1])/(FX[i]-FX[i-1])
		FX[i+1] = secantf_x(X[i+1])
		ER[i+1] = math.Abs(X[i+1]-X[i]) / math.Abs(X[i+1])

		fmt.Printf("%d\t%.10e\t%.10e\t%.10e\n", i, X[i], FX[i], ER[i])
	}

	// Exibir o resultado final
	fmt.Printf("%d\t%.10e\t%.10e\t%.10e\n", i, X[i], FX[i], ER[i])

	// Exibir a raiz aproximada encontrada
	fmt.Printf("\nRaiz aproximada encontrada: x = %.10e, após %d iterações\n", X[i], i)
}

// Função polinomial P(x) para Secante
func secantf_x(x float64) float64 {
	return math.Pow(x, 5) - 54.9493*math.Pow(x, 4) + 1199.71*math.Pow(x, 3) - 13007.3*math.Pow(x, 2) + 70025.1*x - 149746
}

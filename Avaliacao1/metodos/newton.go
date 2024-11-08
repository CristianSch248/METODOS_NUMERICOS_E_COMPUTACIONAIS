package metodos

import (
	"fmt"
	"math"
)

// Função para o Método de Newton
func Newton(x0 float64, nmax int, tol float64) {
	var i int
	X := make([]float64, nmax)
	FX := make([]float64, nmax)
	DER := make([]float64, nmax)
	ER := make([]float64, nmax)

	X[0] = x0
	FX[0] = newtonf_x(X[0])
	DER[0] = df(X[0])

	fmt.Println("k\txk\t\t\tf(xk)\t\t\tf'(xk)\t\t\tERk")

	// Critério de parada: erro relativo menor que a tolerância
	for i = 0; i < nmax-1 && math.Abs(FX[i]) > tol && (i == 0 || ER[i-1] > tol); i++ {
		if DER[i] == 0.0 {
			fmt.Println("Erro: Derivada igual a zero.")
			break
		}
		// Verificação para evitar derivada muito próxima de zero
		if math.Abs(DER[i]) < 1e-10 {
			fmt.Println("Erro: Derivada muito próxima de zero.")
			break
		}

		X[i+1] = X[i] - FX[i]/DER[i]
		FX[i+1] = newtonf_x(X[i+1])
		DER[i+1] = df(X[i+1])
		ER[i+1] = math.Abs(X[i+1]-X[i]) / math.Abs(X[i+1])

		fmt.Printf("%d\t%.10e\t%.10e\t%.10e\t%.10e\n", i, X[i], FX[i], DER[i], ER[i])
	}

	fmt.Printf("%d\t%.10e\t%.10e\t%.10e\t%.10e\n", i, X[i], FX[i], DER[i], ER[i])

	if math.Abs(FX[i]) <= tol {
		fmt.Printf("\nRaiz aproximada encontrada: x = %.10e, com erro relativo = %.10e, após %d iterações\n", X[i], ER[i], i)
	} else {
		fmt.Println("\nO método não convergiu dentro do número máximo de iterações.")
	}
}

// Função polinomial P(x) para Newton
func newtonf_x(x float64) float64 {
	return math.Pow(x, 5) - 54.9493*math.Pow(x, 4) + 1199.71*math.Pow(x, 3) - 13007.3*math.Pow(x, 2) + 70025.1*x - 149746
}

// Derivada de P(x)
func df(x float64) float64 {
	return 5*math.Pow(x, 4) - 4*54.9493*math.Pow(x, 3) + 3*1199.71*math.Pow(x, 2) - 2*13007.3*x + 70025.1
}

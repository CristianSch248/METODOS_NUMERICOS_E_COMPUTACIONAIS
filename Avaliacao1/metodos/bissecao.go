package metodos

import (
	"fmt"
	"math"
)

// Função para o Método da Bisseção
func Bisection(a, b, tol float64, nmax int) {
	var i int
	x := make([]float64, nmax)
	ER := make([]float64, nmax)
	FA := make([]float64, nmax)
	FB := make([]float64, nmax)
	FX := make([]float64, nmax)

	// Definindo o intervalo inicial
	x[0] = 0.5 * (a + b)
	FA[0] = f_x(a)
	FB[0] = f_x(b)
	FX[0] = f_x(x[0])
	ER[0] = 1.0

	// Loop para o método da bissecção
	for i = 0; ER[i] > tol && i < nmax; i++ {
		// Atualizando o intervalo
		if FX[i]*FA[i] < 0.0 {
			b = x[i]
		} else {
			a = x[i]
		}

		x[i+1] = 0.5 * (a + b)

		// Calculando novos valores da função
		FX[i+1] = f_x(x[i+1])
		FA[i+1] = f_x(a)
		FB[i+1] = f_x(b)

		// Calculando o erro relativo
		ER[i+1] = math.Abs(x[i+1]-x[i]) / math.Abs(x[i+1])
	}

	// Exibindo os resultados em uma tabela com 10 dígitos significativos
	fmt.Println("k\t\tA\t\t\tX\t\t\tB\t\t\tFA\t\t\tFX\t\t\tFB\t\tErro Relativo")
	for j := 0; j <= i; j++ {
		fmt.Printf("%d\t%.10e\t%.10e\t%.10e\t%.10e\t%.10e\t%.10e\t%.10e\n", j, a, x[j], b, FA[j], FX[j], FB[j], ER[j])
	}

	// Exibir a raiz aproximada encontrada
	fmt.Printf("\nRaiz aproximada encontrada: x = %.10e, após %d iterações\n", x[i], i)
}

// Função polinomial P(x) para Bisseção
func f_x(x float64) float64 {
	return math.Pow(x, 5) - 54.9493*math.Pow(x, 4) + 1199.71*math.Pow(x, 3) - 13007.3*math.Pow(x, 2) + 70025.1*x - 149746
}

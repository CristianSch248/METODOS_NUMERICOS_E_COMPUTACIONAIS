package main

import (
	"fmt"
	"math"
)

func main() {
	// Parâmetros
	var i int
	nmax := 50
	epsilon := 1e-10

	// Chute inicial
	x0 := 0.5

	// Vetores para armazenar valores
	X := make([]float64, nmax)
	FX := make([]float64, nmax)
	DER := make([]float64, nmax)
	ER := make([]float64, nmax)

	// Definindo o valor inicial
	X[0] = x0
	FX[0] = f(X[0])
	DER[0] = de(X[0])
	ER[0] = math.NaN() // Não há erro relativo na primeira iteração

	// Exibindo o cabeçalho
	fmt.Println("n\t\txn\t\t\tf(xn)\t\t\tf'(xn)\t\t\tERn")

	// Iteração do método de Newton
	for i = 0; i < nmax-1 && math.Abs(FX[i]) > epsilon; i++ {
		// Evitar divisão por zero
		if DER[i] == 0.0 {
			fmt.Println("Erro: Derivada igual a zero.")
			break
		}

		// Fórmula do Método de Newton: x_{n+1} = x_n - f(x_n) / f'(x_n)
		X[i+1] = X[i] - FX[i]/DER[i]

		// Calcular f(x), derivada f'(x) e o erro relativo
		FX[i+1] = f(X[i+1])
		DER[i+1] = de(X[i+1])
		ER[i+1] = math.Abs(X[i+1]-X[i]) / math.Abs(X[i+1])

		// Exibir os resultados da iteração
		fmt.Printf("%d\t\t%.14f\t\t%.10f\t\t%.5f\t\t%.10f\n", i, X[i], FX[i], DER[i], ER[i])
	}

	// Exibir o último valor
	fmt.Printf("%d\t\t%.14f\t\t%.10f\t\t%.5f\t\t%.10f\n", i, X[i], FX[i], DER[i], ER[i])

	// Exibir o resultado final
	fmt.Printf("\nRaiz aproximada encontrada: X = %.14f, f(X) = %.10f, Iterações = %d\n", X[i], FX[i], i)
}

// Função f(x) - Função cujas raízes estamos encontrando
func f(x float64) float64 {
	return x*x*x - 2*x*x - 4*x + 4
}

// Derivada de f(x), f'(x)
func de(x float64) float64 {
	return 3*x*x - 4*x - 4
}

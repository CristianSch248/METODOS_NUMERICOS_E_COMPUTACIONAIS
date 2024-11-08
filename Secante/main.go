package main

import (
	"fmt"
	"math"
)

func main() {
	// Parâmetros
	var i int
	nmax := 50
	epsilon := 0.0000001

	// Estimativas iniciais
	x0 := 0.5
	x1 := 0.75

	// Vetores para armazenar valores
	X := make([]float64, nmax)
	FX := make([]float64, nmax)
	ER := make([]float64, nmax)

	// Definindo os valores iniciais
	X[0] = x0
	X[1] = x1
	FX[0] = f(X[0])
	FX[1] = f(X[1])
	ER[0] = math.NaN() // O erro relativo da primeira iteração é indefinido

	// Exibindo o cabeçalho no formato solicitado
	fmt.Println("n\t\txn\t\t\tf(xn)\t\t\tERn")

	// Iteração do método da secante
	for i = 1; i < nmax && math.Abs(FX[i]) > epsilon; i++ {
		// Evitar divisão por zero
		if FX[i] == FX[i-1] {
			fmt.Println("Erro: Divisão por zero iminente.")
			break
		}

		// Fórmula do método da secante: x_{n+1} = x_n - f(x_n) * (x_n - x_{n-1}) / (f(x_n) - f(x_{n-1}))
		X[i+1] = X[i] - FX[i]*(X[i]-X[i-1])/(FX[i]-FX[i-1])

		// Calcular f(x) e o erro relativo
		FX[i+1] = f(X[i+1])
		ER[i+1] = math.Abs(X[i+1]-X[i]) / math.Abs(X[i+1])

		// Exibir valores de cada iteração no formato solicitado
		fmt.Printf("%d\t\t%.15f\t\t%.10f\t\t%.10f\n", i, X[i], FX[i], ER[i])
	}

	// Exibir o último valor após a convergência
	fmt.Printf("%d\t\t%.15f\t\t%.10f\t\t%.10f\n", i, X[i], FX[i], ER[i])

	// Exibir o resultado final
	fmt.Printf("\nRaiz aproximada encontrada: X = %.15f, f(X) = %.10f, Iterações = %d\n", X[i], FX[i], i)
}

// Função f(x) - Função alvo cujas raízes estamos encontrando
func f(x float64) float64 {
	return x*x*x - 2*x*x - 4*x + 4
}

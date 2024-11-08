package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Método da Secante")
	secant(0.0, 1.0, 0.0000001, 50)
}

// Função para realizar o método da Secante
func secant(a, b, tol float64, nmax int) {
	var i int
	var xn, fa, fb, fxn, ER float64

	// Inicializar as estimativas iniciais
	fa = f(a)
	fb = f(b)

	// Exibindo o cabeçalho
	fmt.Println("n\t\tan\t\t\t\txn\t\t\t\tbn\t\t\t\tf(xn)\t\t\tERn")

	// Iterações do método da Secante
	for i = 1; i <= nmax; i++ {
		// Calcular a próxima aproximação usando a fórmula da Secante
		xn = b - (fb*(b-a))/(fb-fa)
		fxn = f(xn)

		// Calcular o erro relativo
		ER = math.Abs(xn-b) / math.Abs(xn)

		// Exibir os resultados
		fmt.Printf("%d\t\t%.10f\t\t%.10f\t\t%.10f\t\t%.10f\t\t%.10f\n", i, a, xn, b, fxn, ER)

		// Verificar se atingiu a tolerância ou se f(xn) == 0 (raiz exata)
		if math.Abs(fxn) < tol || ER < tol {
			break
		}

		// Atualizar as variáveis para a próxima iteração
		a = b
		fa = fb
		b = xn
		fb = fxn
	}

	fmt.Printf("\nRaiz aproximada encontrada: X = %.10f, Iterações = %d\n", xn, i)
}

// Função f(x) - Função cujas raízes estamos encontrando
func f(x float64) float64 {
	return x*x*x - 2*x*x - 4*x + 4
}

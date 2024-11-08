package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Método da Bisseção")
	bisection(0.0, 1.0, 0.0000001, 50)
}

// Função para realizar o método da Bisseção
func bisection(a, b, tol float64, nmax int) {
	var i int
	var xn, fa, fb, fxn, ER float64

	// Verificar se f(a) * f(b) < 0, se não, o intervalo não é válido
	fa = f(a)
	fb = f(b)
	if fa*fb > 0 {
		fmt.Println("Erro: O intervalo não contém uma raiz ou contém múltiplas raízes.")
		return
	}

	// Exibindo o cabeçalho
	fmt.Println("n\t\tan\t\t\t\txn\t\t\t\tbn\t\t\t\tf(xn)\t\t\tERn")

	// Iterações do método da Bisseção
	for i = 1; i <= nmax; i++ {
		// Calcular o ponto médio
		xn = (a + b) / 2.0
		fxn = f(xn)

		// Calcular o erro relativo
		ER = math.Abs(b-a) / 2.0

		// Exibir os resultados
		fmt.Printf("%d\t\t%.10f\t\t%.10f\t\t%.10f\t\t%.10f\t\t%.10f\n", i, a, xn, b, fxn, ER)

		// Verificar se atingiu a tolerância ou se f(xn) == 0 (raiz exata)
		if math.Abs(fxn) < tol || ER < tol {
			fmt.Printf("\nConvergiu com erro relativo %.10f e valor de f(xn) %.10f\n", ER, fxn)
			break
		}

		// Atualizar o intervalo com base no sinal de f(xn)
		if fa*fxn < 0 {
			b = xn
			fb = fxn
		} else {
			a = xn
			fa = fxn
		}
	}

	// Exibir a raiz aproximada e o número de iterações
	fmt.Printf("\nRaiz aproximada encontrada: X = %.10f, Iterações = %d\n", xn, i)
}

// Função f(x) - Função alvo cujas raízes estamos encontrando
func f(x float64) float64 {
	return x*x*x - 2*x*x - 4*x + 4
}

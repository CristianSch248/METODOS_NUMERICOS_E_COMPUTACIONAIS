package main

import (
	"fmt"
	"math"
)

// Função que queremos integrar
func f(x float64) float64 {
	return math.Sin(x) // Exemplo: f(x) = sin(x)
}

// Função para calcular a regra do trapézio composta
func trapezoidalComposite(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := f(a) + f(b)

	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		sum += 2 * f(x)
	}

	return (h / 2) * sum
}

// Função para calcular a regra de Simpson composta
func simpsonComposite(a, b float64, n int) float64 {
	// n precisa ser par para a regra de Simpson
	if n%2 != 0 {
		fmt.Println("Erro: o número de subintervalos (n) precisa ser par para a regra de Simpson.")
		return 0.0
	}

	h := (b - a) / float64(n)
	sum := f(a) + f(b)

	// Soma termos com peso 4 e 2 intercaladamente
	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		if i%2 == 0 {
			sum += 2 * f(x)
		} else {
			sum += 4 * f(x)
		}
	}

	return (h / 3) * sum
}

func main() {
	// Intervalo de integração e número de subintervalos
	a, b := 0.0, math.Pi // Exemplo: integração de 0 a pi
	n := 10              // Número de subintervalos (par para Simpson)

	// Calcula a integral usando a regra do trapézio composta
	resultTrapezoidal := trapezoidalComposite(a, b, n)
	fmt.Printf("Integral aproximada pela regra do trapézio composta: %.7f\n", resultTrapezoidal)

	// Calcula a integral usando a regra de Simpson composta
	resultSimpson := simpsonComposite(a, b, n)
	fmt.Printf("Integral aproximada pela regra de Simpson composta: %.7f\n", resultSimpson)
}

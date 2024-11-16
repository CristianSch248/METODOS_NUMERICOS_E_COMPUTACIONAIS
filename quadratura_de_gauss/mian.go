package main

import (
	"fmt"
	"math"
)

// Função que desejamos integrar
func f(x float64) float64 {
	// Exemplo: função f(x) = x^2
	return x * x
}

// Função para calcular a quadratura de Gauss de 2 pontos em um intervalo [a, b]
func gaussQuadrature2Points(a, b float64) float64 {
	// Calcula sqrt(3)/3 uma vez e armazena em uma variável
	sqrt3Div3 := math.Sqrt(3.0) / 3.0

	// Pontos e pesos para a quadratura de Gauss de 2 pontos
	x1 := -sqrt3Div3
	x2 := sqrt3Div3
	c1, c2 := 1.0, 1.0

	// Transformação para o intervalo [a, b]
	mid := (b + a) / 2.0
	halfLength := (b - a) / 2.0

	// Calcula a integral aproximada
	integral := halfLength * (c1*f(mid+halfLength*x1) + c2*f(mid+halfLength*x2))

	return integral
}

func main() {
	// Intervalo de integração
	a := -1.0
	b := 1.0

	// Calcula a integral
	result := gaussQuadrature2Points(a, b)

	// Exibe o resultado
	fmt.Printf("A integral aproximada de f(x) no intervalo [%.1f, %.1f] é %.7f\n", a, b, result)
}

package metodos

import (
	"fmt"
	"math"
)

// Função para o Método da Falsa Posição
func FalsePosition(a, b, tol float64, nmax int) {
	var i int
	var xn, fa, fb, fxn, ER float64

	fa = P(a)
	fb = P(b)

	fmt.Println("k\tak\t\t\txk\t\t\tbk\t\t\tf(ak)\t\t\tf(xk)\t\t\tf(bk)\t\t\tERk")

	for i = 1; i <= nmax; i++ {
		xn = b - fb*(b-a)/(fb-fa)
		fxn = P(xn)
		ER = math.Abs(xn-b) / math.Abs(xn)

		fmt.Printf("%d\t%.10e\t%.10e\t%.10e\t%.10e\t%.10e\t%.10e\t%.10e\n", i, a, xn, b, fa, fxn, fb, ER)

		if math.Abs(fxn) < tol || ER < tol {
			break
		}

		if fa*fxn < 0 {
			b = xn
			fb = fxn
		} else {
			a = xn
			fa = fxn
		}
	}

	fmt.Printf("\nRaiz aproximada encontrada: X = %.10e, Iterações = %d\n", xn, i)
}

// Função polinomial P(x) para Falsa Posição
func P(x float64) float64 {
	return math.Pow(x, 5) - 54.9493*math.Pow(x, 4) + 1199.71*math.Pow(x, 3) - 13007.3*math.Pow(x, 2) + 70025.1*x - 149746
}

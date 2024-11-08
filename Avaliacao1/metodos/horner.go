package metodos

import (
	"fmt"
	"math"
)

func Horner(nmax int, tol float64) {
	const N = 5
	var i int
	var a [N + 1]float64
	var x []float64 = make([]float64, nmax)
	var ER []float64 = make([]float64, nmax)
	var b [][]float64 = make([][]float64, nmax)
	var c [][]float64 = make([][]float64, nmax)

	for k := range b {
		b[k] = make([]float64, N+1)
		c[k] = make([]float64, N+1)
	}

	// Ajuste no valor inicial de x[0] para algo maior que z4
	x[0] = 13.5
	a[5] = 1.0
	a[4] = -54.9493
	a[3] = 1199.71
	a[2] = -13007.3
	a[1] = 70025.1
	a[0] = -149746.0

	// Cálculo inicial de b[0][i]
	b[0][N] = a[N]
	for i = N - 1; i >= 0; i-- {
		b[0][i] = a[i] + b[0][i+1]*x[0]
	}

	// Cálculo inicial de c[0][i]
	c[0][N] = a[N]
	for i = N - 1; i > 0; i-- {
		c[0][i] = b[0][i] + c[0][i+1]*x[0]
	}

	// Inicializa o erro relativo
	ER[0] = 1.0
	i = 0

	// Critério de parada ER[i] > tol e i < nmax-1
	for ER[i] > tol && i < nmax-1 {
		// Atualiza x[i+1]
		x[i+1] = x[i] - b[i][0]/c[i][1]
		i++

		// Calcula o erro relativo
		ER[i] = math.Abs(x[i]-x[i-1]) / math.Abs(x[i])

		// Recalcular b[i][j]
		b[i][N] = a[N]
		for j := N - 1; j >= 0; j-- {
			b[i][j] = a[j] + b[i][j+1]*x[i]
		}

		// Recalcular c[i][j]
		c[i][N] = a[N]
		for j := N - 1; j > 0; j-- {
			c[i][j] = b[i][j] + c[i][j+1]*x[i]
		}
	}

	fmt.Println("\nCOEFICIENTES DO POLINÔMIO PN(x)")
	fmt.Println("k\tb5\t\t\tb4\t\t\tb3\t\t\tb2\t\t\tb1\t\t\tb0")
	for j := 0; j <= i; j++ {
		fmt.Printf("%d\t", j)
		for k := N; k >= 0; k-- {
			fmt.Printf("%.10e\t", b[j][k])
		}
		fmt.Println()
	}

	fmt.Println("\nCOEFICIENTES DO POLINÔMIO P'N(x)")
	fmt.Println("k\tc5\t\t\tc4\t\t\tc3\t\t\tc2\t\t\tc1")
	for j := 0; j <= i; j++ {
		fmt.Printf("%d\t", j)
		for k := N; k > 0; k-- {
			fmt.Printf("%.10e\t", c[j][k])
		}
		fmt.Println()
	}

	fmt.Println("\nESTIMATIVAS")
	fmt.Println("k\txk\t\t\tf(xk)\t\t\tf'(xk)\t\t\tERk")
	for j := 0; j <= i; j++ {
		fmt.Printf("%d\t%.10e\t%.10e\t%.10e\t%.10e\n", j, x[j], b[j][0], c[j][1], ER[j])
	}

	// Exibir a raiz aproximada encontrada
	fmt.Printf("\nRaiz z5 = %.10e, com erro relativo = %.10e, após %d iterações\n", x[i], ER[i], i)
}

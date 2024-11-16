package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// Estrutura para armazenar os pares ordenados
type Point struct {
	x, y float64
}

// Função para calcular as diferenças divididas
func dividedDifferences(points []Point) [][]float64 {
	n := len(points)
	ddTable := make([][]float64, n)

	// Inicializa a tabela com os valores y_k
	for i := 0; i < n; i++ {
		ddTable[i] = make([]float64, n-i) // Corrige o tamanho da linha
		ddTable[i][0] = points[i].y
	}

	// Calcula as diferenças divididas
	for j := 1; j < n; j++ {
		for i := 0; i < n-j; i++ {
			ddTable[i][j] = (ddTable[i+1][j-1] - ddTable[i][j-1]) / (points[i+j].x - points[i].x)
		}
	}

	return ddTable
}

// Função para calcular o polinômio interpolador de Newton no ponto z até uma ordem específica
func newtonInterpolation(points []Point, ddTable [][]float64, z float64, order int) float64 {
	approx := ddTable[0][0]
	product := 1.0

	for i := 1; i <= order; i++ {
		product *= (z - points[i-1].x)
		approx += ddTable[0][i] * product
	}

	return approx
}

// Função para gerar o gráfico
func createChart(points []Point, ddTable [][]float64, order int, color string) *charts.Line {
	xValues := make([]float64, len(points))
	yValues := make([]opts.LineData, len(points))
	originalY := make([]opts.ScatterData, len(points))

	// Calcular os valores do polinômio de ordem especificada
	for i, point := range points {
		xValues[i] = point.x
		yValues[i] = opts.LineData{Value: newtonInterpolation(points, ddTable, point.x, order)}
		originalY[i] = opts.ScatterData{Value: point.y}
	}

	// Criar o gráfico de linha e adicionar os pontos originais como dispersão
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: fmt.Sprintf("y = P_%d(x)", order)}),
		charts.WithColorsOpts(opts.Colors{color}),
	)
	line.SetXAxis(xValues).
		AddSeries(fmt.Sprintf("P_%d(x)", order), yValues).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{}))

	// Adicionar a série de dispersão para os pontos originais
	scatter := charts.NewScatter()
	scatter.SetXAxis(xValues).
		AddSeries("Dados Originais", originalY)

	// Incorporar o gráfico de dispersão no gráfico de linha
	line.Overlap(scatter)

	return line
}

func main() {
	// Dados de entrada
	points := []Point{
		{0.19778, 1.99133},
		{0.394267, 0.649714},
		{0.512318, 0.000948577},
		{0.73689, 0.0615512},
		{0.896797, 0.237989},
		{0.988253, 0.271851},
		{1.16666, 0.353404},
		{1.57668, 4.94699},
		{1.74318, 4.96514},
		{1.65267, 5.25835},
	}

	// Calcular a tabela de diferenças divididas
	ddTable := dividedDifferences(points)

	// Exibir a tabela de diferenças divididas com o cabeçalho corrigido
	fmt.Println("Tabela de Diferenças Divididas:")
	fmt.Printf("%-10s %-10s", "x", "y")
	for j := 1; j < len(points); j++ {
		fmt.Printf(" DD%d        ", j)
	}
	fmt.Println()
	for i := 0; i < len(points); i++ {
		fmt.Printf("%-10.6f %-10.6f", points[i].x, points[i].y)
		for j := 0; j < len(ddTable[i]); j++ { // Ajuste para colunas
			fmt.Printf(" %-10.6f", ddTable[i][j])
		}
		fmt.Println()
	}

	// Valor de z para interpolação
	z := 0.726
	fmt.Println("\nEstimativas | f(z) no ponto z = 0.726")
	fmt.Println("k\tP_k(z)\t\tER_k")
	var mostReliableOrder int
	var minError float64 = math.MaxFloat64
	var bestApproximation float64

	for k := 0; k <= 9; k++ {
		approx := newtonInterpolation(points, ddTable, z, k)
		errorRelative := 0.0
		if k > 0 {
			errorRelative = math.Abs(approx-ddTable[0][k]) / math.Abs(ddTable[0][k])
		}
		fmt.Printf("%d\t%.6f\t%.6f\n", k, approx, errorRelative)

		if errorRelative < minError {
			minError = errorRelative
			mostReliableOrder = k
			bestApproximation = approx
		}
	}
	fmt.Printf("\nAproximação mais confiável: k = %d, P_k(z) = %.6f\n", mostReliableOrder, bestApproximation)

	// Gerar gráficos para P2(x), P4(x), P6(x) e P8(x)
	page := components.NewPage()
	page.AddCharts(
		createChart(points, ddTable, 2, "blue"),
		createChart(points, ddTable, 4, "red"),
		createChart(points, ddTable, 6, "orange"),
		createChart(points, ddTable, 8, "green"),
	)

	// Salvar o gráfico em um arquivo HTML
	f, err := os.Create("polynomial_interpolation.html")
	if err != nil {
		log.Fatal("Erro ao criar o arquivo HTML:", err)
	}
	defer f.Close()

	page.Render(f)
	fmt.Println("Gráficos gerados com sucesso! Abra o arquivo 'polynomial_interpolation.html' para visualizar.")
}

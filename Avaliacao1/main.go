package main

import (
	"avaliacao1/metodos"
	"fmt"
)

func main() {
	fmt.Printf("Iniciando métodos numéricos...\n")

	fmt.Println("MÉTODO DA BISSEÇÃO | DETERMINAÇÃO DA RAÍZ z1")
	metodos.Bisection(8.0, 10.0, 0.000001, 50)
	fmt.Println("--------------------------------------------------------------------------------")

	fmt.Println("MÉTODO DE NEWTON | DETERMINAÇÃO DA RAÍZ z2")
	metodos.Newton(11.0, 50, 1e-10)
	fmt.Println("--------------------------------------------------------------------------------")

	fmt.Println("MÉTODO DA SECANTE | DETERMINAÇÃO DA RAÍZ z3")
	metodos.Secant(12.0, 12.5, 0.0000001, 50)
	fmt.Println("--------------------------------------------------------------------------------")

	fmt.Println("MÉTODO DA FALSA POSIÇÃO | DETERMINAÇÃO DA RAÍZ z4")
	metodos.FalsePosition(12.0, 13.0, 0.0000001, 50)
	fmt.Println("--------------------------------------------------------------------------------")

	fmt.Println("MÉTODO DE HORNER | DETERMINAÇÃO DA RAÍZ z5")
	metodos.Horner(50, 1e-6)
	fmt.Println("--------------------------------------------------------------------------------")

}

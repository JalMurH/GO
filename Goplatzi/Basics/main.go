package main

import "fmt"

func main() {

	// Declaracion de constantes
	const pi float64 = 3.1416
	const pi2 = 3.1416
	fmt.Println(pi)
	fmt.Println(pi2)
	fmt.Println(pi, pi2)

	//Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	//zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	//print hello world
	fmt.Println("\nHello Golang")

	//Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del cuadrado:", areaCuadrado)

	//Operadores aritmeticos
	x := 10
	y := 50
	// Suma
	result := x + y
	fmt.Println("Suma:", result)
	// Resta
	result = x - y
	fmt.Println("Resta:", result)
	// Producto
	result = x * y
	fmt.Println("Producto:", result)
	// Division
	result = y / x
	fmt.Println("Div:", result)
	//Modulo o residuo
	result = y % x
	fmt.Println("Residuo:", result)
	// Incremental
	x++
	fmt.Println("Incremental:", x)
	// Decremental
	y--
	fmt.Println("Incremental:", y)

	// Reto Area de Rectangulo, trapecio, circulo
	// Recatangulo
	ARectangulo := x * y
	fmt.Printf("\nArea del rectangulo: %d m^2", ARectangulo)
	// Trapecio
	const h = 20
	ATrapecio := h * (y - x) / 2
	fmt.Printf("\nArea del trapecio: %d m^2", ATrapecio)
	// Circulo
	var r float64 = 5
	ACirculo := pi * (r * r)
	fmt.Printf("\nArea del circulo: %f m^2", ACirculo)

	//Complex64 complex128
	const w complex64 = 10 + 2i
	fmt.Println("\n", w)
}

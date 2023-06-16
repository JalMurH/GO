package main

import "fmt"

type figuras2D interface {
	area() float64
}
type cuadrado struct {
	base float64
}
type rectangulo struct {
	base float64
	alto float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}
func (c rectangulo) area() float64 {
	return c.base * c.alto
}

func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, alto: 4}

	calcular(myCuadrado)
	calcular(myRectangulo)

	//Lista de interfaces
	miInterface := []interface{}{"hola", 12, 4.8}
	fmt.Println(miInterface...)
}

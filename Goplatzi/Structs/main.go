package main

import "fmt"

//Structs son el equivamente de las class
type car struct {
	//atributos
	brand string
	model int
}

func main() {
	//forma de instancia de structura
	miCar := car{brand: "ford", model: 2020}
	fmt.Println(miCar)

	//como clase vacia
	var otherCar car
	otherCar.brand = "ferrari"
	otherCar.model = 2023

	fmt.Println(otherCar)
}

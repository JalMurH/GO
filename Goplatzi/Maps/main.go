package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//recorrido
	for i, v := range m {
		fmt.Println(i, v)
	}
	//Verificar si la llave a la que accedemos si se habia creado anteriormente
	value, ok := m["Jose"]

	fmt.Println(value, ok)
}

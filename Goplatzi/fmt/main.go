package main

import "fmt"

func printFunction(message string) {
	fmt.Println(message)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (C, d int) {
	return a, a * 2
}

func main() {
	printFunction("Holamundo")
	value := returnValue(2)
	fmt.Println(value)
	value1, value2 := doubleReturn(3)
	fmt.Println(value1, value2)
}

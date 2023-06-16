package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}
func main() {
	fmt.Println(sum(1, 2, 3))
}

package main

import "fmt"

//types come after variable name
func sum(x int, y int) int {
	return x + y
}

//consecutive variables can share type
func shortSum(x, y int) int {
	return x + y
}

//test
func evaluate(x, y int, z float32) float32 {
	return (float32)(x+y) * z
}

func main() {
	fmt.Println(sum(5, 2))
	fmt.Println(shortSum(5, 2))
	fmt.Println(evaluate(5, 2, 1))
}

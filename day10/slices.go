package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3} //array
	fmt.Println(a)

	b := []int{1, 2, 3} //slice
	fmt.Println(b)

	//slices are dynamically-sized while arrays are fixed-size
}

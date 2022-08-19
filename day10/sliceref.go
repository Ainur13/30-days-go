package main

import "fmt"

func addSlice(b []int) {
	b[0] += 1
	fmt.Println("B in func: ", b)
}

func addArray(a [3]int) {
	a[0] += 1
	fmt.Println("A in func:", a)
}

func main() {
	a := [3]int{1, 2, 3}
	//arrays contain just values
	b := []int{1, 2, 3}
	//slices contain references to values

	addSlice(b)
	fmt.Println("B: ", b)

	addArray(a)
	fmt.Println("A:", a)

}

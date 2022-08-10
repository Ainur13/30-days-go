package main

import "fmt"

//return values can be named
//however they harm readability in long functions
func split(n int) (x, y int, z float32) {
	x = n / 2
	y = n + 1

	// so return without arguments return named variables
	//this is known as "naked" return
	return
}

func main() {
	fmt.Println(split(5))
}

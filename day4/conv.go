package main

import "fmt"

func main() {
	var i int = 5

	//golang always requires explicit conversion
	var f float32 = float32(i)
	fmt.Println(f)
}

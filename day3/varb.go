package main

import "fmt"

//variable of package level
var d int

func main() {
	//variable of function level
	var a int
	fmt.Println(a)

	//long init
	var b, c int = 1, 2
	//mid init
	var f, g = 7, 8
	//short init
	//can be used only inside functions
	d, e := 3, 4

	fmt.Println(b, c, d, e, f, g)
}

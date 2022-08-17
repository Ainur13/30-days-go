package main

import "fmt"

func main() {
	//arrays cannot be resize
	//cause size is the part of their type
	var a [2]int
	a[0] = 1
	a[1] = 2

	fmt.Println(a)

	var b = [3]rune{'a', 'b', 'c'}
	fmt.Println(b)
}

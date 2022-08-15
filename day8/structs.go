package main

import "fmt"

//collection of fields
type Point struct {
	X int
	Y int
}

func main() {
	var p = Point{1, 2}
	fmt.Println(p)
}

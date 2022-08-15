package main

import "fmt"

func main() {
	//pointers holds the memory address of a value
	var p *int //it's value is nill
	i := 42
	p = &i

	fmt.Println(*p)
	fmt.Println(p)

	*p = 1
	fmt.Println(*p)
	fmt.Println(p)
}

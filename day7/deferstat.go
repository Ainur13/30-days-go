package main

import "fmt"

func main() {
	//defer executes tha last in function
	//function executes defer even it it panics
	//so we can think of it like a try - finally statement
	//multiple defer statements allowed and will be executed in reverse order (like in stack)
	defer fmt.Println("world")
	fmt.Println("hello")
}

package main

import "fmt"

//function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "ainur")
	fmt.Println(a, b) //to print elements seperated by space use comma
}

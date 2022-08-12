package main

import "fmt"

func main() {
	if true {
		fmt.Println("True!")
	}

	// if can start  with init statement  like for loop
	// declared varb are scoped only until the end of if statement
	if c := 7; c < 6 {
		fmt.Println("True!")
	} else {
		fmt.Println(c)
	}
}

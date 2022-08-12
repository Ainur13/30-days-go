package main

import "fmt"

func main() {
	//the only loop in golang
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i)
	}

	//init and post statements are optional
	c := 5
	for c < 100 {
		fmt.Printf("%v ", c)
		c += 10
	}

	//infinite loop
	// for{
	// }

}

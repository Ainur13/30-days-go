package main

import "fmt"

func main() {
	var a []int            //it's a s;ice with cap = 0
	var b = make([]int, 3) //it's a slice with cap = 3

	fmt.Println(a, b)

	a = append(a, 1)
	b = append(b, 1, 2)

	fmt.Println(a, b)

	for i, v := range b {
		fmt.Printf("index: %v  value: %v \n", i, v)
	}

	//we also can skip index or value like
	// for _, v := range b
	// for i, _ := range b or just for i := range b

}

package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2

	fmt.Println(m)

	m["two"] = 3
	fmt.Println(m)

	delete(m, "two")
	fmt.Println(m)

	elem, ok := m["two"]
	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println("deleted")
	}
}

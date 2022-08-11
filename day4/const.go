package main

import "fmt"

func main() {
	const c = "constant"
	const cc int = 4

	//can not be declared with :=
	//const ccc := false

	fmt.Println(c, cc)
}

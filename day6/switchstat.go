package main

import "fmt"

func main() {
	var a int = 5
	switch a {
	case 1:
		fmt.Println("It's 1")
		//break is automated in golang
		//that's why in golang switch statement runs only the selected case
	case 2:
		fmt.Println("It's 2")
	default:
		fmt.Println("Whatever")
	}

	switch ex := "mom"; ex {
	case "mom":
		fmt.Println("It's mom!")
	case "dad":
		fmt.Println("It's dad!")
	default:
		fmt.Println("Whatever")
	}
}

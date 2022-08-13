package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Monday?")
	today := time.Now().Weekday()
	//fmt.Println(today)
	//t := today + 1
	//fmt.Printf("%T\n", t)
	//why it's not working&
	switch time.Monday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tommorow")
	case today + 2:
		fmt.Println("After a day")
	case today + 3:
		fmt.Println("After a 2 days")
	default:
		fmt.Println("Too far away...")
	}
}

package main

import "fmt"

//by default all variables without initial value have zero value
// 0 for numbers
// false for bool
// empty for string

func main() {
	//basic types
	var b bool
	fmt.Printf("type: %T value: %v\n", b, b)

	var s string
	var r rune // is like char in c
	// alias for int32
	fmt.Printf("type: %T\n", s)
	fmt.Printf("type: %T value: %v\n", r, r)

	var bb byte // alias for uint8
	fmt.Printf("type: %T\n", bb)

	var i int //by default is int32 in 32bit system and int64 in 64bit system
	//int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64
	//uintptr - integer representation of memory address
	fmt.Printf("type: %T\n", i)

	var f float32 // also there're float64
	fmt.Printf("type: %T\n", f)

	var c complex64 // also there're comples128
	fmt.Printf("type: %T\n", c)
}

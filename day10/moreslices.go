package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	a = a[2:4]
	// [3, 4] but unerlying array is [3, 4, 5]
	fmt.Println(a, len(a), cap(a))

	a = a[1:]
	// [4] but unerlying array is [4, 5]
	fmt.Println(a, len(a), cap(a))

	//painc, cap is only 2
	//a = a[:3]
	//fmt.Println(a, len(a), cap(a))

	a = a[1:]
	//[] but underlying array is [5]
	fmt.Println(a, len(a), cap(a))

	//panic, cap is only 1, out of range
	//a = a[2:]
	//fmt.Println(a, len(a), cap(a))

	//len is a length of slice
	//cap is a capacity(length) of underlying array

	//underlying array is cutted from start
	//if slice is cutted from start
	//not end
}

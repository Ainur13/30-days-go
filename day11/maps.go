package main

import "fmt"

type Coord struct {
	X, Y float32
}

func main() {
	capitals := make(map[string]Coord)

	capitals["Astana"] = Coord{X: 1, Y: 2}
	capitals["London"] = Coord{X: 3, Y: 4}

	fmt.Println(capitals)

	capitals2 := map[string]Coord{
		"Astana": {1, 2},
		"London": {3, 4},
	}

	fmt.Println(capitals2)

	//nil map
	var capitals3 map[string]Coord
	//panic
	//we cannot add keys to nil map
	//capitals3["Astana"] = Coord{X: 1, Y: 2}
	fmt.Println(capitals3)

	var capitals4 map[string]Coord = map[string]Coord{
		"Astana": {1, 2},
		"London": {3, 4},
	}
	fmt.Println(capitals4)
}

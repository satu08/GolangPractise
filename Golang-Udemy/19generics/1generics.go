package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func addI(x int, y int) int {
	return x + y
}

func addF(x float64, y float64) float64 {
	return x + y
}

// set Interface
type myNumbers interface {
	//~int | ~float64 // ~tells to include all type of underlying integers
	constraints.Integer | constraints.Float // constraints package
}

// generic function - type constraint
func addT[T myNumbers](x, y T) T {
	return x + y
}

type myalias int

func main() {
	var n myalias = 42
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.3, 2.8))

	fmt.Println(addT(n, 2))
	fmt.Println(addT(1.9, 2.8))
}

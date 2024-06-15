package main

import (
	"fmt"
)

func intdelta(n *int) {
	*n = 42
}

func intdelta2(n *int) {
	*n = *n + 1
}

func intdelta1(n int) int {
	return n + 1
}
func slicedelta(xi []int) {
	xi[0] = 99
}

func mapdelta(md map[string]int, s string) {
	md[s] = 56
}

// everything in go is pass by value
func main() {
	a := 21

	fmt.Println(a)
	fmt.Println(intdelta1(a))
	fmt.Println(a)
	intdelta(&a)
	fmt.Println(a)
	intdelta2(&a)
	fmt.Println(a)

	slice1 := []int{1, 2, 4, 5, 5, 6}
	fmt.Println(slice1)
	slicedelta(slice1)
	fmt.Println(slice1)

	map1 := make(map[string]int)
	map1["satya"] = 24
	fmt.Println(map1)
	mapdelta(map1, "satya")
	fmt.Println(map1)

}

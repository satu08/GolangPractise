package main

import "fmt"

func main() {
	fmt.Println(withEscapeAnalysis())
	z := withoutEscapeAnalysis()
	fmt.Println(z)
}

func withEscapeAnalysis() *int {
	x := 43
	return &x // x escapes to the heap
}
func withoutEscapeAnalysis() int {
	y := 42
	return y
}

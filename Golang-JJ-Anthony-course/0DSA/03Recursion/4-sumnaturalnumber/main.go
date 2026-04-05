package main

import "fmt"

func main() {
	value := sumNaturalNumber(5)
	fmt.Println(value)
	value2 := mathFormulaSolution(5)
	fmt.Println(value2)
}

func sumNaturalNumber(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumNaturalNumber(n-1)
}

func mathFormulaSolution(n int) int {
	return (n * (n + 1)) / 2
}

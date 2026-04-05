package main

import "fmt"

func main() {
	factorial := factorial(5)
	fmt.Println(factorial)
}

func factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

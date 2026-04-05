package main

import "fmt"

func main() {
	value := nthFibonacci(3)
	fmt.Println(value)
}

func nthFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return nthFibonacci(n-1) + nthFibonacci(n-2)
}

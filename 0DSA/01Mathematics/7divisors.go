package main

import "fmt"

func main() {
	arr := divisors(9)
	fmt.Println(arr)
	fmt.Println("printing each divisor")
	for _, x := range arr {
		fmt.Println(x)
	}
}

func divisors(x int) []int {
	var divisors []int
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

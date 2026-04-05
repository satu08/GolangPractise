package main

import "fmt"

func main() {
	value := sumOfDigits(9987)
	fmt.Println(value)
}

func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 10) + sumOfDigits(n/10)
}

package main

import "fmt"

func main() {

	// time complexity is 3^n
	value := RopeCuttingProblem(9, 2, 2, 2)
	fmt.Println(value)
}

func RopeCuttingProblem(n, a, b, c int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	result := maximum(RopeCuttingProblem(n-a, a, b, c), RopeCuttingProblem(n-b, a, b, c), RopeCuttingProblem(n-c, a, b, c))
	if result == -1 {
		return -1
	}
	return result + 1
}

// maximum of three numbers
func maximum(number1, number2, number3 int) int {
	var largest int
	if number1 >= number2 && number1 >= number3 {
		largest = number1
	} else if number2 >= number1 && number2 >= number3 {
		largest = number2
	} else {
		largest = number3
	}
	return largest
}

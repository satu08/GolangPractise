package main

import "fmt"

func main() {
	factorial := factorial2(5)
	fmt.Println(factorial)
	fmt.Println(trailingZeroesInFactorialNaiveSolution(20))
	fmt.Println(TrailingZeroesEfficient(100))
}

func factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

func factorial2(x int) int {
	fact := 1

	for x > 0 {
		fact *= x
		x--
	}
	return fact
}

func trailingZeroesInFactorialNaiveSolution(x int) int {
	fact := factorial(x)

	trailingZeroes := 0
	for fact > 0 {
		if fact%10 == 0 {
			trailingZeroes++
		}
		fact /= 10
	}
	return trailingZeroes
}

func TrailingZeroesEfficient(x int) int {
	zeroes := 0
	for i := 5; i <= x; i *= 5 {
		zeroes = zeroes + x/i
	}
	return zeroes
}

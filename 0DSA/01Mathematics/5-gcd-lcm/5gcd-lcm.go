package main

import "fmt"

func main() {
	lcm := lcm(4, 6)
	fmt.Println(lcm)
	gcd := gcd(4, 6)
	fmt.Println(gcd)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	min := min(a, b)
	for min > 0 {
		if a%min == 0 && b%min == 0 {
			return min
		}
		min--
	}
	return 1
}

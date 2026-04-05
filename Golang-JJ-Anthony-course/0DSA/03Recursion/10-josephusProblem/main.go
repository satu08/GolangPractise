package main

import "fmt"

func main() {
	value := josephus(3, 3)
	fmt.Println(value)
}

func josephus(n, k int) int {
	if n == 1 {
		return 0
	}
	return ((josephus(n-1, k) + k) % n)
}

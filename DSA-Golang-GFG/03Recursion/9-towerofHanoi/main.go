package main

import "fmt"

var count int

func main() {
	toh(3, "a", "b", "c")
	fmt.Println(count)
}

// number of movements (2^n - 1)
func toh(n int, a, b, c string) {
	count++
	if n == 1 {
		fmt.Println("move 1 from", a, "to", c)
		return
	}
	toh(n-1, a, c, b)
	fmt.Println("move", n, "from", a, "to", c)
	toh(n-1, b, a, c)
}

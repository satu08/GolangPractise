package main

import "fmt"

func main() {
	value := logbase2n(16)
	fmt.Println(value)
	binaryofnumber(8)
}

// return log base 2 of n
func logbase2n(n int) int {
	if n == 1 {
		return 0
	} else {
		return 1 + logbase2n(n/2)
	}
}

func binaryofnumber(n int) {
	if n == 0 {
		return
	}
	binaryofnumber(n / 2)
	fmt.Printf("%v", n%2)
}

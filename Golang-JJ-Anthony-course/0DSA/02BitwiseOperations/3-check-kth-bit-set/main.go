package main

import "fmt"

// we have two numbers x and y  if we convert x into binary then we need to check if yth bit of x is 1 or not from right

func main() {

	x := 5
	y := 3

	kthBitSetUsingLeftShift(x, y)
	kthBitSetUsingRightShift(x, y)
}

func kthBitSetUsingLeftShift(x, y int) {
	if (x & (1 << (y - 1))) != 0 {
		fmt.Println("k-th bit is set")
	} else {
		fmt.Println("k-th bit is not set")
	}
}

func kthBitSetUsingRightShift(x, y int) {
	if (1 & (x >> (y - 1))) != 0 {
		fmt.Println("k-th bit is set")
	} else {
		fmt.Println("k-th bit is not set")
	}
}

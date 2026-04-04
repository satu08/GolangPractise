package main

import "fmt"

func main() {
	power := pow(4, 2)
	fmt.Println(power)
}

func pow(x int, y int) int {
	var pow int = 1
	for i := 0; i < y; i++ {
		pow *= x
	}
	return pow
}

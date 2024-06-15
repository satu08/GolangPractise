package main

import "fmt"

func main() {
	fmt.Println(add(3, 4))
}
func add(num ...int) int {
	sum := 0
	for _, num := range num {
		sum += num
	}
	return sum
}

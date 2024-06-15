package main

import (
	"fmt"
)

func main() {
	// normal for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five:%v \t first for loop\n", i)
	}
	// do while loop
	var y int
	for y < 10 {
		fmt.Printf("y is %v \t second for loop\n", y)
		y++
	}
	// break statement
	for {
		fmt.Printf("y is %v \t third for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}
	// continue statements
	for i := 0; i < 20; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println("counting even numbers", i)
	}

	// nested loops

	for i := 0; i < 5; i++ {
		fmt.Println("---")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
	}
}

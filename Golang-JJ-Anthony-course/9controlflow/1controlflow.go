package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("this is first statement")
	fmt.Println("this i second statement")
	var x, y int
	fmt.Println("Enter a number x and y")
	_, err := fmt.Scan(&x, &y)
	if err != nil {
		fmt.Println("error occurred while taking input", err)
	}

	fmt.Printf("x=%v and y=%v", x, y)
	// if,else if,else
	if x < 42 {
		fmt.Println("less than 42")
	}

	if x < 42 {
		fmt.Println("less than 42")
	} else {
		fmt.Println("greater than 42")
	}

	if x < 42 {
		fmt.Println("less than 42")
	} else if x == 42 {
		fmt.Println("equal to 42")
	} else {
		fmt.Println("greater than 42")
	}

	// logical operators

	if x < 42 || y > 10 {
		fmt.Println("at least second variable met the condition")
	}

	if x < 42 && y > 10 {
		fmt.Println("both conditions are false")
	}

	if x != 42 || y != 10 {
		fmt.Println("at least one condition is true")
	}

	// statement statement idiom

	if z := x * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is grater than or equal to x=%v\n", z, x)
	} else {
		fmt.Printf("z=%v is less than x=%v\n", z, x)
	}

	//switch case

	switch {
	case x < 42:
		fmt.Println("x is less than 42")
		fallthrough // workout next case as well
	case x == 42:
		fmt.Println("X is equal to 42")
	case x > 42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("this is default case ")
	}
}

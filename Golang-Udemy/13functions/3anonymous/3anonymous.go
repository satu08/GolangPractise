package main

import "fmt"

func main() {
	rs()

	x := func() {
		fmt.Println("an anonymous function")
	}
	x() // func expression

	y := func(s string) {
		fmt.Println("my name is ", s)
	}
	y("satyanarayan") // func expression

	z := returnfunc()
	fmt.Println(z())
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", returnfunc())
	fmt.Printf("%T\n", z)
}
func rs() {
	fmt.Println("satya")
}

// returning a function
func returnfunc() func() int {
	return func() int {
		return 43
	}
}

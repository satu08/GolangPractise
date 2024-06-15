package main

import "fmt"

func main() {

	fmt.Println(doMath(12, 8, addition))
	fmt.Println(doMath(12, 8, subtraction))

	f := incrementor()
	fmt.Println(f())
	f = incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

// callback
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func addition(a int, b int) int {
	return a + b
}
func subtraction(a int, b int) int {
	return a - b
}

// closure
func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
	fmt.Println(variadicfunc(1, 2, 3, 4, 5, 6))
	fmt.Println(funcwithslice([]int{1, 2, 3, 4, 5, 6}))
	// defer functions are LIFO
	defer fmt.Println("satya1")
	defer fmt.Println("satya2")
	defer fmt.Println("satya3")
}
func foo() int {
	return 42
}

func bar() (string, int) {
	return "satya", 23
}

func variadicfunc(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func funcwithslice(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

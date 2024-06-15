package main

import "fmt"

func main() {
	fmt.Println("introduction to functions")
	defer foo() // executes at the last - waits for other functions to executes
	bar("satyanarayan")
	name := foofoo("satya")
	fmt.Println(name)
	name1, age1 := barbar("satya", 7)
	fmt.Println(name1, age1)
	// passing a values from a slice to a function
	// unfurling a slice in a variadic function
	xi := []int{1, 3, 4, 5, 6, 6, 76}
	total := sum(xi...)
	fmt.Println("total is", total)
}

// no params no return
func foo() {
	fmt.Println("i am from foo")
}

// One param, no return
func bar(s string) {
	fmt.Println("my name is ", s)
}

// one param, one return

func foofoo(s string) string {
	return fmt.Sprint("they call me Mr.", s)
}

// two param, two returns

func barbar(s string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(s, " is elder than dogs age by "), age
}

// variadic parameters
// variadic parameter should be final incoming parameter to a func

func sum(numbers ...int) int {
	fmt.Printf("%T\n", numbers)
	fmt.Printf("%v\n", numbers)

	sum1 := 0
	for _, v := range numbers {
		sum1 += v
	}
	return sum1
}

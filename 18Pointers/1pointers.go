package main

import "fmt"

func main() {
	x := 21
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%v \t %T \n", &x, &x)

	//dereferencing the pointers

	y := &x
	fmt.Printf("%v \t %T \n", y, y)
	fmt.Println(*y)
	fmt.Println(*&x)

	*y = 43
	fmt.Println(x)
	fmt.Println(*y)
	fmt.Println(&y)

}

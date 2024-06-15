package main

import "fmt"

func main() {
	a := [3]int{21, 21, 323}
	fmt.Println(a)
	fmt.Println("length of array ", len(a))

	b := [...]string{"satya", "jadhav"}
	fmt.Println(b)

	var c [2]int
	c[0] = 3
	c[1] = 4
	fmt.Println(c)
	fmt.Printf("%T\n", a)

}

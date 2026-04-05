package main

import "fmt"

func main() {
	var a int = 10
	fmt.Println(a)

	b := 20 // short declaration operator
	fmt.Println(b)

	//c, d, e, f := 30, 40, 50, 60
	//fmt.Println(c, d, e) //this will not work as we have declared the f but not used it

	g, h, _, i := 30, 40, 50, 60
	fmt.Println(g, h, i) // this will work as we are not declaring the variable

}

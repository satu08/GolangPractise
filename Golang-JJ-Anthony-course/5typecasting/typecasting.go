package main

import "fmt"

func main() {
	a := 20
	b := 30.67
	fmt.Printf("%v of type %T \n", a, a)
	fmt.Printf("%v of type %T \n", b, b)

	var m float32 = 43.67
	fmt.Printf("%v of type %T\n", m, m)
	b = float64(m)
	fmt.Printf("%v of type %T \n", b, b)
	k := int64(m)
	fmt.Printf("%v of type %T", k, k)
}

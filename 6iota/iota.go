package main

import "fmt"

type catt int

const (
	c0 catt = iota
	c1
	c2
)

func main() {
	fmt.Println(c0, c1, c2)
}

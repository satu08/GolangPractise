package main

import "fmt"

func addOne(v int) int {
	return v + 1
}
func addOneP(v *int) {
	*v += 1
}

type person struct {
	first string
}

func main() {
	//value semantics

	a := 1
	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Println(a)

	// pointer semantic
	b := 1
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)

	p1 := person{
		first: "satya",
	}
	fmt.Println(p1)
	p1 = changeName(p1, "ravi")
	fmt.Println(p1)
	changeNamePtr(&p1, "swati")
	fmt.Println(p1)
}
func changeName(p person, s string) person {
	p.first = s
	return p
}

func changeNamePtr(p *person, s string) {
	p.first = s
}

package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	name string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("i am", p.name)
}

func (sa secretAgent) speak() {
	fmt.Println("i am a secret agent ", sa.name)
}

func saySomething(human2 human) {
	human2.speak()
}

func main() {
	p1 := secretAgent{
		person: person{
			name: "satya",
		},
		ltk: true,
	}
	p2 := person{
		name: "ravi",
	}
	saySomething(p1)
	saySomething(p2)
}

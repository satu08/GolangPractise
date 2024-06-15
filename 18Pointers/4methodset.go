package main

import "fmt"

type dog struct {
	firstname string
}

func (d dog) walk() {
	fmt.Println("hi my name is ", d.firstname, " and i am walking")
}

func (d *dog) run() {
	fmt.Println("hi my name is ", d.firstname, " and i am running")
}

type young interface {
	walk()
	run()
}

func youngin(y young) {
	y.run()
}

func main() {
	d1 := dog{
		"tommy",
	}

	d2 := &dog{
		"padget",
	}
	d1.walk()
	d1.run()
	//youngin(d1) this is not going to work because it is a value, so it can use value receiver methods to implement interfaces

	d2.walk()
	d2.run()
	youngin(d2)

}

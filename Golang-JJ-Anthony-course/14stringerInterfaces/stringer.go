package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

// attaching a function as a method to a type using receiver
func (b book) String() string {
	return fmt.Sprint("the title of book is --> ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("the number is -->", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("logger:-", s.String())
}
func main() {
	b2 := book{
		title: "we met in a wrong universe",
	}
	var n count = 32

	logInfo(b2)
	logInfo(n)
}

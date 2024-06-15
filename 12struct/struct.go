package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
	favIc []string
}

// embedded structs
type humanbeing struct {
	person
	knowledge bool
}

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	fmt.Println("introduction to structs")
	p1 := humanbeing{
		person: person{
			first: "satya",
			last:  "jadhav",
			age:   23,
			favIc: []string{"chocolate", "vanilla", "buttersctoch", "strawberry"},
		},
		knowledge: true,
	}
	// anonymous struct
	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "vedant",
		last:  "deshpande",
		age:   23,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Printf("%T\n", p2)
	fmt.Println(p1.first, p1.age, p1.person)
	for _, v := range p1.person.favIc {
		fmt.Println("p1 favorite icecream is", v)
	}

	// making map of struct
	fmt.Println("printing map")
	map1 := map[string]humanbeing{
		p1.last: p1,
	}

	for _, v := range map1 {
		fmt.Println(v)
		for _, v2 := range v.favIc {
			fmt.Println(v.first, v.last, v2)
		}
	}

	// embed struct

	veh1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "yes",
		model: " fortuner",
		doors: 4,
		color: "white",
	}

	veh2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "no",
		model: " endevour",
		doors: 5,
		color: "black",
	}

	fmt.Println(veh2)
	fmt.Println(veh1)

	fmt.Println(veh1.color)
	fmt.Println(veh2.color)

}

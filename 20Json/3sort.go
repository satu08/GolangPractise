package main

import (
	"fmt"
	"sort"
)

type people struct {
	name string
	age  int
}

func (p people) String() string {
	return fmt.Sprintf("%s: %d", p.name, p.age)
}

func main() {
	s := []int{90, 87, 0, 3, 554, 729, 23}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)

	ss := []string{"satya", "a", "dr.shraddha", "#$%", "i am a enginner"}
	fmt.Println(ss)
	sort.Strings(ss)
	fmt.Println(ss)

	p1 := people{
		name: "satya",
		age:  23,
	}
	p2 := people{
		name: "shraddha",
		age:  26,
	}
	p3 := people{
		name: "vedant",
		age:  24,
	}
	p4 := people{
		name: "ravi",
		age:  27,
	}
	humanbeings := []people{p1, p2, p3, p4}
	fmt.Println(humanbeings)
}

package main

import "fmt"

func main() {
	// slice data structure
	xi := []int{21, 31, 31, 32, 24, 324, 31}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}
	// range over map data structure
	m := map[string]string{
		"name":    "Satyanarayan",
		"company": "Siemens",
	}
	for k, v := range m {
		fmt.Println("ranging over map", k, v)
	}

	nameOfStudent := m["name"]
	fmt.Println("name of student is :", nameOfStudent)

	if n, ok := m["name"]; ok {
		fmt.Printf("name of student is %v\n", n)
	}
}

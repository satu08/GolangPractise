package main

import "fmt"

func main() {
	slice1 := []int{1, 3, 4, 5, 42}
	fmt.Println(slice1)
	fmt.Println(len(slice1))

	for key, value := range slice1 {
		fmt.Printf("key= %v , value =%v\n", key, value)
	}
	// blank identifier to not use a returned value
	for _, value := range slice1 {
		fmt.Printf("value =%v\n", value)
	}
	//printing by taking indexes
	for key, _ := range slice1 {
		fmt.Printf("value =%v\n", slice1[key])
	}

	// appending to a slice
	slice1 = append(slice1, 45, 32, 21, 54)
	fmt.Println(slice1)
	fmt.Println("--------")

	// slicing of slices

	//1 [inclusive:exclusive]
	fmt.Printf("%#v\n", slice1[1:5])
	fmt.Println("--------")

	//2 [:exclusive]
	fmt.Printf("%#v\n", slice1[:4])
	fmt.Println("--------")

	//3 [inclusive:]
	fmt.Printf("%#v\n", slice1[3:])
	fmt.Println("--------")
	// 4 [:]
	fmt.Printf("%#v\n", slice1[:])

	// deleting from slices
	slice1 = append(slice1[:4], slice1[5:]...) // appending one slice to another
	fmt.Printf("%#v\n", slice1[:])

	// make slice
	slice2 := make([]int, 0, 10)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	slice2 = append(slice2, 12, 23, 343, 4, 3, 6, 1, 9)
	fmt.Println(slice2)
	fmt.Println("-------")
	slice2 = append(slice2, 23, 4, 54, 6, 9, 0)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// multidimensional slices
	fmt.Println("multidimensional slice")
	slice3 := [][]int{slice1, slice2}
	fmt.Println(slice3)
	fmt.Println("printing each value of multidimensional array")
	for i, v := range slice3 {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}

	// underlying array - slice internals
	a := []int{1, 3, 4, 5, 42}

	b := make([]int, 5)
	copy(b, a)
	fmt.Println("a", a)
	fmt.Println("b", b)
	a[0] = 7
	fmt.Println("a", a)
	fmt.Println("b", b) // 0th position of b remains as 1 because we are assigning a new space with make

}

package main

import "fmt"

func main() {
	value := Search([]int{1, 5, 4, 233, 6}, 233)
	fmt.Println(value)
	arr := insert([]int{1, 3, 4, 9, 10}, 3, 7)
	fmt.Println(arr)
	arr1 := delete([]int{1, 3, 4, 9, 10}, 4)
	fmt.Println(arr1)
}

/* array elements stored at contiguous memory
Random access
 cache friendliness */

// time complexity O(n)
func Search(arr []int, key int) int {
	length := len(arr)
	for i := 0; i < length; i++ {
		if arr[i] == key {
			return i

		}
	}
	return -1
}

// time complexity O(n)
func insert(arr []int, pos int, value int) []int {
	arr = append(arr[:pos-1], append([]int{value}, arr[pos-1:]...)...)
	return arr
}

// time complexity O(n)
func delete(arr []int, value int) []int {
	for i, v := range arr {
		if v == value {
			// Remove the element by its index
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

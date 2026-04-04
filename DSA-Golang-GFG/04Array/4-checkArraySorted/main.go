package main

import "fmt"

func main() {
	value := isArraySorted([]int{1, 4, 10, 7, 9})
	fmt.Println(value)
}

// time complexity is O(n)
func isArraySorted(array []int) bool {
	l := len(array)
	for i := l - 1; i > 0; i-- {
		if array[i] < array[i-1] {
			return false
		}
	}
	return true
}

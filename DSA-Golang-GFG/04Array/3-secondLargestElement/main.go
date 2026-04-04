package main

import "fmt"

func main() {
	value := secondLargestElement([]int{8, 78, 35, 82, 32, 90})
	fmt.Println(value)
}

// time complexity is Θ(n)
// returns the index of second largest
func secondLargestElementIndex(arr []int) int {
	largest := 0
	secondLargest := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[largest] {
			secondLargest = largest
			largest = i
		}
	}
	return secondLargest
}

// returns the value of second largest
func secondLargestElement(arr []int) int {
	largest := 0
	secondLargest := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			secondLargest = largest
			largest = arr[i]
		}
	}
	return secondLargest
}

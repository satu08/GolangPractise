package main

import "fmt"

func main() {
	index := largestElement([]int{1, 89, 78, 43, 10})
	fmt.Println(index)
}

func largestElement(nums []int) int {
	largest := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] > nums[largest] {
			largest = i
		}
	}
	return largest
}

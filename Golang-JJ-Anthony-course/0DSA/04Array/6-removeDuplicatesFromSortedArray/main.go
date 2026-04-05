package main

import "fmt"

func removeDuplicates(nums []int) ([]int, int) {
	length := len(nums)
	res := 1
	for i := 1; i < length; i++ {
		if nums[i] != nums[res-1] {
			nums[res] = nums[i]
			res++

		}
	}
	return nums[:res], res
}

func main() {
	arr, size := removeDuplicates([]int{1, 1, 1})
	fmt.Println(arr)
	fmt.Println(size)
}

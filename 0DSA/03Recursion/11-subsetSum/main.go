package main

import "fmt"

func main() {
	//arr := []int{20, 10, 15}
	value := subsetSum([]int{20, 10, 15}, 3, 25)
	fmt.Println(value)
}

func subsetSum(nums []int, n, sum int) int {
	if n == 0 {
		if sum == 0 {
			return 1
		} else {
			return 0
		}
	}
	return subsetSum(nums, n-1, sum) + subsetSum(nums, n-1, sum-nums[n-1])
}

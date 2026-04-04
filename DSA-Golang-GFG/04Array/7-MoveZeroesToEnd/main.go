package main

import "fmt"

func main() {
	zerosArray := moveZeroes([]int{10, 20, 0})
	fmt.Println(zerosArray)
}

func moveZeroes(nums []int) []int {
	res := len(nums) - 1
	for i := res; i >= 0; i-- {
		if nums[i] == 0 {
			temp := nums[res]
			nums[res] = 0
			nums[i] = temp
			res--
		}
	}
	return nums
}

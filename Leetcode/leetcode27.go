package main

import "fmt"

//func main() {
//	fmt.Println(removeElement([]int{2, 3, 3, 2}, 3))
//}

func removeElement(nums []int, val int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	fmt.Println(nums)
	return count
}

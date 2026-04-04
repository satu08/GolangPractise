package main

import "fmt"

func main() {
	fmt.Println(minConsFlips([]int{0, 0, 1, 1, 0, 0, 1, 1, 0}))
}

func minConsFlips(nums []int) int {
	count0 := 0
	count1 := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i-1] == 0 {
				count0++
			} else {
				count1++
			}
		}
	}
	if nums[len(nums)-1] == 0 {
		count0++
	} else {
		count1++
	}
	if count0 < count1 {
		return count0
	}
	return count1

}

package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 0, 1, 1, 1, 1, 0, 0, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	var currentCount int
	for _, num := range nums {
		if num != 1 {
			currentCount = 0
		} else {
			currentCount++
			maxCount = max(maxCount, currentCount)
		}
	}
	return maxCount
}

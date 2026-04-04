package main

import "fmt"

func main() {
	fmt.Println(trap([]int{5, 0, 6, 2, 3}))
}

func trap(height []int) int {
	n := len(height)

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	var result int
	for i := 0; i < n; i++ {
		result += min(leftMax[i], rightMax[i]) - height[i]
	}
	return result
}

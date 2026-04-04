package main

import "fmt"

func main() {
	fmt.Println(maxDifference([]int{30, 10, 8, 2}))
}

func maxDifference(arr []int) int {
	n := len(arr)
	minValue := arr[0]
	maxDiff := arr[1] - arr[0]

	for i := 1; i < n; i++ {
		if arr[i]-minValue > maxDiff {
			maxDiff = arr[i] - minValue
		}
		if arr[i] < minValue {
			minValue = arr[i]
		}
	}
	return maxDiff
}

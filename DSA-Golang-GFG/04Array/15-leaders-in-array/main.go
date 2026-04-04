package main

import "fmt"

func main() {
	fmt.Println(leadersInArray([]int{7, 10, 4, 3, 6, 5, 2}))
}

func leadersInArray(arr []int) []int {
	n := len(arr)
	var leaders []int
	maxFromRight := arr[n-1]
	leaders = append(leaders, maxFromRight)

	for i := n - 1; i >= 0; i-- {
		if arr[i] > maxFromRight {
			maxFromRight = arr[i]
			leaders = append(leaders, maxFromRight)
		}
	}
	return leaders
}

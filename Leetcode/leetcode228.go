package main

import "fmt"

//	func main() {
//		fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
//	}
func summaryRanges(nums []int) []string {
	var resultarray []string
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 != nums[i+1] {
			fmt.Println(i)
			resultarray = append(resultarray, fmt.Sprintf("%v->%v", nums[count], nums[i]))
			count = i
		}
	}
	return resultarray
}

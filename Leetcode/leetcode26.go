package main

// import "fmt"
//
//	func main() {
//		fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 2}))
//	}
func removeDuplicates(nums []int) int {
	length := len(nums)
	res := 1
	for i := 1; i < length; i++ {
		if nums[i] != nums[res-1] {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}

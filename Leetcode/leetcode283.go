package main

//	func main() {
//		fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
//	}
func moveZeroes(nums []int) []int {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			temp := nums[i]
			nums[i] = 0
			nums[l] = temp
			l++
		}
	}
	return nums
}

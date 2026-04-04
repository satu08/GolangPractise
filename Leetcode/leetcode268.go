package main

//	func main() {
//		fmt.Println(missingNumber([]int{3, 0, 1}))
//	}
func missingNumber(nums []int) int {
	sum := 0
	length := len(nums)
	actualSum := length * (length + 1) / 2

	for _, v := range nums {
		sum += v
	}
	return actualSum - sum
}

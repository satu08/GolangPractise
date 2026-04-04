package main

//	func main() {
//		fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
//	}
func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	missingMap := make(map[int]int)
	for i := 1; i <= length; i++ {
		missingMap[i] = 0
	}
	for i := 0; i < length; i++ {
		delete(missingMap, nums[i])
	}
	var missingArray []int

	for k, _ := range missingMap {
		missingArray = append(missingArray, k)
	}
	return missingArray
}

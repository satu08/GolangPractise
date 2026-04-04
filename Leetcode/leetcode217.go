package main

//	func main() {
//		fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
//	}
func containsDuplicate(nums []int) bool {
	duplicateMap := make(map[int]int)

	for _, v := range nums {
		if duplicateMap[v] > 0 {
			return true
		}
		duplicateMap[v]++
	}

	//for _, v := range duplicateMap {
	//	if v > 1 {
	//		return true
	//	}
	//}
	return false
}

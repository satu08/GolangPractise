package main

//	func main() {
//		fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
//	}
func majorityElement(nums []int) int {
	singleNumMap := make(map[int]int)
	for _, v := range nums {
		singleNumMap[v]++
	}
	var majority, count int
	for k, v := range singleNumMap {
		if v > count {
			majority = k
			count = v
		}
	}
	return majority
}

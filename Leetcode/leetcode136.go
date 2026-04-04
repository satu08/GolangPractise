package main

//	func main() {
//		fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
//		fmt.Println(singleNumber1([]int{4, 1, 2, 1, 2}))
//	}
func singleNumber(nums []int) int {
	singleNumMap := make(map[int]int)
	for _, v := range nums {
		singleNumMap[v]++
	}
	for k, _ := range singleNumMap {
		if singleNumMap[k] == 1 {
			return k
		}
	}
	return nums[0]
}

func singleNumber1(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}

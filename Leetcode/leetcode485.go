package main

//	func main() {
//		fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
//	}
func findMaxConsecutiveOnes(nums []int) int {
	maxConsOnes := 0
	ones := 0
	for _, v := range nums {
		if v == 1 {
			ones = ones + 1
			if ones > maxConsOnes {
				maxConsOnes = ones
			}
		} else if v == 0 {
			ones = 0
		}
	}
	return maxConsOnes
}

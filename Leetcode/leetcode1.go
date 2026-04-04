package main

// import "fmt"
//
//	func main() {
//		fmt.Println(twoSum([]int{11, 2, 7, 15}, 9))
//	}
func twoSum(nums []int, target int) []int {
	var output1 []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return append(output1, i, j)
				break
			}
		}
	}
	valuemap := make(map[int]int)
	var output []int
	for k, v := range nums {
		subs := target - v
		if _, ok := valuemap[subs]; !ok {
			valuemap[v] = k
		} else {
			return append(output, valuemap[subs], k)
		}
	}

	return output
}

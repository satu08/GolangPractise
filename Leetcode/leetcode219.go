package main

//	func main() {
//		fmt.Println(containsDuplicateII([]int{1, 0, 1, 1}, 1))
//	}
func containsDuplicateII(nums []int, k int) bool {
	duplicateMap := make(map[int][]int)

	for i, v := range nums {
		duplicateMap[v] = append(duplicateMap[v], i)
	}
	for _, v := range duplicateMap {
		for i := 0; i < len(v)-1; i++ {
			if absInt(v[i]-v[i+1]) <= k {
				return true
			}
		}
	}
	return false
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

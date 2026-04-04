package main

import "fmt"

//func main() {
//	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
//}

func findRelativeRanks(score []int) []string {

	n := len(score)
	sorted := make([]int, n)
	copy(sorted, score)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sorted[j] < sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	rankMap := make(map[int]string)
	for i, val := range sorted {
		switch i {
		case 0:
			rankMap[val] = "Gold Medal"
		case 1:
			rankMap[val] = "Silver Medal"
		case 2:
			rankMap[val] = "Bronze Medal"
		default:
			rankMap[val] = fmt.Sprintf("%d", i+1)
		}
	}
	result := make([]string, n)
	for i, val := range score {
		result[i] = rankMap[val]
	}
	return result
}

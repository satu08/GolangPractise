package main

import "fmt"

func main() {

	fmt.Println("Subsets:", subsets([]int{1, 2, 3}))

	fmt.Println("Permutations:", permute([]int{1, 2, 3}))

	fmt.Println("Combination Sum:", combinationSum([]int{2, 3, 6, 7}, 7))
}

// ===== Subsets =====
func subsets(nums []int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(int)
	backtrack = func(start int) {
		res = append(res, append([]int{}, path...))

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}

// ===== Permutations =====
func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])

			backtrack()

			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}

// ===== Combination Sum =====
func combinationSum(nums []int, target int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(int, int)
	backtrack = func(start int, sum int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i, sum+nums[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)
	return res
}

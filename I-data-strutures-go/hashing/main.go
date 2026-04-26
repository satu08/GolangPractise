package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 2, 1, 4}

	// 🔹 Frequency
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}
	fmt.Println("Frequency:", freq)

	// 🔹 Two Sum
	fmt.Println("Two Sum:", twoSum(arr, 5))

	// 🔹 Duplicate
	fmt.Println("Contains Duplicate:", containsDuplicate(arr))

	// 🔹 Subarray Sum
	fmt.Println("Subarray Sum =", subarraySum(arr, 3))
}

// helper functions
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		if idx, ok := mp[target-v]; ok {
			return []int{idx, i}
		}
		mp[v] = i
	}
	return nil
}

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, v := range nums {
		if set[v] {
			return true
		}
		set[v] = true
	}
	return false
}

func subarraySum(nums []int, k int) int {
	mp := map[int]int{0: 1}
	sum, count := 0, 0

	for _, v := range nums {
		sum += v
		if c, ok := mp[sum-k]; ok {
			count += c
		}
		mp[sum]++
	}
	return count
}

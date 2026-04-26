package main

import "fmt"

// 🔹 Linear Search (O(n))
func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// 🔹 Binary Search - Iterative (O(log n))
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 🔹 Binary Search - Recursive
func binarySearchRecursive(arr []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearchRecursive(arr, mid+1, right, target)
	} else {
		return binarySearchRecursive(arr, left, mid-1, target)
	}
}

func main() {

	// Unsorted array (for linear search)
	arr1 := []int{10, 5, 3, 8, 2}

	// Sorted array (for binary search)
	arr2 := []int{2, 3, 5, 8, 10}

	target := 8

	fmt.Println("🔹 Linear Search:")
	index1 := linearSearch(arr1, target)
	fmt.Println("Element found at index:", index1)

	fmt.Println("\n🔹 Binary Search (Iterative):")
	index2 := binarySearch(arr2, target)
	fmt.Println("Element found at index:", index2)

	fmt.Println("\n🔹 Binary Search (Recursive):")
	index3 := binarySearchRecursive(arr2, 0, len(arr2)-1, target)
	fmt.Println("Element found at index:", index3)
}

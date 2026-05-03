package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("sorting algorithms")
	fmt.Println("bubble sort")
	fmt.Println(bubbleSort([]int{1, 9, 5, 6, 3, 6}))
	fmt.Println("selection sort")
	fmt.Println(selectionSort([]int{1, 9, 5, 6, 3, 6}))
	fmt.Println("insertion sort")
	fmt.Println(insertionSort([]int{1, 9, 5, 6, 3, 6}))
	fmt.Println("merge sort")
	fmt.Println(mergeSort([]int{1, 9, 5, 6, 3, 6}))
	fmt.Println("intersection of two sorted arrays")
	fmt.Println(intersectionSorted([]int{1, 2, 3, 3, 4, 5}, []int{3, 3, 4, 5, 6, 7}))
	fmt.Println("intersection of two unsorted arrays")
	fmt.Println(intersection([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 6, 7}))
	fmt.Println("union of two sorted arrays")
	fmt.Println(unionSorted([]int{1, 2, 3, 3, 4, 5}, []int{3, 3, 4, 5, 6, 7}))

	fmt.Println("partitioned array is -->", partitionNaive([]int{3, 8, 6, 12, 10, 7}, 5))

	fmt.Println(maximumProduct([]int{-100, -98, -1, 2, 3, 4}))

	fmt.Println(findErrorNums([]int{1, 2, 4, 4}))
	fmt.Println(dominantIndex([]int{1, 0}))
	fmt.Println(checkIfExist([]int{3, 1, 7, 11}))
	fmt.Println(findKthPositive([]int{1, 3, 4}, 2))
}

// 🔹 Bubble Sort theta(n*n)  stable algorithm
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 🔹 Selection Sort  theta(n*n) find out minimum element in each loop and put it in first places -not stable (does not gurantee ordering  of equal elements)
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

// 🔹 Insertion Sort  O(n*n) stable algo
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

// 🔹 Merge Sort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func intersection(arr1, arr2 []int) []int {
	result := []int{}
	for _, v1 := range arr1 {
		for _, v2 := range arr2 {
			if v1 == v2 {
				result = append(result, v1)
				break
			}
		}
	}
	return result
}

func intersectionSorted(arr1, arr2 []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if i > 0 && arr1[i] == arr1[i-1] {
			i++
			continue
		}
		if arr1[i] < arr2[j] {
			i++
		} else if arr1[i] > arr2[j] {
			j++
		} else {
			result = append(result, arr1[i])
			i++
			j++
		}
	}
	return result
}

func unionSorted(arr1, arr2 []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if i > 0 && arr1[i] == arr1[i-1] {
			i++
			continue
		}
		if j > 0 && arr2[j] == arr2[j-1] {
			j++
			continue
		}
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else if arr1[i] > arr2[j] {
			result = append(result, arr2[j])
			j++
		} else {
			result = append(result, arr1[i])
			i++
			j++
		}
	}
	for i < len(arr1) {
		if i == 0 || arr1[i] != arr1[i-1] {
			result = append(result, arr1[i])
		}
		i++
	}
	for j < len(arr2) {
		if j == 0 || arr2[j] != arr2[j-1] {
			result = append(result, arr2[j])
		}
		j++
	}
	return result
}

func partitionNaive(arr []int, pos int) []int {
	i, j := 0, len(arr)-1
	result := make([]int, len(arr))
	for k := 0; k < len(arr); k++ {
		if arr[k] < arr[pos] {
			result[i] = arr[k]
			i++
		} else if arr[k] > arr[pos] {
			result[j] = arr[k]
			j--
		} else {
			continue
		}
	}
	result[i] = arr[pos]
	return result
}

func lomutoPartition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// place pivot in correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// 🔹 Quick Sort
func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func partitionTemplate(arr []int, condition func(int) bool) {
	i := -1
	for j := 0; j < len(arr); j++ {
		if condition(arr[j]) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
}

func sort012(arr []int) {
	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else { // arr[mid] == 2
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}

func threeWayPartition(arr []int, pivot int) {
	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {
		if arr[mid] < pivot {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == pivot {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}

func partitionRange(arr []int, lowVal, highVal int) {
	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {
		if arr[mid] < lowVal {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] >= lowVal && arr[mid] <= highVal {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}

func partitionZeroOne(arr []int) {
	i := -1
	for j := 0; j < len(arr); j++ {
		if arr[j] == 0 {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
}

func partitionEvenOdd(arr []int) {
	i := -1
	for j := 0; j < len(arr); j++ {
		if arr[j]%2 == 0 {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
}

func partitionPosNeg(arr []int) {
	i := -1
	for j := 0; j < len(arr); j++ {
		if arr[j] < 0 {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
}

func maximumProduct(nums []int) int {
	nums = mergeSort(nums)
	length := len(nums) - 1
	product1 := nums[length] * nums[length-1] * nums[length-2]
	product2 := nums[0] * nums[1] * nums[2]
	fmt.Println(nums)
	if math.Abs(float64(product2)) > math.Abs(float64(product1)) {
		return product2
	}
	return product1
}

func findErrorNums(nums []int) []int {
	length := len(nums)
	sum := (length * (length + 1)) / 2
	fmt.Println(sum)

	sum2 := 0
	for i := 0; i < len(nums); i++ {
		sum2 += nums[i]
	}
	var missingNum int
	valueMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if valueMap[nums[i]] {
			missingNum = nums[i]
		}
		valueMap[nums[i]] = true
	}
	if sum2 > sum {
		return []int{missingNum, missingNum - (sum2 - sum)}
	}
	return []int{missingNum, missingNum + (sum - sum2)}
}

func dominantIndex(nums []int) int {
	largest, secondlargest := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] >= nums[largest] {
			secondlargest = largest
			largest = i
		}
	}

	if nums[largest] >= 2*nums[secondlargest] {
		return largest
	}
	return -1
}

func checkIfExist(arr []int) bool {
	valueMap := make(map[float64]bool)

	for i := 0; i < len(arr); i++ {
		double := float64(arr[i]) * 2
		half := float64(arr[i]) / 2

		if _, ok := valueMap[double]; ok {
			return true
		} else if _, ok = valueMap[half]; ok {
			return true
		}
		valueMap[float64(arr[i])] = true
	}
	return false
}

func countNegatives(grid [][]int) int {
	isnegative := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				isnegative++
			}
		}
	}
	return isnegative
}

func findKthPositive(arr []int, k int) int {
	missingArray := []int{}
	first := 1
	for i := 0; i < len(arr); i++ {
		if arr[i] == first {
			first++
			continue
		}

		missingArray = append(missingArray, first)
		first++
		i--
	}
	if len(missingArray) >= k {
		return missingArray[k-1]
	}

	return arr[len(arr)-1] + (k - len(missingArray))
}

func findKthPositiveOptimised(arr []int, k int) int {
	j := 1
	for i := 0; i < len(arr); j++ {
		if arr[i] != j {
			k--
		} else {
			i++
		}
		if k == 0 {
			return j
		}
	}
	return j + k - 1
}

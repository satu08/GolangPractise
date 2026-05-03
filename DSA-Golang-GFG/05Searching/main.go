package main

import "fmt"

func main() {
	// fmt.Println(linearSearch([]int{1, 78, 89, 3, 8}, 3))
	// fmt.Println(binarySearch([]int{1, 3, 8, 78, 89}, 3))
	// fmt.Println(firstOccurrence([]int{1, 10, 10, 10, 10, 20, 20}, 10))
	// fmt.Println(lastOccurrence([]int{1, 10, 10, 10, 10, 20, 20}, 10))
	// fmt.Println(count1sinSortedBinaryArray([]int{0, 0, 1, 1, 1, 1, 1, 1}))

	// fmt.Println(mergeTwoSortedArrays([]int{1, 2, 3, 4, 5, 9, 70}, []int{3, 4, 5, 6, 7, 8}))
	fmt.Println(arrangingCoins(10))
}
func mergeTwoSortedArrays(arr1 []int, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		result = append(result, arr1[i])
		i++

	}
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}
	return result
}

// works on both sorted and unsorted arrays
func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// works only on sorted arrays
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func binarySearchRecursive(arr []int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] < target {
		return binarySearchRecursive(arr, target, mid+1, right)
	}
	return binarySearchRecursive(arr, target, left, mid-1)
}

func firstOccurrence(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			if mid == 0 || arr[mid-1] != arr[mid] {
				return mid
			} else {
				right = mid - 1
			}
		}
	}
	return result
}

func lastOccurrence(arr []int, target int) int {
	left, right := 0, len(arr)-1
	length := len(arr) - 1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			if mid != length || arr[mid] != arr[mid+1] {
				return mid
			} else {
				left = mid + 1
			}
		}
	}
	return result
}

func totalOccurrences(arr []int, target int) int {
	first := firstOccurrence(arr, target)
	if first == -1 {
		return 0
	}
	last := lastOccurrence(arr, target)
	return last - first + 1
}

func count1sinSortedBinaryArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == 0 {
			left = mid + 1
		} else {
			if mid == 0 || arr[mid-1] == 0 {
				return len(arr) - mid
			} else {
				right = mid - 1
			}
		}
	}
	return 0
}

func searchInSortedRotatedArray(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[left] <= arr[mid] {
			if target >= arr[left] && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > arr[mid] && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func getPeakElement(arr []int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if (mid == 0 || arr[mid] >= arr[mid-1]) && (mid == len(arr)-1 || arr[mid] >= arr[mid+1]) {
			return mid
		}
		if mid > 0 && arr[mid-1] > arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// two pointer approach, works only on sorted arrays
func pairwithgivensum(arr []int, target int) bool {
	left, right := 0, len(arr)-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return true
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}

func tripletwithgivensum(arr []int, target int) bool {
	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum == target {
				return true
			}
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return false
}

func medianofTwoSortedArraysMergeApproach(arr1 []int, arr2 []int) float64 {
	merged := mergeTwoSortedArrays(arr1, arr2)
	totalLength := len(merged)
	if totalLength%2 == 1 {
		return float64(merged[totalLength/2])
	}
	return float64(merged[totalLength/2-1]+merged[totalLength/2]) / 2.0
}

func medianofTwoSortedArraysMergeApproachOptimized(arr1 []int, arr2 []int) float64 {
	totalLength := len(arr1) + len(arr2)
	mid := totalLength / 2
	i, j := 0, 0
	var prev, current int
	for k := 0; k <= mid; k++ {
		prev = current
		if i < len(arr1) && (j >= len(arr2) || arr1[i] < arr2[j]) {
			current = arr1[i]
			i++
		} else {
			current = arr2[j]
			j++
		}
	}
	if totalLength%2 == 1 {
		return float64(current)
	}
	return float64(prev+current) / 2.0
}

func squareRoot(n int) int {
	if n < 2 {
		return n
	}
	left, right := 1, n/2
	for left <= right {
		mid := left + (right-left)/2
		midSquare := mid * mid
		if midSquare == n {
			return mid
		}
		if midSquare < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func arrangingCoins(n int) int {
	i := 1
	sum := 0
	for i = 1; i <= n; i++ {
		sum += i
		if sum > n {
			return i - 1
		}
	}
	return i
}

func arrangeCoins(n int) int {
	low, high := 0, n

	for low <= high {
		mid := low + (high-low)/2
		coins := mid * (mid + 1) / 2

		if coins == n {
			return mid
		} else if coins < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}

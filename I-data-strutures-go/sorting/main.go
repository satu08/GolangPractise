// | Algorithm      | Time           | Stable | Use Case              |
// | -------------- | -------------- | ------ | --------------------- |
// | Bubble Sort    | O(n²)          | ✅      | Learning only         |
// | Selection Sort | O(n²)          | ❌      | Rarely used           |
// | Insertion Sort | O(n²)          | ✅      | Small / nearly sorted |
// | Merge Sort     | O(n log n)     | ✅      | Stable sorting        |
// | Quick Sort     | O(n log n) avg | ❌      | Fast in practice      |
// | Heap Sort      | O(n log n)     | ❌      | Priority queues       |
// | Counting Sort  | O(n+k)         | ✅      | Small range integers  |

package main

import (
	"fmt"
)

// 🔹 Bubble Sort
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 🔹 Selection Sort
func selectionSort(arr []int) {
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
}

// 🔹 Insertion Sort
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
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

// 🔹 Heap Sort
func heapSort(arr []int) {
	n := len(arr)

	// build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// extract elements
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// 🔹 Counting Sort
func countingSort(arr []int) []int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	count := make([]int, max+1)

	for _, v := range arr {
		count[v]++
	}

	result := []int{}
	for i := 0; i <= max; i++ {
		for count[i] > 0 {
			result = append(result, i)
			count[i]--
		}
	}
	return result
}

// 🔹 MAIN
func main() {
	arr := []int{5, 2, 9, 1, 5, 6}

	fmt.Println("Original:", arr)

	// Bubble
	a1 := append([]int{}, arr...)
	bubbleSort(a1)
	fmt.Println("Bubble:", a1)

	// Selection
	a2 := append([]int{}, arr...)
	selectionSort(a2)
	fmt.Println("Selection:", a2)

	// Insertion
	a3 := append([]int{}, arr...)
	insertionSort(a3)
	fmt.Println("Insertion:", a3)

	// Merge
	a4 := mergeSort(arr)
	fmt.Println("Merge:", a4)

	// Quick
	a5 := append([]int{}, arr...)
	quickSort(a5, 0, len(a5)-1)
	fmt.Println("Quick:", a5)

	// Heap
	a6 := append([]int{}, arr...)
	heapSort(a6)
	fmt.Println("Heap:", a6)

	// Counting
	a7 := countingSort(arr)
	fmt.Println("Counting:", a7)
}

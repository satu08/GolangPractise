package main

import "fmt"

//	func main() {
//		merge([]int{0}, 0, []int{1}, 1)
//	}
func merge(nums1 []int, m int, nums2 []int, n int) {

	var arr = []int{}

	arr = append(arr, nums1[:m]...)
	arr = append(arr, nums2[:n]...)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		nums1[i] = arr[i]
	}

	fmt.Println(nums1)

}

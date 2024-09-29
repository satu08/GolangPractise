package main

import "fmt"

func main() {
	arr := LeftRotateByD([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println(arr)

	arr2 := leftRotate([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println(arr2)
}

func LeftRotateByD(arr []int, d int) []int {
	n := len(arr)
	for i := 0; i < d; i++ {
		LeftRotateByOne(arr, n)
	}
	return arr
}

func LeftRotateByOne(A []int, n int) []int {
	temp := A[0]
	for i := 1; i < n; i++ {
		A[i-1] = A[i]
	}
	A[n-1] = temp
	return A
}

func leftRotate(A []int, d int) []int {
	n := len(A)
	reverse(A, 0, d-1)
	reverse(A, d, n-1)
	reverse(A, 0, n-1)
	return A
}
func reverse(s []int, low, high int) {
	for low < high {
		temp := s[low]
		s[low] = s[high]
		s[high] = temp
		low++
		high--
	}
}

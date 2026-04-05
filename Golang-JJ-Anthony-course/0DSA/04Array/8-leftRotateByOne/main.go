package main

import "fmt"

func main() {
	arr := LeftRotateByOne([]int{30, 5, 20})
	fmt.Println(arr)
}

func LeftRotateByOne(A []int) []int {
	temp := A[0]
	for i := 1; i < len(A); i++ {
		A[i-1] = A[i]
	}
	A[len(A)-1] = temp
	return A
}

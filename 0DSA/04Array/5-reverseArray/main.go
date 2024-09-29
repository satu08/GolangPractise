package main

import "fmt"

func main() {
	reversedArray := reverseArray([]int{12, 78, 43, 67, 90})
	fmt.Println(reversedArray)
	reversedArrayWithoutSecondArray := reverseArrayWithoutSecondArray([]int{12, 78, 43, 67, 90})
	fmt.Println(reversedArrayWithoutSecondArray)
}

// time complexity is Θ(n)
func reverseArray(s []int) []int {
	l := len(s)
	var arr []int
	for i := l - 1; i >= 0; i-- {
		arr = append(arr, s[i])
	}
	return arr
}

func reverseArrayWithoutSecondArray(s []int) []int {
	high := len(s) - 1
	low := 0
	for low < high {
		temp := s[low]
		s[low] = s[high]
		s[high] = temp
		low++
		high--
	}
	return s
}

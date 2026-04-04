package main

//func main() {
//	fmt.Println(check([]int{1, 3, 4}))
//}

func check(nums []int) bool {
	count := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			break
		} else {
			count++
		}
	}
	nums = rightRotateByD(nums, len(nums)-count)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			return false
		}
	}
	return true
}

func rightRotateByD(arr []int, d int) []int {
	n := len(arr)
	for i := 0; i < d; i++ {
		rightRotateByOne(arr, n)
	}
	return arr
}

func rightRotateByOne(A []int, n int) []int {
	temp := A[n-1]
	for i := n - 1; i > 0; i-- {
		A[i] = A[i-1]
	}
	A[0] = temp
	return A
}

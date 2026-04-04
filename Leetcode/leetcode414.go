package main

//
//func main() {
//	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
//}

func thirdMax(nums []int) int {
	largest := 0
	secondLargest := 0
	thirdlargest := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > largest {
			thirdlargest = secondLargest
			secondLargest = largest
			largest = nums[i]
		}
	}
	return thirdlargest
}

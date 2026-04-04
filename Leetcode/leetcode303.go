package main

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.nums[i]
	}
	return sum
}

//func main() {
//	arr := Constructor([]int{-2, 0, 3, -5, 2, -1})
//	val := arr.SumRange(0, 2)
//	fmt.Println(val)
//}

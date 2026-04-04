package main

//	func main() {
//		fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
//	}
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var returnArray []int
	for _, v := range nums1 {
		for i1, v1 := range nums2 {
			if v1 == v {
				var greaterexist bool
				for i2 := i1; i2 < len(nums2); i2++ {
					if nums2[i2] > v1 {
						greaterexist = true
						returnArray = append(returnArray, nums2[i2])
						break
					}
				}
				if !greaterexist {
					returnArray = append(returnArray, -1)
				}

			}
		}
	}
	return returnArray
}

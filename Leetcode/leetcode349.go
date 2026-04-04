package main

//func main() {
//	fmt.Println(intersection2([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
//	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
//	fmt.Println(intersection3([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
//	fmt.Println(intersection4([]int{4, 9, 5, 9}, []int{9, 4, 9, 8, 4}))
//}

func intersection(nums1 []int, nums2 []int) []int {
	intmap := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == nums1[i] {
				if _, exist := intmap[nums1[i]]; !exist {
					intmap[nums1[i]]++
				}
			}
		}
	}
	var arr []int
	for k, _ := range intmap {
		arr = append(arr, k)
	}
	return arr
}

func intersection2(nums1 []int, nums2 []int) []int {
	intmap := make(map[int]int)
	int2map := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		intmap[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		int2map[nums2[i]]++
	}
	var arr []int
	for k, v := range intmap {
		for k2, v2 := range int2map {
			if k == k2 {
				mini := min(v, v2)
				for mini > 0 {
					arr = append(arr, v2)
					mini--
				}
			}
		}
	}
	return arr
}

func intersection3(nums1 []int, nums2 []int) []int {
	intmap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		intmap[nums1[i]]++
	}
	var arr []int

	for i := 0; i < len(nums2); i++ {
		if _, exist := intmap[nums2[i]]; exist {
			arr = append(arr, nums2[i])
			delete(intmap, nums2[i])
		}
	}
	return arr
}

func intersection4(nums1 []int, nums2 []int) []int {
	intmap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		intmap[nums1[i]]++
	}
	var arr []int

	for i := 0; i < len(nums2); i++ {
		if _, exist := intmap[nums2[i]]; exist {
			if intmap[nums2[i]] > 0 {
				arr = append(arr, nums2[i])
			}
			intmap[nums2[i]]--
		}
	}
	return arr
}

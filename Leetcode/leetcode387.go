package main

//func main() {
//	fmt.Println(firstUniqChar("aabb"))
//}

func firstUniqChar(s string) int {
	uniqueMap := make(map[string][]int)
	for i, v := range s {
		uniqueMap[string(v)] = append(uniqueMap[string(v)], i)
	}
	for _, v := range s {
		if len(uniqueMap[string(v)]) == 1 {
			return uniqueMap[string(v)][0]
		}
	}
	return -1
}

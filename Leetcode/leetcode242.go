package main

//	func main() {
//		fmt.Println(isAnagram("anagram", "nagaram"))
//	}
func isAnagram(s string, t string) bool {
	anaMap := make(map[string]int)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		anaMap[string(s[i])]++
		anaMap[string(t[i])]--
	}
	for _, v := range anaMap {
		if v != 0 {
			return false
		}
	}
	return true
}

package main

//	func main() {
//		fmt.Println(longestCommonPrefix([]string{"car", "cat", "cam"}))
//	}
func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for _, s := range strs[1:] {
			if i >= len(s) || s[i] != char {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

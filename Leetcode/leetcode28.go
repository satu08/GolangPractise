package main

//	func main() {
//		fmt.Println(strStr("a", "a"))
//	}
func strStr(haystack string, needle string) int {
	length := len(needle)

	for i := 0; i < len(haystack); i++ {
		if (i + length) <= len(haystack) {
			if haystack[i:i+length] == needle {
				return i
			}
		}
	}
	return -1
}

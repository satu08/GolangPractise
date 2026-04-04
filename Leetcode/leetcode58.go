package main

//	func main() {
//		fmt.Println(lengthOfLastWord("luffy is still joyboy"))
//	}
func lengthOfLastWord(s string) int {
	length := 0
	var islettervisited bool
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			islettervisited = true
			length++
		}
		if islettervisited && string(s[i]) == " " {
			break
		}

	}
	return length
}

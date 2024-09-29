package main

import "fmt"

func main() {
	flag := isPalindrome("satya", 0, 3)
	fmt.Println(flag)
}

func isPalindrome(s string, i, j int) bool {
	if i >= j {
		return true
	}
	return string(s[i]) == string(s[j]) && isPalindrome(s, i+1, j-1)
}

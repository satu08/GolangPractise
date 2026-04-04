package main

import "strings"

// import (
//
//	"fmt"
//	"strings"
//
// )
//
//	func main() {
//		fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
//	}
func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		var locali string
		var localj string
		if s[i] >= 65 && s[i] <= 90 || s[i] >= 97 && s[i] <= 122 || s[i] >= 48 && s[i] <= 57 {
			locali = strings.ToLower(string(s[i]))
			i++
			continue
		}
		if s[j] >= 65 && s[j] <= 90 || s[j] >= 97 && s[j] <= 122 || s[j] >= 48 && s[j] <= 57 {
			localj = strings.ToLower(string(s[j]))
			j--
			continue
		}
		if locali != localj {
			return false
		}
		i++
		j--
	}

	return true
}

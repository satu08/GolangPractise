package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "abcabcbb"

	fmt.Println("Longest Unique:", longestUnique(s))

	fmt.Println("Palindrome (racecar):", isPalindrome("racecar"))

	fmt.Println("Anagram (listen, silent):", isAnagram("listen", "silent"))

	// String Builder
	var sb strings.Builder
	sb.WriteString("Go ")
	sb.WriteString("Lang")
	fmt.Println("Builder:", sb.String())
}

// 🔹 Longest substring
func longestUnique(s string) int {
	set := make(map[byte]bool)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		for set[s[right]] {
			delete(set, s[left])
			left++
		}
		set[s[right]] = true
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// 🔹 Palindrome
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 🔹 Anagram
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

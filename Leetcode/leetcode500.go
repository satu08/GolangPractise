package main

import (
	"strings"
)

//func main() {
//	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
//}

func findWords(words []string) []string {
	s1 := "qwertyuiop"
	s2 := "asdfghjkl"
	s3 := "zxcvbnm"

	valueMap := make(map[rune]int)
	for _, c := range s1 {
		valueMap[c] = 1
	}
	for _, c := range s2 {
		valueMap[c] = 2
	}
	for _, c := range s3 {
		valueMap[c] = 3
	}

	var result []string

	for _, word := range words {
		runes := []rune(strings.ToLower(word))
		row := valueMap[runes[0]]
		isWritable := true

		for i := 1; i < len(runes); i++ {
			if valueMap[runes[i]] != row {
				isWritable = false
				break
			}
		}

		if isWritable {
			result = append(result, word)
		}
	}

	return result
}

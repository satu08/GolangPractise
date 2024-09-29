package main

import "fmt"

func main() {
	subsets("abcd", "", 0)
}

func subsets(s string, curr string, i int) {
	if i == len(s) {
		fmt.Println(curr)
		return
	}
	subsets(s, curr, i+1)
	subsets(s, curr+string(s[i]), i+1)
}

package main

import "fmt"

func main() {
	fmt.Println(namedreturn([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func namedreturn(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}

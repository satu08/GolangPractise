package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(hammingWeight(11))
	fmt.Println(hammingWeight2(11))
}

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}

func hammingWeight2(num int) int {
	count := 0

	binary := strconv.FormatInt(int64(num), 2)

	for i := 0; i < len(binary); i++ {
		if binary[i] == '1' {
			count++
		}
	}

	return count
}

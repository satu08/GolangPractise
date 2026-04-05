package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(1431))

}

func isPalindrome(x int) bool {
	temp := x
	reverse := 0
	for x > 0 {
		digit := x % 10
		reverse = reverse*10 + digit
		x = x / 10
	}
	return reverse == temp
}

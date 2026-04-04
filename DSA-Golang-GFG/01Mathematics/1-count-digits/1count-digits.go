package main

import "fmt"

func main() {
	c := CountDigits(2345111)
	fmt.Println(c)
}
func CountDigits(num int) int {
	digits := 0
	for num > 0 {
		digits++
		num /= 10
	}
	return digits
}

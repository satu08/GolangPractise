package main

//func main() {
//	fmt.Println(isHappy(19))
//}

func isHappy(n int) bool {
	for {
		return getSum(n) == 1
	}
	return false
}

func getSum(n int) int {
	var sum int
	for n > 0 {
		modulo := n % 10
		sum += modulo * modulo
		n /= 10
	}
	return sum
}

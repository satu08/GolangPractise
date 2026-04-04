package main

//	func main() {
//		fmt.Println(plusOne([]int{9, 9}, 9))
//	}
func plusOne(digits []int, target int) []int {
	carry := target
	i := len(digits) - 1
	for i >= 0 {
		sum := digits[i] + carry
		if sum >= 10 {
			digits[i] = sum % 10
			carry = sum / 10
		} else {
			digits[i] = sum
			carry = sum / 10
		}
		i--
	}
	if carry >= 1 {
		return append([]int{1}, digits...)
	}

	return digits
}

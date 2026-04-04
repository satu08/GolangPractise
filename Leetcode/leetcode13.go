package main

//	func main() {
//		fmt.Println(romanToInt("MCMXCIV"))
//	}
func romanToInt(s string) int {
	valueMap := make(map[string]int)
	valueMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	lenth := len(s)
	sum := 0
	for i := 0; i < lenth; i++ {
		if i+1 > len(s)-1 {
			sum += valueMap[string(s[i])]
			return sum
		}
		if s[i] == s[i+1] {
			sum += valueMap[string(s[i])]
		} else {
			if valueMap[string(s[i])] > valueMap[string(s[i+1])] {
				sum += valueMap[string(s[i])]
			} else {
				sum += valueMap[string(s[i+1])] - valueMap[string(s[i])]
				i++
			}
		}
	}
	return sum
}

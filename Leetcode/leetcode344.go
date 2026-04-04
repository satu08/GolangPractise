package main

//func main() {
//	fmt.Println(string(reverseString([]byte{'h', 'e', 'l', 'l', 'o'})))
//}

func reverseString(s []byte) []byte {
	i := 0
	l := len(s) - 1
	for i < l {
		s[i], s[l] = s[l], s[i]
		i++
		l--
	}
	return s
}

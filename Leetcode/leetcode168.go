package main

//	func main() {
//		fmt.Println(convertToTitle(701))
//	}
func convertToTitle(columnNumber int) string {
	var arr []byte
	for columnNumber > 0 {
		columnNumber--
		remainder := columnNumber % 26
		arr = append(arr, byte('A'+remainder))
		columnNumber /= 26
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return string(arr)
}

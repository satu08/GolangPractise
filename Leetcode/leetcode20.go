package main

//	func main() {
//		fmt.Println(isValid("()[]{}"))
//	}
func isValid(s string) bool {
	run := []rune(s)
	closingMap := make(map[string]string)

	closingMap[")"] = "("
	closingMap["}"] = "{"
	closingMap["]"] = "["
	var tempStack []string
	for _, ch := range run {
		if _, ok := closingMap[string(ch)]; ok {
			if len(tempStack) > 0 {
				if tempStack[len(tempStack)-1] == closingMap[string(ch)] {
					tempStack = tempStack[:len(tempStack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			tempStack = append(tempStack, string(ch))
		}
	}
	if len(tempStack) != 0 {
		return false
	}
	return true
}

package main

//	func main() {
//		fmt.Println(pascalTriangle(2))
//	}
func pascalTriangle(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)

		row[0] = 1
		row[i] = 1

		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		triangle[i] = row
	}
	return triangle
}

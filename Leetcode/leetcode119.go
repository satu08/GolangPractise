package main

//	func main() {
//		fmt.Println(pascalTriangle(1))
//	}
func pascalTriangle1(rowindex int) []int {
	triangle := make([][]int, rowindex+1)

	for i := 0; i < rowindex+1; i++ {
		row := make([]int, i+1)

		row[0] = 1
		row[i] = 1

		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		triangle[i] = row
	}
	return triangle[rowindex]
}

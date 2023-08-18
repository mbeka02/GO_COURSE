package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	// ?
	//2-D matrix
	matrix := [][]int{}

	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < cols; j++ {
			//mult row indx with col indx
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}



func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
	//test (10000,10000)
}

package main

func rotate2(matrix [][]int) {
	row, col := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		for j := i; j < col-i-1; j++ {
			matrix[j][col-i-1], matrix[row-i-1][col-j-1], matrix[row-j-1][i], matrix[i][j] =
				matrix[i][j], matrix[j][col-i-1], matrix[row-i-1][col-j-1], matrix[row-j-1][i]
		}
	}
}

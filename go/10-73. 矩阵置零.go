package main

func setZeroes(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])
	m := make([]bool, row)
	n := make([]bool, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				m[i] = true
				n[j] = true
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if m[i] == true || n[j] == true {
				matrix[i][j] = 0
			}
		}
	}
}

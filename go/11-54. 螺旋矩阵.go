package main

func spiralOrder(matrix [][]int) []int {
	row, col := len(matrix), len(matrix[0])
	var (
		total      = row * col
		res        = make([]int, 0, total)
		directions = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
		k          = 0
		i          = 0
		j          = -1
	)

	for len(res) < total {
		k %= 4
		for i+directions[k][0] >= 0 && i+directions[k][0] < row &&
			j+directions[k][1] >= 0 && j+directions[k][1] < col {
			if matrix[i+directions[k][0]][j+directions[k][1]] <= 100 {
				i += directions[k][0]
				j += directions[k][1]
				res = append(res, matrix[i][j])
				matrix[i][j] = 101
			} else {
				break
			}
		}
		k++
	}
	return res
}

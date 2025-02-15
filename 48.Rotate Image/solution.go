func rotate(matrix [][]int)  {

	row := len(matrix)
	col := len(matrix[0])

	for i := 0; i < row; i++ {
		for j := i + 1 ; j < col; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}

	for i := 0; i < row; i++ {
		for j, k := 0, len(matrix[i]) - 1 ; j< k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}

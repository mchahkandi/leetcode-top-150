func numIslands(grid [][]byte) int {
	islands := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				islands++
				dfs(grid,i,j)
			}
		}
	}

	return islands
}

func dfs(grid [][]byte,i, j int) {

	grid[i][j] = '0'

	directions := [][]int{[]int{i,j - 1}, []int{i + 1, j}, []int{i - 1, j}, []int{i, j + 1}}

	for _, direction := range directions {
		row := direction[0]
		col := direction[1]

		if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) && grid[row][col] == '1' {
			dfs(grid,row,col)
		}
	}
}

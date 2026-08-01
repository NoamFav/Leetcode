package leetcode

// @leet start
func inBounds(grid [][]int, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	var border [][2]int
	item := grid[row][col]

	var dfs func(row, col int)
	dfs = func(row, col int) {
		visited[row][col] = true
		for _, dir := range dirs {
			nr, nc := row+dir[0], col+dir[1]
			if !inBounds(grid, nr, nc) {
				border = append(border, [2]int{row, col})
			} else {
				if grid[nr][nc] == item && !visited[nr][nc] {
					dfs(nr, nc)
				} else if grid[nr][nc] != item {
					border = append(border, [2]int{row, col})
				} else if !visited[nr][nc] {
					dfs(nr, nc)
				}
			}
		}
	}

	dfs(row, col)

	for _, i := range border {
		grid[i[0]][i[1]] = color
	}

	return grid
}

// @leet end

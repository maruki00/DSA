package graph

func MatrixDFS(grid [][]int, r, c int, visited map[[2]int]bool) int {

	rows, cols := len(grid), len(grid[0])

	if min(r, c) < 0 || visited[[2]int{r, c}] == true || grid[r][c] == 1 {
		return 0
	}

	if r == rows-1 || c == cols-1 {
		return 1
	}

	visited[[2]int{r, c}] = true
	count := 0
	count += dfs(grid, r-1, c, visited)
	count += dfs(grid, r+1, c, visited)
	count += dfs(grid, r, c-1, visited)
	count += dfs(grid, r, c+1, visited)
	delete(visited, [2]int{r, c})
	return count
}

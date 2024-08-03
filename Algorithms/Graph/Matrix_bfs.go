package graph

type Node struct {
	r int
	c int
}

func pop(node []Node) (int, int, []Node) {
	return node[0].r, node[0].c, node[1:]
}

func MatrixDFS(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	visited := make(map[[2]int]bool, 0)
	q := []Node{}

	q = append(q, Node{0, 0})
	visited[[2]int{0, 0}] = true
	l := 0
	for len(q) > 0 {
		for range len(q) {
			r, c, q := pop(q)
			if r == rows-1 && c == cols-1 {
				return l
			}

			neighbors := [][]int{
				{0, 1},
				{0, -1},
				{1, 0},
				{-1, 0},
			}
			for _, neibord := range neighbors {
				dr := neibord[0]
				dc := neibord[1]
				if min(r+dr, c+dc) < 0 || r+dr == rows || dc+c == cols || visited[[2]int{r + dr, c + d}] == true || grid[r+dr][c+d] == 1 {
					continue
				}
				q = append(q, Node{r + dr, c + dc})
				visited[[2]int{r + dr, c + dc}] = true
			}
		}
		l++
	}
	return 0
}

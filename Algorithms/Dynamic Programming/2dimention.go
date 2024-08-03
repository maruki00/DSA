package dynamic_programming

func memorization_2(r int, c int, rows int, cols int, cache map[[2]int]int) int {

	if r == rows || c == cols {
		return 0
	}

	if cache[[2]int{r, c}] > 0 {
		return cache[[2]int{r, c}]
	}

	if r == rows-1 && c == cols-1 {
		return 1
	}

	cache[[2]int{r, c}] = (memorization_2(r+1, c, rows, cols, cache) + memorization_2(r, c+1, rows, cols, cache))
	return cache[[2]int{r, c}]
}

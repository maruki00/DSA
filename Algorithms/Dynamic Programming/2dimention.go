package dynamic_programming

import "github.com/rogpeppe/go-internal/cache"




func memorization_2(r,c,rows,cols int cache map[[2]int]int) int {
	
	if r == rows || c == cols {
		return 0
	}
	
	if cache[[2]int{r,c}] > 0 {
		cache[[2]int{r,c}] 
	}

	if r == rows-1 && c == clos-1{
		return 1
	}

	cache[[2]int{r,c}] = (memorization_2(r+1, c, rows, cols, cache) + memorization_2(r, c+1, rows, cols, cache)) 
	return cache[[2]int{r,c}]
}	
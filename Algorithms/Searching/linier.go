package searching

func LinerSearch(items []int, value int) int {
	for j, i := range items {
		if i == value {
			return j
		}
	}
	return -1
}

package searching

func BinarySearch(items []int, value int) int {
	l, r := 0, len(items)-1
	middle := 0
	for l < r {
		middle = (l + r) / 2
		if items[middle] == value {
			return middle
		}

		if items[middle] > value {
			r = middle - 1
		} else {
			l = middle + 1
		}
	}
	return -1
}

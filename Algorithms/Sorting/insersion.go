package sorting

func InsertionSort(items []int) []int {

	for i := 1; i < len(items); i++ {

		for j := i; items[j] < items[j-1] && j > 0; j-- {
			items[j], items[j-1] = items[j-1], items[j]
		}
	}
	return items
}

package sorting

func CustomSort(items []int) []int {

	for i := 0; i < len(items); i++ {
		for j := i; j < len(items); j++ {
			if items[i] > items[j] {
				items[i], items[j] = items[j], items[i]
			}
		}
	}
	return items
}

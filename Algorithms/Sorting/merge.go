package sorting

func MergeSort(items1, items2 []int) []int {
	l := len(items1) - 1
	r := len(items2) - 1
	result := make([]int, 0)
	i, j := 0, 0
	for l >= i && r >= j {
		if items1[i] <= items2[j] {
			result = append(result, items1[i])
			i++
		} else {
			result = append(result, items2[j])
			j++
		}
	}

	for l >= i {
		result = append(result, items1[i])
		i++
	}

	for r >= j {
		result = append(result, items2[j])
		j++
	}
	return result
}

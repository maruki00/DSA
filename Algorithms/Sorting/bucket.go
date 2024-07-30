package sorting

import "sort"

func BucketSort(items []int) []int {
	// var bucket map[int]int
	bucket := make(map[int]int)
	result := []int{}
	keys := []int{}
	for _, i := range items {
		bucket[i]++
	}
	for i := range bucket {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for _, j := range keys {
		for k := 0; k < bucket[j]; k++ {
			result = append(result, j)
		}
	}
	return result
}

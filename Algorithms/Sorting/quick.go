package sorting

func QuickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)-1]
	left := []int{}
	right := []int{}

	for _, item := range arr[:len(arr)-1] {
		if item > pivot {
			right = append(right, item)
			continue
		}
		left = append(left, item)
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

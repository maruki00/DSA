package dynamic_programming

func memorization_1(number int, cache map[int]int) int {

	if number <= 1 {
		return number
	}

	if val, ok := cache[number]; ok {
		return val
	}

	cache[number] = memorization_1(number-1, cache) + memorization_1(number-2, cache)
	return cache[number]
}

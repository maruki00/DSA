package dynamic_programming

func memorization(number int, cache map[int]int) int {

	if number <= 1 {
		return number
	}

	if val, ok := cache[number]; ok {
		return val
	}

	cache[number] = memorization(number-1, cache) + memorization(number-2, cache)
	return cache[number]
}

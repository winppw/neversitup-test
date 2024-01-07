package main

func FindOddNumber(text []int) int {
	// TODO : start your code here
	numCounts := make(map[int]int)
	for _, num := range text {
		numCounts[num]++
	}

	for num, count := range numCounts {
		if count%2 != 0 {
			return num
		}
	}

	return 0
}

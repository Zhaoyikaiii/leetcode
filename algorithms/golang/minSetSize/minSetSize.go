package minSetSize

import "sort"

func Solution(arr []int) int {
	m := make(map[int]int)
	for _, num := range arr {
		m[num]++
	}
	var freq []int
	for _, v := range m {
		freq = append(freq, v)
	}
	sort.Ints(freq)
	var ans, removed int
	for i := len(freq) - 1; i >= 0; i-- {
		removed += freq[i]
		ans++
		if removed >= len(arr)/2 {
			return ans
		}
	}
	return ans
}

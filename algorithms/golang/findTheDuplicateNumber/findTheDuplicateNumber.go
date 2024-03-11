package findtheduplicatenumber

func findDuplicate(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		if m[v] {
			return v
		}
		m[v] = true
	}
	return -1
}

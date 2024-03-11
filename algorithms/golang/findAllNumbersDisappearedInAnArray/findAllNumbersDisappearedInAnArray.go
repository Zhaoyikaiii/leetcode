package findallnumbersdisappearedinanarray

func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		m[v] = true
	}
	result := make([]int, 0, len(nums))
	for i := 1; i <= len(nums); i++ {
		if !m[i] {
			result = append(result, i)
		}
	}
	return result
}

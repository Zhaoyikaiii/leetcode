package squaresofasortedarray

// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	l, r, k := 0, len(nums)-1, len(nums)-1
	for k >= 0 {
		if abs(nums[l]) > abs(nums[r]) {
			res[k] = abs(nums[l]) * abs(nums[l])
			l++
		} else {
			res[k] = abs(nums[r]) * abs(nums[r])
			r--
		}
		k--
	}
	return res
}

func abs(i int) int {
	if i > 0 {
		return i
	} else {
		return -i
	}
}

package main

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	start, end, sum := 0, 0, 0
	minLength := n + 1
	for end < n {
		sum += nums[end]
		for sum >= target {
			if end-start+1 < minLength {
				minLength = end - start + 1
			}
			sum -= nums[start]
			start++
		}
		end++
	}
	if minLength == n+1 {
		return 0
	}
	return minLength
}

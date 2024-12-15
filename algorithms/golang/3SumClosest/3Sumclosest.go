package sumclosest

import (
	"math"
	"sort"
)

func thressSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-closest)) {
				closest = sum
			}
			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				return sum
			}
		}
	}
	return closest
}

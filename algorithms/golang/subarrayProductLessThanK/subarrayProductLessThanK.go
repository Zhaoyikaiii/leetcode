package subarrayproductlessthank

func numSubarrayProductLessTanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	l, r, c := 0, 0, 0
	acc := 1
	for r < len(nums) {
		acc *= nums[r]
		for acc >= k { // if discovery current acc is bigger than k,recovery to last step and move on l pointer.
			acc /= nums[l]
			l++
		}
		c += r - l + 1 // add the distance between l and r pointer.
		r++
	}
	return c
}

package main

//You are given an integer array nums and an integer k. Find the maximum subarray sum of all the subarrays of nums that meet the following conditions:
//The length of the subarray is k, and
//All the elements of the subarray are distinct.
//Return the maximum subarray sum of all the subarrays that meet the conditions. If no subarray meets the conditions, return 0.
//A subarray is a contiguous non-empty sequence of elements within an array.

func maximumSubarraySum(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}

	window := make(map[int]int)
	sum, maxSum := 0, 0
	for i := 0; i < len(nums); i++ {
		window[nums[i]]++
		sum += nums[i]

		if i >= k {
			if window[nums[i-k]] > 1 {
				window[nums[i-k]] -= 1
			} else {
				delete(window, nums[i-k]) // When the window is full, eliminate the first element(nums[i-k]) from the window.
			}
			sum -= nums[i-k]
		}

		if i >= k-1 && len(window) == k { //len(window) == k ensure every element in the window is distinct.
			maxSum = max(maxSum, sum) // in case the duplicate element was pushed out of the window at earlier step.Later it will be replaced by the new element.until  len(window) == k
		}
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

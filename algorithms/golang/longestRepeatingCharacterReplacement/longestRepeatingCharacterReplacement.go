package longestrepeatingcharacterreplacement

// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.
// Return the length of the longest substring containing the same letter you can get after performing the above operations.

func characterReplacement(s string, k int) int {
	count := make([]int, 26)
	maxCount, maxLen := 0, 0

	for left, right := 0, 0; right < len(s); right++ {
		count[s[right]-'A']++
		maxCount = max(maxCount, count[s[right]-'A'])
		if right-left+1-maxCount > k { // k also represents the maximum number of characters that are not the same in the window.If right - left + 1 - maxCount begin to exceed k, it means that the window is too large and needs to be reduced.
			count[s[left]-'A']--
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

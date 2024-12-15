package longestsubstringwithatleastkrepeating

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}

	for i, ch := range s {
		if count[ch] < k {
			return max(longestSubstring(s[:i], k), longestSubstring(s[i+1:], k))
		}
	}

	return len(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

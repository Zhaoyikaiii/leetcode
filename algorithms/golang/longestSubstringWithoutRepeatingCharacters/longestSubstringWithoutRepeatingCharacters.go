package validparentheses

import "strings"

func lengthOfLongestSubstring(s string) int {
	strs := strings.Split(s, "")
	l, r, m, charMap := 0, 0, 0, make(map[string]int)

	for r < len(strs) {
		if _, ok := charMap[strs[r]]; ok {
			l = max(l, charMap[strs[r]]+1)
		}
		charMap[strs[r]] = r
		m = max(m, r-l+1)
		r = r + 1
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

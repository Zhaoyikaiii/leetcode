package minimumwindowsubstring

// Given two strings s and t of lengths m and n respectively, return the minimum window substring
//  of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".
// The testcases will be generated such that the answer is unique.

// in s = "ADOBECODEBANC", t = "ABC", the minimum window substring is "BANC"

func minWindow(s, t string) string {
	// create a map to store the frequency of each character in t
	tMap := make(map[string]int)
	for _, v := range t {
		tMap[string(v)]++
	}

	// create a map to store the frequency of each character in s
	sMap := make(map[string]int)

	// create a counter to count the number of characters in t
	counter := len(tMap)

	// create two pointers to track the start and end of the window
	l, r, m, start := 0, 0, len(s)+1, 0

	// iterate through the string s
	for r < len(s) {
		// if the character at r is in t, update the frequency of the character in sMap
		if _, ok := tMap[string(s[r])]; ok {
			sMap[string(s[r])]++
			// if the frequency of the character in sMap is equal to the frequency of the character in tMap, decrement the counter
			if sMap[string(s[r])] == tMap[string(s[r])] {
				counter--
			}
		}

		// increment r
		r++

		// while the counter is 0
		for counter == 0 {
			// if the length of the window is less than m, update m and start
			if r-l < m {
				m = r - l
				start = l
			}

			// if the character at l is in t, update the frequency of the character in sMap
			if _, ok := tMap[string(s[l])]; ok {
				sMap[string(s[l])]--
				// if the frequency of the character in sMap is less than the frequency of the character in tMap, increment the counter
				if sMap[string(s[l])] < tMap[string(s[l])] {
					counter++
				}
			}

			// increment l
			l++
		}
	}

	// if m is equal to the length of s, return an empty string
	if m == len(s)+1 {
		return ""
	}

	// return the minimum window substring
	return s[start : start+m]
}

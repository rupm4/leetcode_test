package main

func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	m := make(map[byte]int, 0)
	ans := 0

	for r < len(s) {
		m[s[r]]++
		for m[s[r]] > 1 {
			m[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
		r++
	}

	return ans
}

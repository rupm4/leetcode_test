package main

import "math"

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}

	for i := range t {
		ori[t[i]]++
	}

	long := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for i, v := range ori {
			if cnt[i] < v {
				return false
			}
		}

		return true
	}

	for left, right := 0, 0; right < len(s); right++ {
		if right < len(s) {
			cnt[s[right]]++
		}

		for check() && left <= right {
			if long > right-left+1 {
				long = right - left + 1
				ansL, ansR = left, right
			}
			cnt[s[left]]--
			left++
		}
	}

	if ansL != -1 {
		return s[ansL : ansR+1]
	} else {
		return ""
	}
}

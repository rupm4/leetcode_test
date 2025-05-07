package main

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	res := 0
	for l < r {
		ans := (r - l) * min(height[l], height[r])
		res = max(res, ans)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}

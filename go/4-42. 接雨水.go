package main

func trap(height []int) int {
	l := 0
	r := len(height) - 1
	lmax := 0
	rmax := 0
	ans := 0
	for l < r {
		lmax = max(lmax, height[l])
		rmax = max(rmax, height[r])
		if height[l] >= height[r] {
			ans += rmax - height[r]
			r--
		} else {
			ans += lmax - height[l]
			l++
		}
	}

	return ans
}

package main

import "sort"

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for i := 0; i < len(intervals); i++ {
		l := len(ans)
		if l == 0 || ans[l-1][1] < intervals[i][0] {
			ans = append(ans, intervals[i])
		} else {
			ans[l-1][1] = max(ans[l-1][1], intervals[i][1])
		}
	}

	return ans
}

package main

func subarraySum(nums []int, k int) int {
	m := map[int]int{}
	count, pre := 0, 0
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]

		if v, has := m[pre-k]; has {
			count += v
		}
		m[pre]++
	}

	return count
}

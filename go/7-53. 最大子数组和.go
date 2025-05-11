package main

func maxSubArray(nums []int) int {
	f := make([]int, len(nums))
	f[0] = nums[0]
	res := f[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(nums[i], nums[i]+f[i-1])
		res = max(res, f[i])
	}

	return res
}

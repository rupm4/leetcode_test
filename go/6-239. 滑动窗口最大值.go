package main

func maxSlidingWindow(nums []int, k int) []int {
	m := map[int]int{}
	arr := make([]int, len(nums))
	res := []int{}
	right := -1
	left := 0
	for i := 0; i < len(nums); i++ {
		if right == -1 {
			arr[right+1] = nums[i]
			m[nums[i]] = i
		} else {
			if i-m[arr[left]] >= k {
				left++
			}

			for right >= left && arr[right] <= nums[i] {
				right--
			}
			arr[right+1] = nums[i]
			m[nums[i]] = i
		}

		if i >= k-1 {
			res = append(res, arr[left])
		}
		right++
	}

	return res
}

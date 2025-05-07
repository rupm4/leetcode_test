package main

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	longestStreak := 0
	for _, num := range nums {
		numSet[num] = true
	}

	for num := range numSet {
		if numSet[num-1] {
			continue
		}

		currentStreak := 1
		currentNum := num + 1
		for numSet[currentNum] {
			currentStreak++
			currentNum++
		}

		if longestStreak < currentStreak {
			longestStreak = currentStreak
		}
	}
	return longestStreak
}

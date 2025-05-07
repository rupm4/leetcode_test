package main

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i, v := range nums {
		if p, has := hashTable[target-v]; has {
			return []int{i, p}
		}
		hashTable[v] = i
	}

	return nil
}

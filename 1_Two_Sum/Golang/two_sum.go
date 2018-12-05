func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)

	for idx := range nums {
		hashTable[nums[idx]] = idx
	}

	for idx, value := range nums {
		tmpValue := target - value
		_, ok := hashTable[tmpValue]
		if ok != false && idx != hashTable[tmpValue] {
			return []int{min(idx, hashTable[tmpValue]), max(idx, hashTable[tmpValue])}
		}
	}
	return nil
}
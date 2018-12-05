package Golang
// When submit the code on leetcode, please remove the line package Golang.

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
		if _, ok := hashTable[tmpValue]; ok && idx != hashTable[tmpValue] {
			return []int{min(idx, hashTable[tmpValue]), max(idx, hashTable[tmpValue])}
		}
	}
	return nil
}
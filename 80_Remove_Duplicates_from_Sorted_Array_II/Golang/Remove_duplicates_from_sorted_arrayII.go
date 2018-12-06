package Golang
// When submit the code on leetcode, please remove the line package Golang.

func removeDuplicates(nums []int) int {
	ret := 0
	for idx := range nums {
		if idx - 2 >= 0 {
			if nums[idx] == nums[ret-2] {
				continue
			}
		}
		nums[ret] = nums[idx]
		ret++
	}
	return ret
}
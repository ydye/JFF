package Golang
// When submit the code on leetcode, please remove the line package Golang.

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	ret := 1
	for j := 1; j < length; j++ {
		if nums[j] != nums[j-1] {
			nums[ret] = nums[j]
			ret = ret + 1
		}
	}
	return ret
}
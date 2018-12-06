package Golang
// When submit the code on leetcode, please remove the line package Golang.

func removeElement(nums []int, val int) int {
	ret := 0

	for _, valueInSlice := range nums {
		if valueInSlice != val {
			nums[ret] = valueInSlice
			ret = ret + 1
		}
	}
	return ret
}


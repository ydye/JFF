package Golang
// When submit the code on leetcode, please remove the line package Golang.

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	st, ret := 0, 0
	hashTb := make(map[rune]int)

	for idx, ch := range s {
		_, ok := hashTb[ch]
		if ok == false || hashTb[ch] < st {
			hashTb[ch] = idx
		} else {
			st = hashTb[ch] + 1
			hashTb[ch] = idx
		}
		ret = max(ret, idx - st + 1)
	}
	return ret

}
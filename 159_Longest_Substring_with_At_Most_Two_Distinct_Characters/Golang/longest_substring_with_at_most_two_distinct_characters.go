package Golang

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	hashtbCount := make(map[rune]int)
	st, ret, dist := 0, 0, 0
	for idx, value := range s {
		_, ok := hashtbCount[value]
		if ok == false {
			hashtbCount[value] = 0
		}
		hashtbCount[value] ++
		if hashtbCount[value] == 1 {
			dist++
		}
		if dist > 2 {
			for ; dist > 2 && st <= idx; st++ {
				hashtbCount[rune(s[st])]--
				if hashtbCount[rune(s[st])] == 0 {
					dist--
				}
			}

		}
		ret = max(ret, idx - st + 1)
	}
	return ret
}
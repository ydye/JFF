package Golang

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}


func maxProfit(prices []int) int {
	ret := 0
	minValue := -1
	for _, v := range prices {
		if minValue == -1 {
			minValue = v
		} else if minValue > v {
			minValue = v
		}
		ret = max(ret, v - minValue)
	}
	return ret

}
package Golang



func maxProfit(prices []int) int {
	var ret int = 0
	length := len(prices)
	if length <= 1 {
		return ret
	}
	for i := 1; i < length; i++ {
		if prices[i] > prices[i-1] {
			ret = ret + prices[i] - prices[i-1]
		}
	}
	return ret
}

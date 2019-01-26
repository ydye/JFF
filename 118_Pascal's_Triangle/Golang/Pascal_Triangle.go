package Golang

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	ret := make([][]int, 0)
	for i := 1; i <= numRows; i++ {
		line := make([]int, i)
		line[0] = 1
		line[i-1] = 1
		for j := 1; j < i - 1; j++ {
			line[j] = ret[i-2][j-1] + ret[i-2][j]
		}
		ret = append(ret, line)
	}
	return ret
}

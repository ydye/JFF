package Golang

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	flag := 0
	ret := make([][]int, 0)
	que := []*TreeNode{root}

	for len(que) != 0 {
		nxtLevel := make([]*TreeNode, 0)
		retLevel := make([]int, 0)
		for len(que) != 0 {
			nxtNd := que[0]
			que = que[1:]
			retLevel = append(retLevel, nxtNd.Val)
			if nxtNd.Left != nil {
				nxtLevel = append(nxtLevel, nxtNd.Left)
			}
			if nxtNd.Right != nil {
				nxtLevel = append(nxtLevel, nxtNd.Right)
			}
		}
		if flag == 1 {
			for i := 0; i < len(retLevel) / 2; i++ {
				opp := len(retLevel)-1-i
				retLevel[i], retLevel[opp] = retLevel[opp], retLevel[i]
			}
			flag = 0
		} else {
			flag = 1
		}
		que = nxtLevel
		ret = append(ret, retLevel)
	}
	return ret
}
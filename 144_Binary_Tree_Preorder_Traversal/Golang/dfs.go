package Golang

/**
 * Definition for a binary tree node.
 **/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Solver struct {
	ans []int
}

func (s *Solver)dfs(root *TreeNode) {
	if root == nil {
		return
	}
	s.ans = append(s.ans, root.Val)
	s.dfs(root.Left)
	s.dfs(root.Right)
}


func preorderTraversal(root *TreeNode) []int {
	ret := new(Solver)
	ret.dfs(root)
	return ret.ans
}
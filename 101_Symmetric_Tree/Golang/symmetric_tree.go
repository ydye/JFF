package Golang

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func dfs(rt1, rt2 *TreeNode) bool {
	if rt1 == nil && rt2 == nil {
		return true
	}
	if rt1 == nil || rt2 == nil {
		return false
	}
	return rt1.Val == rt2.Val && dfs(rt1.Left, rt2.Right) && dfs(rt1.Right,rt2.Left)
}

func isSymmetric(root *TreeNode) bool {
	return dfs(root, root)
}

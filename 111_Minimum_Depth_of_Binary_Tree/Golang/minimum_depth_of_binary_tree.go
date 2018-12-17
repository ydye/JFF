package Golang


/**
 * Definition for a binary tree node.
**/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftShortestPath := minDepth(root.Left)
	rightShortestPath := minDepth(root.Right)

	if leftShortestPath == 0 && rightShortestPath != 0 {
		return rightShortestPath + 1
	}

	if leftShortestPath != 0 && rightShortestPath == 0 {
		return leftShortestPath + 1
	}

	return min(leftShortestPath, rightShortestPath) + 1
}

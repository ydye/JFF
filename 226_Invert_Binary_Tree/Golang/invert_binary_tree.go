package Golang

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	leftPtr := root.Left
	rightPtr := root.Right
	root.Left = rightPtr
	root.Right = leftPtr
	return root

}
package Golang

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var l int = len(nums)
	if l == 0 {
		return nil
	}
	var mid int = (l - 1) / 2

	rt := new(TreeNode)
	rt.Val = nums[mid]
	rt.Left = sortedArrayToBST(nums[:mid])
	rt.Right = sortedArrayToBST(nums[mid + 1: ])
	return rt
}

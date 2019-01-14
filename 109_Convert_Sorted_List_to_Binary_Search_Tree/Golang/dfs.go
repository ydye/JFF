package Golang

/**
 * Definition for singly-linked list.
 **/
type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func dfs(head, tail *ListNode) *TreeNode {
	if head == tail || head.Next == tail {
		return nil
	}
	first, second := head, head
	for second != tail {
		first = first.Next
		second = second.Next
		if second != tail {
			second = second.Next
		}
	}
	rt := new(TreeNode)
	rt.Val = first.Val
	rt.Left = dfs(head, first)
	rt.Right = dfs(first, tail)
	return rt
}

func sortedListToBST(head *ListNode) *TreeNode {
	dumpy := new(ListNode)
	dumpy.Next = head
	return dfs(dumpy, nil)
}
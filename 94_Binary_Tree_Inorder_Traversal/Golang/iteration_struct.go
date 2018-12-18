package Golang

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Stack struct {
	items []*TreeNode
}

func (s *Stack) New() *Stack {
	s.items = []*TreeNode{}
	return s
}

func (s *Stack) Push(t *TreeNode) {
	s.items = append(s.items, t)
}

func (s *Stack) Pop() *TreeNode {
	item := s.items[len(s.items)-1]
	s.items = s.items[0:len(s.items)-1]
	return item
}


type Solver struct {
	st Stack
	ans []int
}


func ()

func inorderTraversal(root *TreeNode) []int {

}

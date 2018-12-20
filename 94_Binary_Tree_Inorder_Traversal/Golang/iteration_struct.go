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
	Items []*TreeNode
}

func (s *Stack) New() *Stack {
	s.Items = []*TreeNode{}
	return s
}

func (s *Stack) Push(t *TreeNode) {
	s.Items = append(s.Items, t)
}

func (s *Stack) Pop() *TreeNode {
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[0:len(s.Items)-1]
	return item
}

func (s *Stack) Empty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}


type Solver struct {
	st Stack
	ans []int
}


func (s *Solver) iteration(root *TreeNode) {
	pCur := root
	//s.st.New()
	for pCur != nil || s.st.Empty() == false {
		if pCur != nil {
			s.st.Push(pCur)
			pCur = pCur.Left
		} else {
			pCur = s.st.Pop()
			s.ans = append(s.ans, pCur.Val)
			pCur = pCur.Right
		}
	}
}

func inorderTraversal(root *TreeNode) []int {
	ret := new(Solver)
	ret.iteration(root)
	return ret.ans
}

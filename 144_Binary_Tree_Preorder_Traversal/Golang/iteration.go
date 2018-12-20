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

func (s *Stack) Pop() *TreeNode{
	top := s.Items[len(s.Items) - 1]
	s.Items = s.Items[:len(s.Items)-1]
	return top
}

func (s *Stack) Push(node *TreeNode) {
	s.Items = append(s.Items, node)
}

func (s *Stack) Empty() bool{
	return len(s.Items) == 0
}


type Solver struct {
	ans []int
	st Stack
}

func (s *Solver) iteration(root *TreeNode) {
	pCur := root
	for pCur != nil || s.st.Empty() != true {
		if pCur == nil {
			pCur = s.st.Pop()
		}
		s.ans = append(s.ans, pCur.Val)
		if pCur.Right != nil {
			s.st.Push(pCur.Right)
		}
		pCur = pCur.Left
	}
}



func preorderTraversal(root *TreeNode) []int {
	ret := new(Solver)
	ret.iteration(root)
	return ret.ans
}

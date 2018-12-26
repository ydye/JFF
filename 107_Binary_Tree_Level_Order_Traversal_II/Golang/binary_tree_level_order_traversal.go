package Golang

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Solver struct {
	ret [][]int
}

func (s *Solver) dfs(root *TreeNode, dep int) {
	if root == nil {
		return
	}
	if dep > len(s.ret) {
		s.ret = append([][]int{{}}, s.ret...)
	}
	s.ret[len(s.ret)-dep] = append(s.ret[len(s.ret)-dep], root.Val)
	s.dfs(root.Left, dep+1)
	s.dfs(root.Right, dep+1)
}


func levelOrderBottom(root *TreeNode) [][]int {
	s := new(Solver)
	s.dfs(root, 1)
	return s.ret
}
package Golang


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type solver struct {
	ans []int
}


func (p *solver) dfs(root *TreeNode) {
	if root == nil {
		return
	}
	p.dfs(root.Left)
	p.ans = append(p.ans, root.Val)
	p.dfs(root.Right)
}

func inorderTraversal(root *TreeNode) []int {
	ret := new(solver)
	ret.dfs(root)
	return ret.ans
}


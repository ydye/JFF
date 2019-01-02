package Golang

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}


type solver struct {
	res [][]int
}


func (s *solver) dfs(root *TreeNode, sum int, pathSlice []int){
	if root == nil{
		return
	}
	curPath := append(pathSlice, root.Val)
	curSum := sum - root.Val
	if root.Right == nil && root.Left == nil  {
		if curSum == 0 {
			s.res = append(s.res, append([]int{}, curPath...))
		}
		return
	}

	s.dfs(root.Left, curSum, curPath)
	s.dfs(root.Right, curSum, curPath)
}

func pathSum(root *TreeNode, sum int) [][]int {
	s := new(solver)
	s.dfs(root, sum, []int{})
	return s.res
}

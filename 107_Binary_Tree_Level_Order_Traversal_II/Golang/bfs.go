package Golang

func levelOrderBottom(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		levelData := make([]int, 0)
		nxtLevel := make([]*TreeNode, 0)
		for _, node := range queue {
			levelData = append(levelData, node.Val)
			if node.Left != nil {
				nxtLevel = append(nxtLevel, node.Left)
			}
			if node.Right != nil {
				nxtLevel = append(nxtLevel, node.Right)
			}
		}
		queue = nxtLevel
		ret = append([][]int{levelData}, ret...)
	}
	return ret
}
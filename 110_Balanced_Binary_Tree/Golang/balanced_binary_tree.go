package Golang


/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	 Val   int
	 Left  *TreeNode
	 Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func treeDepth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	l, lok := treeDepth(root.Left)
	r, rok := treeDepth(root.Right)
	cur := max(l, r) + 1

	return cur, lok && rok && abs(l - r) <= 1

}



func isBalanced(root *TreeNode) bool {
	_, ok := treeDepth(root)
	return ok
}
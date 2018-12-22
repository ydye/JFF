package Golang

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


type pair struct {
	node *TreeNode
	depth int
}

type Queue struct {
	Items []pair
	ret [][]int
}

func (q *Queue) Push(item pair) {
	q.Items = append(q.Items, item)
}

func (q *Queue) Pop() pair {
	ret := q.Items[0]
	q.Items = q.Items[1:]
	return ret
}

func (q *Queue) Empty() bool {
	return len(q.Items) == 0
}


func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	que := new(Queue)

	if root == nil {
		return ret
	}

	que.Push(pair{root, 1})
	for que.Empty() != true {
		p := que.Pop()
		if len(ret) < p.depth {
			ret = append(ret, []int{})
		}
		ret[p.depth-1] = append(ret[p.depth-1], p.node.Val)
		if p.node.Left != nil {
			que.Push(pair{p.node.Left, p.depth+1})
		}
		if p.node.Right != nil {
			que.Push(pair{p.node.Right,p.depth+1})
		}
	}

	return ret
}
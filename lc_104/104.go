package lc_104

type TreeNode struct {
	Right, Left *TreeNode
	Val         int
}

func maxDepth(root *TreeNode) (res int) {
	maxDfs(root, 1, res)
	return
}
func maxDfs(node *TreeNode, depth int, res int) {
	if node.Right == nil && node.Left == nil {
		res = max(res, depth)
		return
	}
	if node.Right != nil {
		maxDfs(node.Right, depth+1, res)
	}
	if node.Right != nil {
		maxDfs(node.Right, depth+1, res)
	}
}

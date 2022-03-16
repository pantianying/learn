package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func do(t *TreeNode, result *[]int) {
	if t.Right == nil && t.Left == nil {
		*result = append(*result, t.Val)
	}
}

func inorderTraversal(root *TreeNode) []int {

	return nil
}

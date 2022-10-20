package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func main() {
	var t = new(TreeNode)
	t.Val = 1

	t.Left = newTreeNode(2)
	t.Right = newTreeNode(2)

	t.Left.Left = newTreeNode(3)
	t.Left.Right = newTreeNode(3)

	t.Left.Left.Left = newTreeNode(4)
	t.Left.Left.Right = newTreeNode(4)
	fmt.Println(isBalanced(t))
}
func newTreeNode(v int) *TreeNode {
	return &TreeNode{Val: v}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	hl := highLevel(root.Left)
	hr := highLevel(root.Right)
	if math.Abs(float64(hl-hr)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

// 计算深度
func highLevel(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Right == nil && node.Left == nil {
		return 1
	}
	if node.Left == nil && node.Right != nil {
		return highLevel(node.Right) + 1
	}
	if node.Right == nil && node.Left != nil {
		return highLevel(node.Left) + 1
	}
	return max(highLevel(node.Left), highLevel(node.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

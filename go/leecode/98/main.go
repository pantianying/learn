package main

import "fmt"

var a *int

func main() {
	testTree := &TreeNode{
		Val: 0,
	}
	result := isValidBST(testTree)
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return trans1(root)
}

func trans1(node *TreeNode) bool {
	if node == nil {
		return true
	}

	if !trans1(node.Left) {
		return false
	}

	if a != nil {
		if node.Val <= *a {
			return false
		}
	}
	a = &node.Val

	if !trans1(node.Right) {
		return false
	}
	return true
}

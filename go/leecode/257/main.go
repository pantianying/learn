package main

import "strconv"

/*
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
叶子节点 是指没有子节点的节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res := make([]string, 0)
	dfs(root, "", &res)
	return res
}

func dfs(root *TreeNode, path string, res *[]string) {
	if path == "" && root != nil {
		path = strconv.Itoa(root.Val)
	} else if root != nil {
		path = path + "->" + strconv.Itoa(root.Val)
	}
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return
	}
	if root.Left != nil {
		dfs(root.Left, path, res)
	}
	if root.Right != nil {
		dfs(root.Right, path, res)
	}
}

package main

import "fmt"

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
示例：
二叉树：[3,9,20,null,null,15,7],
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
*/

func main() {
	treeNode := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(levelOrder(treeNode))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	res = dfs(root, 0, res)
	return res
}

func dfs(root *TreeNode, level int, res [][]int) [][]int {
	fmt.Println("level", level, "res", res)
	if root != nil {
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		res = dfs(root.Left, level+1, res)
		res = dfs(root.Right, level+1, res)
	}
	return res
}

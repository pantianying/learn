package main

import (
	"container/list"
	"fmt"
)

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

// 层序遍历
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

// 先序遍历 非递归
func PreOrderNoRecursion(root *TreeNode) []int {
	stack := list.New()
	var result []int
	t := root
	for t != nil || stack.Len() != 0 {
		for t != nil {
			result = append(result, t.Val)
			stack.PushBack(t)
			t = t.Left
		}
		for stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*TreeNode)
			t = t.Right
			stack.Remove(v)
		}
	}
	return result
}

// 中序遍历-非递归
func InOrderNoRecursion(root *TreeNode) []int {
	t := root
	stack := list.New()
	res := make([]int, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*TreeNode)
			res = append(res, t.Val) // visit
			t = t.Right
			stack.Remove(v)
		}
	}
	return res
}

// 后根遍历
func PostOrderNoRecursion(root *TreeNode) []int {
	t := root
	stack := list.New()
	res := make([]int, 0)
	var preVisited *TreeNode

	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}
		v := stack.Back()
		top := v.Value.(*TreeNode)
		if (top.Left == nil && top.Right == nil) ||
			(top.Right == nil && preVisited == top.Left) ||
			preVisited == top.Right {
			res = append(res, top.Val) //visit
			preVisited = top
			stack.Remove(v)
		} else {
			t = top.Right
		}
	}
	return res
}

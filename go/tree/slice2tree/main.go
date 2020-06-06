package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历的逆运算
func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var t *TreeNode
	var treeChan = make(chan *TreeNode, 100)
	// 取出的根节点
	var root *TreeNode
	for i := 0; i < len(data); i++ {
		if i == 0 {
			t = &TreeNode{
				Val: data[i],
			}
			root = t
			continue
		}
		// 取出要加左右子节点的根节点
		if root.Right != nil {
			root = <-treeChan
		}

		if root.Left == nil {
			root.Left = &TreeNode{
				Val: data[i],
			}
			treeChan <- root.Left
			continue
		}
		if root.Right == nil {
			root.Right = &TreeNode{
				Val: data[i],
			}
			treeChan <- root.Right
			continue
		}
	}
	bytes, _ := json.Marshal(t)
	fmt.Printf("%v", string(bytes))
}

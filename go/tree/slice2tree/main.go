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

func main() {
	//test := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 4,
	//		},
	//		Right: &TreeNode{
	//			Val: 5,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val: 6,
	//		},
	//		Right: &TreeNode{
	//			Val: 7,
	//		},
	//	},
	//}
	//levelOrder([]*TreeNode{test})
	levelOrder2()
}

// 层序遍历的逆运算
func levelOrder2() {
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

// 层序遍历
func levelOrder(queue []*TreeNode) {
	if queue == nil {
		return
	}
	for len(queue) > 0 {
		fmt.Println(queue)
		node := queue[0]
		// pop
		queue = queue[1:]
		fmt.Println(node.Val)
		if node.Left != nil {
			// push back
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			// push back
			queue = append(queue, node.Right)
		}
	}
}

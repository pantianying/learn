package main

import "fmt"

//func main() {
//	stack := list.New()
//	stack.PushBack(1)
//	stack.PushBack(2)
//	stack.PushBack(3)
//	e := stack.Back()
//	stack.Remove(e)
//
//	fmt.Println(e.Value)
//}

func testForList() {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5

	pre := reversrList(head)
	for pre != nil {
		fmt.Println(pre.Val)
		pre = pre.Next
	}
}

//链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversrList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		tmp := cur.Next //记录原来cur的next位置
		cur.Next = pre  // 当前节点的next转到新链表作为头部
		pre = cur
		cur = tmp // 计算下一个检点
	}
	return pre
}

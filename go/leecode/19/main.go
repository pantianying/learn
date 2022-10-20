package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := removeNthFromEnd(&ListNode{Val: 100, Next: &ListNode{Val: 200}}, 2)
	fmt.Println(head)
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var a = head
	var b = head
	var i = 0
	for a != nil {
		a = a.Next
		if i > n {
			b = b.Next
		}
		i++
	}
	if b.Next == nil {
		return nil
	}
	b.Next = b.Next.Next
	return head
}

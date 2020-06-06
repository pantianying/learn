package main

import (
	"container/list"
	"fmt"
)

func main() {
	stack := list.New()
	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)
	e := stack.Back()
	fmt.Println(e.Value)
	stack.
		e = stack.Back()

	fmt.Println(e.Value)
}

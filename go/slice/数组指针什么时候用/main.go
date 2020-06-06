package main

import "fmt"

// 吸取教训 ：指针传入 指针改变了就没卵用了
// 要替换指针指向的值
func main() {
	s := make([]int, 0, 10)
	s = append(s, 1)
	fmt.Printf("%v 指针1:%p\n", s, &s[0])
	test(s)
	fmt.Printf("%v 指针4:%p\n", s, s)
}
func test(s []int) {
	fmt.Printf("%v 指针2:%p\n", s, &s[0])
	s[0] = 2
	s = append(s, 1)
	fmt.Printf("%v 指针3:%p\n", s, &s[0])
}

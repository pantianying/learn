package main

import "fmt"

func main() {
	test1()
}
func test2() {
	s1 := []int{1}
	s2 := s1
	s1[0] = 0
	fmt.Println(s1, s2)
}
func test1() {
	s1 := []int{1}
	s2 := s1
	s2 = append(s2, 1)
	s2[0] = 0
	fmt.Println(s1, s2)
}

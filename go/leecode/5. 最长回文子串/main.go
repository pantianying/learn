package main

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：
输入: "cbbd"
输出: "bb"

*/

func main() {
	var s = []int{1, 2, 3}
	fmt.Println(s[1 : 2+1])
}

// func longestPalindrome(s string) string {
// 	var result []rune
// 	rs := []rune(s)
// 	leftM := make(map[rune]int)
// 	rightM := make(map[rune]int)
// 	for i := range rs {
// 		if v, ok := rightM[rs[i]]; ok {
// 			result := rs[i:v] // 这里check 一下
// 			return string(result)
// 		}
// 	}
// }

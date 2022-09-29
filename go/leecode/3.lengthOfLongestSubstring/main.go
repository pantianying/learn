package main

import "fmt"

func main() {
	result := lengthOfLongestSubstringOwn("pwwkew")
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	var max int
	for i := range s {
		l := 0
		tmp := make(map[uint8]int, 0)
		j := i
		for {
			if j >= len(s) {
				break
			}
			if _, ok := tmp[s[j]]; ok {
				break
			} else {
				tmp[s[j]] = 1
			}
			j++
			l++
		}
		if l > max {
			max = l
		}
	}
	return max
}

// 答案

func lengthOfLongestSubstring2(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

//
func lengthOfLongestSubstringOwn(s string) int {
	indexMap := make(map[rune]int, 0)
	start := 0
	maxLen := 1
	for i, v := range s {
		if index, ok := indexMap[v]; ok && index >= start {
			// 计算 start 到 i-1 的长度
			l := i - start
			if l > maxLen {
				maxLen = l
			}
			start = index + 1
			indexMap[v] = i
		} else {
			indexMap[v] = i
		}
	}
	return maxLen
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	helper(root, 0, &result)
	return result
}

func helper(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}
	if len(*result) == level {
		*result = append(*result, make([]int, 0))
	}
	(*result)[level] = append((*result)[level], node.Val)
	helper(node.Left, level+1, result)
	helper(node.Right, level+1, result)
	return
}

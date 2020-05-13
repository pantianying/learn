package main

func main() {
	lengthOfLongestSubstring("abcabcbb")
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

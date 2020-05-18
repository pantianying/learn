package main

// 超出时间限制 需要cache

func main() {
	isMatch("aa", "*")
}

func isMatch(s string, p string) bool {
	p = removeDuplicateStars(p)
	return helper(s, p)
}

func removeDuplicateStars(p string) string {
	bp := []byte(p)
	var result []byte
	var before = false
	for i := range bp {
		// 前一个是*，后面的是*的就不要了
		if before && bp[i] == '*' {
			continue
		}
		if bp[i] == '*' {
			before = true
		} else {
			before = false
		}
		result = append(result, bp[i])
	}
	return string(result)
}
func helper(s string, p string) bool {
	if s == p || p == "*" {
		return true
	}
	if s == "" || p == "" {
		return false
	}
	if p[0] == '*' {
		return helper(s, p[1:]) || helper(s[1:], p)
	}

	if s[0] == p[0] || p[0] == '?' {
		return helper(s[1:], p[1:])
	}

	return false
}

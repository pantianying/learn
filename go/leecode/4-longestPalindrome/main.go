package main

import "fmt"

// 回文数
func main() {
	result := longestPalindrome("aababc")
	fmt.Println(result)
}

//
// resultSet[i][j]=resultSet[i+1][j-1] && s[i]==s[j]
func longestPalindrome(s string) string {
	l := len(s)

	result := s[:1]

	resultSet := make([][]bool, l)
	// int resultSet
	for i := range resultSet {
		resultSet[i] = make([]bool, l)
	}
	for i := 0; i < l; i++ {
		resultSet[i][i] = true
	}
	for j := 0; j <= l-1; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				resultSet[i][j] = false
			} else {
				if j-i < 3 {
					resultSet[i][j] = true
					palindromeStr := s[i : j+1]
					if len(palindromeStr) > len(result) {
						result = palindromeStr
					}
					continue
				}
				if resultSet[i+1][j-1] {
					resultSet[i][j] = true
					palindromeStr := s[i : j+1]
					if len(palindromeStr) > len(result) {
						result = palindromeStr
					}
				} else {
					resultSet[i][j] = false
				}
			}
		}
	}
	return result
}

func palindrome(s string) bool {
	for i := 0; i < len(s)/2+1; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// 中分法
func longestPalindrome2(s string) string {
	l := len(s)
	longestPalindrome := s[:1]
	for i := 0; i < l; i++ {
		var palindrome string
		for j := 1; i-j >= 0 && j+i <= l-1; j++ {
			// 判断单数情况
			if i+j == l-1 {
				palindrome = s[i-j:]
				break
			}
			if i-j == 0 {
				palindrome = s[:i+j+1]
				break
			}
			if s[i-j] != s[i+j] {
				palindrome = s[i-j+1 : i+j]
				break
			} else {
				continue
			}
		}
		if len(longestPalindrome) < len(palindrome) {
			longestPalindrome = palindrome
		}

		for j := 0; i-j+1 > 0 && j+i+1 <= l-1; j++ {
			if i+j+1 == l-1 {
				palindrome = s[i-j:]
				break
			}
			if i-j == 0 {
				palindrome = s[:i+j+1]
				break
			}

			if s[i-j] != s[i+j+1] {
				palindrome = s[i-j+1 : i+j+1]
				break
			}
		}
		if len(longestPalindrome) < len(palindrome) {
			longestPalindrome = palindrome
		}
	}
	return longestPalindrome
}

package main

import "fmt"

func main() {
	result := generateParenthesis(10)
	fmt.Println(result)
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	dfs("", &result, n, 0, 0)
	return result
}

func dfs(s string, result *[]string, n int, leftNum, rightNum int) {
	if leftNum < rightNum {
		return
	}
	if len(s) == 2*n {
		*result = append(*result, s)
	}
	dfs(s+"(", result, n, leftNum+1, rightNum)
	dfs(s+")", result, n, leftNum, rightNum+1)
	return
}

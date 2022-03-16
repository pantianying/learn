package main

import "fmt"

func main() {
	a := permuteUnique([]int{0, 1, 1, 4})
	fmt.Println(a)
}

// 打印全排列
// 深度遍历
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	dfs(nums, path, used, &res)
	return res
}

func dfs(nums, path []int, used []bool, res *[][]int) {
	if len(path) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i, k := range used {
		if k == false {
			path = append(path, nums[i])
			used[i] = true
			dfs(nums, path, used, res)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}

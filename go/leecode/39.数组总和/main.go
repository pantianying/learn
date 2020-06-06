package main

import (
	"sort"
)

/*

给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取
*/

func main() {
	candidates := []int{2, 5, 3}
	combinationSum(candidates, 8)
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(x, y int) bool {
		return candidates[x] < candidates[y]
	})
	res := [][]int{}
	dfs(candidates, nil, target, 0, &res)
	return res
}

func dfs(candidates, nums []int, target, left int, res *[][]int) {
	if target == 0 { // 结算
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := left; i < len(candidates); i++ { // left限定，形成分支
		if target < candidates[i] { // 剪枝
			return
		}
		t := append(nums, candidates[i])
		dfs(candidates, t, target-candidates[i], i, res) // 分支
	}
}

package main

// 打印全排列，不重复
// 深度遍历
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	dfsUnique(nums, path, used, &res)
	return res
}

func dfsUnique(nums, path []int, usedIndex []bool, res *[][]int) {
	if len(path) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	usedValue := make(map[int]bool)
	for i, k := range usedIndex {
		if k == false {
			if ok, _ := usedValue[nums[i]]; ok {
				continue
			}
			usedValue[nums[i]] = true
			path = append(path, nums[i])
			usedIndex[i] = true
			dfsUnique(nums, path, usedIndex, res)
			path = path[:len(path)-1]
			usedIndex[i] = false
		}
	}
}

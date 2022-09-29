package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

/*

给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。



示例 1：
输入：nums = [10,2]
输出："210"
示例 2：

输入：nums = [3,30,34,5,9]
输出："9534330"

提示：
1 <= nums.length <= 100
0 <= nums[i] <= 109
*/

type T struct {
	Nums []int
}

func (t T) Len() int {
	return len(t.Nums)
}

func (t T) Less(i, j int) bool {
	a := t.Nums[i]
	b := t.Nums[j]
	return getWeight(a) > getWeight(b)
}

func (t T) Swap(i, j int) {
	t.Nums[i], t.Nums[j] = t.Nums[j], t.Nums[i]
}

// 获得一个数的排序权重
func getWeight(num int) int {
	return (num+1)*(int(math.Pow10(9-len(strconv.Itoa(num))))) - 1
}

func largestNumber(nums []int) string {
	t := &T{
		Nums: nums,
	}
	sort.Sort(t)
	var result string
	for _, v := range t.Nums {
		result += strconv.Itoa(v)
	}
	return result
}

func main() {
	// 什么
	s := largestNumber([]int{3, 30, 34, 5, 9})
	fmt.Println(s)
}

package main

func main() {

}

func subsets(nums []int) [][]int {
	l := len(nums)
	if l == 1 {
	}
	result := make([][]int, 0)

	subsets(nums[:l])
}

func helper(nums []int, res *[][]int) {

}

package main

func main() {
	req := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	spiralOrder(req)
}

// 理解有问题
func spiralOrder(matrix [][]int) []int {
	side := len(matrix)
	res := make([]int, 4*side-4)

	for i := 0; i < side; i++ {
		res[i] = matrix[0][i]
	}
	for i := 1; i < side; i++ {
		res[side-1+i] = matrix[i][side-1]
	}
	for i := 1; i < side; i++ {
		res[2*side+i-2] = matrix[side-1][side-i-1]
	}
	for i := 1; i < side-1; i++ {
		res[3*side+i-3] = matrix[side-i-1][0]
	}
	return res
}

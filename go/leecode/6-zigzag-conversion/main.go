package main

/**
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：
string convert(string s, int numRows);
*/

func main() {
	convert("ABC", 1)
}
func convert(s string, numRows int) string {

	// 为numRow:1 的情况单独处理
	if numRows == 1 {
		return s
	}
	z := make([][]rune, numRows)
	ss := []rune(s)

	// 可以计算出该Z图的宽度，因为填满平均一列用两个rune
	// 以此来初始化整个Z图
	width := len(ss)/2 + 1
	for weight := range z {
		z[weight] = make([]rune, width)
	}

	var i, j = 0, 0
	for k := range ss {
		z[j][i] = ss[k]
		if k == len(ss)-1 {
			break
		}
		// 向下移动
		if j < numRows-1 && i%(numRows-1) == 0 {
			// 注意 这里j变动了 后续不能用j进行比较，直接用else最为简单
			j++
		} else {
			j--
			i++
		}
	}
	// 后续直接安层遍历即可
	var result []rune
	for j := range z {
		for i := range z[j] {
			v := z[j][i]
			if v != 0 {
				result = append(result, z[j][i])
			}
		}
	}
	return string(result)
}

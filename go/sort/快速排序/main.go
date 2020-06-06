package main

import (
	"fmt"
)

func main() {
	a := []int{314, 1341, 32, 2134, 52345, 7, 423521, 134, 1341, 2341, 41341234, 12341, 23, 6}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func QuickSort(data []int, low, high int) {
	tmpLow := low
	tmpHigh := high

	if tmpLow >= tmpHigh {
		return
	}
	tmp := data[tmpLow]
	for tmpLow < tmpHigh {
		for data[tmpHigh] >= tmp && tmpLow < tmpHigh {
			tmpHigh--
		}
		data[tmpLow] = data[tmpHigh]
		for data[tmpLow] <= tmp && tmpLow < tmpHigh {
			tmpLow++
		}
		data[tmpHigh] = data[tmpLow]
	}
	data[tmpLow] = tmp
	fmt.Println(data)
	QuickSort(data, low, tmpLow-1)
	QuickSort(data, tmpLow+1, high)
	return
}

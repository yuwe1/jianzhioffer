package main

import "fmt"

// 顺时针打印矩阵
func PrintMatrixClockwisely(data [][]int) {
	// 定义四个边界
	// 右边界
	right := len(data[0]) - 2
	// 底部边界
	buttom := len(data) - 2
	// 左边界
	left := 0
	// 顶边界
	top := 0

	// 定义行列索引
	row := 0
	clo := 0

	for left < len(data) {
		for row == left && clo <= right {
			fmt.Print(data[row][clo], " ")
			clo++
		}
		for clo == right+1 && row <= buttom {
			fmt.Print(data[row][clo], " ")
			row++
		}
		for row == buttom+1 && clo >= left+1 {
			fmt.Print(data[row][clo], " ")
			clo--
		}
		for clo == left && row >= top+1 {
			fmt.Print(data[row][clo], " ")
			row--
		}
		row++
		clo++
		right--
		top++
		buttom--
		left++
	}
}
func main() {
	data := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	PrintMatrixClockwisely(data)
}

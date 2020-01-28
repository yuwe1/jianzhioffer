### 顺时针打印矩阵
1  2  3  4
5  6  7  8
9  10 11 12
13 14 15 16

### 思路
***定义四个边界***
打印顺序：
- 1,2,3  右边界至3
- 4,8 12 底部边界至12
- 16 15 14 左边界至14
- 13 9 5 上边界至5
每次经过一个循环，各个边界进行自增或自减操作

```go
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
```
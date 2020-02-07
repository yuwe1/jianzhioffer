package main

import "fmt"

/*
输入数组：

[
  [1,2,8,9]，
  [2,4,9,12]，
  [4,7,10,13]，
  [6,8,11,15]
]

如果输入查找数值为7，则返回true，

如果输入查找数值为5，则返回false。
*/
func searchArray(data [][]int, target int) bool {
	if len(data) <=0{
		return false
	}
	i := 0
	j := len(data[0]) - 1
	for i < len(data) 0 && j >= 0 {
		if data[i][j] == target {
			return true
		} else if data[i][j] > target {
			j--
		} else if data[i][j] < target {
			i++
		}
	}
	return false
}

func main() {
	data := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	fmt.Println(searchArray(data, 7))
}

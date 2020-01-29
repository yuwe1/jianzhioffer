package main

import "fmt"

// 判断一个数组是否可能是一个二叉搜索树的一个后续遍历
func IsAfter(data []int) bool {

	root := data[len(data)-1]
	i := 0
	for i = 0; i < len(data)-2; i++ {
		if data[i] > root {
			break
		}
	}
	j := i
	// 判断右边的数组是i否存在比根节点小的值
	for i := j; i < len(data)-2; i++ {
		if data[i] < root {
			return false
		}
	}
	right := true
	// 继续判断右边的数组
	if i > 0 {
		right = IsAfter(data[:j])
	}
	left := true
	if j < len(data)-1 {
		left = IsAfter(data[j : len(data)-1])
	}
	return left && right
}

func main() {
	data := []int{5, 7, 6, 9, 11, 10}
	fmt.Println(IsAfter(data))
}

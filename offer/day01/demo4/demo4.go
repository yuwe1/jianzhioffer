package main

import (
	"fmt"
)

// 二叉查找树
type biraryTree struct {
	value int
	left  *biraryTree
	right *biraryTree
}

// 构建一个二叉查找树
func createBiraryTree(root *biraryTree, value int) {
	if root.value > value {
		if root.left == nil {
			root.left = &biraryTree{value, nil, nil}
		} else {
			createBiraryTree(root.left, value)
		}
	} else {
		if root.right == nil {
			root.right = &biraryTree{value, nil, nil}
		} else {
			createBiraryTree(root.right, value)
		}
	}
}

// 中序遍历
func middleRead(root *biraryTree) {
	if root == nil {
		return
	}
	middleRead(root.left)
	fmt.Print(root.value, " ")
	middleRead(root.right)
}
func main() {
	data := []int{2, 3, 1, 5, 9, 4, 0, 10}
	root := &biraryTree{data[0], nil, nil}
	for i := 1; i < len(data); i++ {
		createBiraryTree(root, data[i])
	}
	middleRead(root)
}

package main

import (
	"fmt"
)

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

// 创建二叉树
func createTree(data []int, index int) (root *treeNode) {
	if index < len(data) {
		root = &treeNode{data[index], nil, nil}
		root.left = createTree(data, 2*index+1)
		root.right = createTree(data, 2*index+2)
	}
	return
}

// 层序遍历二叉树
func orderRead(root *treeNode) {
	if root == nil {
		return
	}
	data := []*treeNode{root}
	for len(data) > 0 {
		p := data[0]
		fmt.Println(data[0].value)
		if p.left != nil {
			data = append(data, p.left)
		}
		if p.right != nil {
			data = append(data, p.right)
		}
		data = data[1:]
	}
}

var levels [][]int
var data []int

//第二种方法
func helper(root *treeNode, level int) {
	if level == len(levels) {
		levels = append(levels, []int{})
	}
	levels[level] = append(levels[level], root.value)
	if root.left != nil {
		helper(root.left, level+1)
	}
	if root.right != nil {
		helper(root.right, level+1)
	}
}
func order_two(root *treeNode) [][]int {
	if root == nil {
		return levels
	}
	helper(root, 0)
	return levels
}
func preRead(root *treeNode) {
	if root != nil {
		fmt.Println(root.value)
		preRead(root.left)
		preRead(root.right)
	}
}
func main() {
	data := []int{2, 3, 1, 4, 5, 6, 3}
	root := createTree(data, 0)
	preRead(root)
	fmt.Println("***********")
	fmt.Println(order_two(root))
}

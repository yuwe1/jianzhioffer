package main

import "fmt"

// Z字形层次遍历二叉树

// 定义一个treeNode
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// 创建一个二叉树

func createTree(data []int, index int) (t *treeNode) {
	if index >= len(data) {
		return
	}
	t = &treeNode{data[index], nil, nil}
	t.left = createTree(data, 2*index+1)
	t.right = createTree(data, 2*index+2)
	return t
}

var levels [][]int
var level []int

// 递归
func LevelRead(root *treeNode) [][]int {
	if root == nil {
		return levels
	}
	l := 0
	helper(root, l)
	return levels
}

func helper(node *treeNode, lev int) {
	if lev == len(levels) {
		levels = append(levels, []int{})
	}
	levels[lev] = append(levels[lev], node.val)
	if node.left != nil {
		helper(node.left, lev+1)
	}
	if node.right != nil {
		helper(node.right, lev+1)
	}
}

// 使用数组
func LevelRead_two(root *treeNode) {
	data := make([]*treeNode, 1)
	data[0] = root
	for len(data) > 0 {
		fmt.Println(data[0].val)
		if data[0].left != nil {
			data = append(data, data[0].left)
		}
		if data[0].right != nil {
			data = append(data, data[0].right)
		}
		data = data[1:]
	}
}
func main() {
	data := []int{3, 4, 5, 1, 2, 4, 6}
	root := createTree(data, 0)
	// result := LevelRead(root)
	// fmt.Println(result)
	LevelRead_two(root)
}

package main

import "fmt"

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

// 构建一个二叉树
func createtree(data []int, index int) (t *treeNode) {
	if index >= len(data) {
		return
	}
	t = &treeNode{data[index], nil, nil}
	t.left = createtree(data, 2*index+1)
	t.right = createtree(data, 2*index+2)
	return t
}

// 前序遍历二叉树
func preTree(root *treeNode) {
	if root == nil {
		return
	}
	p := root
	fmt.Println(p.value)
	preTree(p.left)
	preTree(p.right)
}

// 从下往上打印二叉树
func printFromtop(root *treeNode) {

}
func main() {
	data := []int{8, 6, 10, 5, 7, 9, 11}
	root := createtree(data, 0)
	printFromtop(root)
}

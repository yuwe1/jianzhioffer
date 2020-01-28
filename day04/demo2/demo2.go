package main

import "fmt"

// 二叉树的镜像
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func createTree(val []int, i int) (root *treeNode) {
	if i >= len(val) {
		return
	}
	root = &treeNode{val[i], nil, nil}
	root.left = createTree(val, 2*i+1)
	root.right = createTree(val, 2*i+2)
	return root
}

// 遍历树
func readTree(root *treeNode) {
	if root == nil {
		return
	}
	p := root
	fmt.Print(p.val, " ")
	readTree(p.left)
	readTree(p.right)
}

//两颗互为镜像的二叉树
func MicrrorRecusively(root *treeNode) {
	if root == nil || (root.left == nil && root.right == nil) {
		return
	}
	root.left, root.right = root.right, root.left
	if root.left != nil {
		MicrrorRecusively(root.left)
	}
	if root.right != nil {
		MicrrorRecusively(root.right)
	}
}
func main() {
	val := []int{8, 6, 10, 5, 7, 9, 11}
	root := createTree(val, 0)
	readTree(root)
	MicrrorRecusively(root)
	fmt.Println()
	readTree(root)
}

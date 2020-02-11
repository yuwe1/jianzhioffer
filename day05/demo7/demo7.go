package main

import "fmt"

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// 构建一个二叉搜索树
func createtree(root *treeNode, value int) {
	if root.val > value {
		if root.left == nil {
			root.left = &treeNode{value, nil, nil}
		} else {
			createtree(root.left, value)
		}
	} else {
		if root.right == nil {
			root.right = &treeNode{value, nil, nil}
		} else {
			createtree(root.right, value)
		}
	}
}
func convent(root *treeNode, last *treeNode) {
	if root == nil {
		return
	}
	convent(root.left, last)
	root.left = last
	if last != nil {
		last.right = root
	}
	last = root
	convent(root.right, last)
}

// 转换成双链表
func ToLink(root *treeNode) *treeNode {
	if root == nil {
		return root
	}
	//
	var last *treeNode = nil
	convent(root, last)
	for root.left != nil {
		root = root.left
	}
	return root
}

// 中序遍历
func midRead(root *treeNode) {
	if root != nil {
		midRead(root.left)
		fmt.Print(root.val, " ")
		midRead(root.right)
	}
}

// 链表的遍历
func LinkRead(root *treeNode) {
	temp := root
	for temp.right != root {
		fmt.Println(temp.val)
		temp = temp.right
	}
}
func main() {
	data := []int{10, 6, 14, 4, 8, 12, 16}
	root := &treeNode{data[0], nil, nil}
	for i := 1; i < len(data); i++ {
		createtree(root, data[i])
	}
	r := ToLink(root)
	LinkRead(r)
}

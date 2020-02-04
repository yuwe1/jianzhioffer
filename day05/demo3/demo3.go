package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 创建一个二叉搜索树的类
type Tree struct {
	root *TreeNode
}

// 创建一个二叉搜索树
func (t *Tree) createTree(val int) {
	if t.root == nil {
		temp := &TreeNode{val, nil, nil}
		t.root = temp
		// fmt.Println(t.root.Val)
		return
	}
	p := t.root
	for p != nil {
		if p.Val > val && p.Left == nil {
			p.Left = &TreeNode{val, nil, nil}
		} else if p.Val > val && p.Left != nil {
			p = p.Left
		} else if p.Val < val && p.Right == nil {
			p.Right = &TreeNode{val, nil, nil}
		} else {
			p = p.Right
		}
	}
}

// 遍历二叉搜索树
func read(root *TreeNode) {
	if root == nil {
		// fmt.Println()
		return
	}

	read(root.Left)
	fmt.Print(root.Val, " ")
	read(root.Right)
}

// 求二叉树的高度
func GetHight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := GetHight(root.Left)
	right := GetHight(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return isBST(root, -1<<63, 1<<63-1)
}
func isBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	val := root.Val
	if min >= val || max <= val {
		return false
	}
	return isBST(root.Left, min, val) && isBST(root.Right, val, max)
}
func main() {
	t := &Tree{nil}
	data := []int{3, 1, 4, 6, 8, 9, 0, 2}
	for i := 0; i < len(data); i++ {
		t.createTree(data[i])
	}
	read(t.root)
	fmt.Println()
	fmt.Println(GetHight(t.root))
	fmt.Println(isValidBST(t.root))
}

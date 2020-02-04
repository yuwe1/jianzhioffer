package main

import "fmt"

// 二叉树的层次遍历
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

var levels [][]int

func helper(treenode *TreeNode, level int) {
	if len(levels) == level {
		levels = append(levels, []int{})
	}
	levels[level] = append(levels[level], treenode.Val)
	if treenode.Left != nil {
		helper(treenode.Left, level+1)
	}
	if treenode.Right != nil {
		helper(treenode.Right, level+1)
	}
}
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return levels
	}
	helper(root, 0)
	return levels
}

//
func main() {
	t := &Tree{nil}
	data := []int{3, 1, 4, 6, 8, 9, 0, 2}
	for i := 0; i < len(data); i++ {
		t.createTree(data[i])
	}
	read(t.root)
	fmt.Println(levelOrder(t.root))
}

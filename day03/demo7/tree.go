package main

import "fmt"

// 定义一个数的结点类
type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}
type Tree interface {
	// 动态创建一个二叉树（默认为二叉查找树）
	CreateTree(val interface{})
}

// 定义一个二叉树类
type binaryTree struct {
	// 根结点
	root *treeNode
}

func NewBinaryTree() *binaryTree {
	return &binaryTree{
		root: nil,
	}
}

// 获取根节点
func (t *binaryTree) Root() *treeNode {
	return t.root
}

func (t *binaryTree) SetRoot(val treeNode) {
	t.root = &val
}

func (t *binaryTree) CreateTree(val int) {
	if t.Root() == nil {
		t.SetRoot(treeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		})
	} else {
		root := t.Root()
		for root != nil {
			if val <= root.Val && root.Left == nil{
				root.Left = &treeNode{val,nil,nil}
				break
			} else if val > root.Val && root.Right == nil {
				root.Right = &treeNode{val,nil,nil}
				break
			} else if val <= root.Val && root.Left != nil {
				root = root.Left
			} else if val > root.Val && root.Right != nil {
				root = root.Right
			}
		}
	}
}

// 前序遍历
func PreTraverseBTree(root *treeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	PreTraverseBTree(root.Left)
	PreTraverseBTree(root.Right)
}

// 中序遍历
func MidTraverseTree(root *treeNode) {
	if root == nil {
		return
	}
	PreTraverseBTree(root.Left)
	fmt.Print(root.Val, " ")
	PreTraverseBTree(root.Right)
}

// 后序遍历
func PostTraverseTree(root *treeNode) {
	if root == nil {
		return
	}
	PreTraverseBTree(root.Left)
	PreTraverseBTree(root.Right)
	fmt.Print(root.Val, " ")
}


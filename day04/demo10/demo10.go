package main

import "fmt"

// 定义二叉搜索树的结点
type BiraryTreeNode struct {
	value int
	left  *BiraryTreeNode
	right *BiraryTreeNode
}

// 定义一个二叉树的类
type Tree struct {
	root *BiraryTreeNode
}

//创建一个二叉搜索树
func (t *Tree) createTree(value int) {
	if t.root == nil {
		temp := &BiraryTreeNode{value, nil, nil}
		t.root = temp
	} else {
		p := t.root
		for p != nil {
			if value <= p.value && p.left != nil {
				p = p.left
			} else if value <= p.value && p.left == nil {
				p.left = &BiraryTreeNode{value, nil, nil}
				break
			} else if value > p.value && p.right != nil {
				p = p.right
			} else if value > p.value && p.right == nil {
				p.right = &BiraryTreeNode{value, nil, nil}
				break
			}
		}
	}
}

// 中遍历二叉树
func preRead(root *BiraryTreeNode) {
	if root == nil {
		return
	}
	preRead(root.left)
	fmt.Print(root.value, " ")
	preRead(root.right)
}

func main() {
	data := []int{10, 6, 4, 8, 14, 12, 16}
	t := new(Tree)
	for _, v := range data {
		t.createTree(v)
		// fmt.Println(t.root)
	}
	preRead(t.root)
}

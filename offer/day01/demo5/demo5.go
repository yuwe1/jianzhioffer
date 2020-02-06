package main

import (
	"fmt"
)

// 红黑树
type RedBLackTree struct {
	value  int
	color  int
	left   *RedBLackTree
	right  *RedBLackTree
	parent *RedBLackTree
}

var (
	redroot *RedBLackTree
)

const (
	red   = 1
	black = 0
)

// 创建根节点
func NewRoot(value int) *RedBLackTree {
	return &RedBLackTree{
		value:  value,
		left:   nil,
		right:  nil,
		color:  black,
		parent: nil,
	}
}

func Add(root *RedBLackTree, value int) {
	if root.value > value {
		if root.left == nil {
			root.left = &RedBLackTree{value, red, nil, nil, root}
			AddFix(&redroot, root.left)
		} else {
			Add(root.left, value)
		}
	} else {
		if root.right == nil {
			root.right = &RedBLackTree{value, red, nil, nil, root}
			AddFix(&redroot, root.right)
		} else {
			Add(root.right, value)
		}
	}
}

//
func AddFix(root **RedBLackTree, node *RedBLackTree) {
	// 获得祖父结点
	for node.parent != nil && node.parent.color == red {
		parent := node.parent
		// 获得祖父结点
		gparent := parent.parent
		if parent == gparent.left {
			// 获得叔叔结点
			uncle := gparent.right
			// 如果叔叔的结点为红色,使用变色
			if uncle != nil && uncle.color == red {
				parent.color = black
				uncle.color = black
				gparent.color = red
				node = gparent
				continue
			}

			// 叔叔结点是黑色，且当前结点时右结点
			if node == parent.right {
				LeftRolate(root, parent)
				parent, node = node, parent
			}

			// 叔叔的结点是黑色
			parent.color = black
			gparent.color = red
			RightRolate(root, gparent)
		} else { //父节点是右子节点
			uncle := gparent.left
			// 叔叔的结点时红色
			if uncle != nil && uncle.color == red {
				parent.color = black
				uncle.color = black
				gparent.color = red
				node = gparent
				continue
			}

			// 叔叔的结点是黑色
			if node == parent.left {
				RightRolate(root, parent)
				parent, node = node, parent
			}

			// 叔叔的结点时黑色，且为右结点
			parent.color = black
			gparent.color = red
			LeftRolate(root, gparent)
		}

	}
	*root.color = black
}

// 中序遍历
func MiddleRead(root *RedBLackTree) {
	if root != nil {
		MiddleRead(root.left)
		fmt.Println(root.value)
		MiddleRead(root.right)
	}
}

// 左旋
func LeftRolate(root **RedBLackTree, node *RedBLackTree) {
	if node != nil {
		// 保留当前节点的右节点
		r := node.right
		node.right = r.left
		if r.left != nil {
			r.left.parent = r
		}
		// 将
		r.parent = node.parent
		if node.parent == nil {
			*root = r
		} else {
			if node == node.parent.left {
				node.parent.left = r
			} else {
				node.parent.right = r
			}
		}
		// 将r的左子节点设为x，将x的父节点设置为r
		r.left = node
		node.parent = r
	}
}

// 右旋
func RightRolate(root **RedBLackTree, node *RedBLackTree) {
	if node != nil {
		// 保存当前结点的左结点
		l := node.left
		node.left = l.right
		if l.right != nil {
			l.right.parent = node
		}
		l.parent = node.parent
		if node.parent == nil {
			*root = l
		} else {
			if node.parent.left == node {
				node.parent.left = l
			} else {
				node.parent.right = l
			}
		}
		l.right = node
		node.parent = l
	}
}

// 前序遍历
func preRead(root *RedBLackTree) {
	if root != nil {
		fmt.Println(root.value)
		preRead(root.left)
		preRead(root.right)
	}
}

func main() {
	data := []int{9, 8, 12, 10, 14}
	redroot = NewRoot(data[0])
	for i := 1; i < len(data); i++ {
		Add(redroot, data[i])
	}
	preRead(redroot)

	preRead(redroot)
}

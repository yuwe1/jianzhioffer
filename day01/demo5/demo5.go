package main

import (
	"fmt"
)

// 定义存储的对象
type ojbect interface{}

// 定义一个二叉树的节点
type Node struct {
	data  ojbect
	left  *Node
	right *Node
}

// 定义一个tree
type TreeNode struct {
	// 树的深度
	h int
	// 根节点
	root *Node
}

func contruct(pre []int, mid []int) (t *Node) {
	l := len(pre)
	if l == 0 {
		return nil
	}
	root := &Node{
		data: pre[0],
	}
	if l == 1 {
		return root
	}
	leftlen, rightlen := 0, 0
	// 遍历我们的中序遍历结果，找到位置分开左子树和右子树
	for _, v := range mid {
		if v == pre[0] {
			break
		}
		leftlen++
	}
	rightlen = l - leftlen - 1
	// 构建我们的左子树
	if leftlen > 0 {
		root.left = contruct(pre[1:leftlen+1], mid[0:leftlen])
	}
	if rightlen > 0 {
		root.right = contruct(pre[leftlen+1:], mid[leftlen+1:])
	}
	return root
}

// 二叉树的前序遍历
func preorder(root *Node) {
	if root == nil {

		return
	}
	fmt.Print(root.data)
	preorder(root.left)
	preorder(root.right)
}

//二叉树的中序遍历
func midorder(root *Node) {
	if root == nil {
		return
	}
	midorder(root.left)
	fmt.Print(root.data)
	midorder(root.right)
}

// 二叉树的后序遍历
func afterorder(root *Node) {
	if root == nil {
		return
	}
	afterorder(root.left)
	afterorder(root.right)
	fmt.Print(root.data)
}
func main() {
	per := []int{1, 2, 4, 7, 3, 5, 6, 8}
	mid := []int{4, 7, 2, 1, 5, 3, 8, 6}
	root := contruct(per, mid)
	preorder(root)
	fmt.Println()
	midorder(root)
	fmt.Println()
	afterorder(root)
	fmt.Println()

	// 测试用例完全二叉树
	preall := []int{4, 5, 1, 7, 8, 9, 6, 2, 3}
	midall := []int{7, 1, 8, 5, 9, 4, 2, 6, 3}
	rootall := contruct(preall, midall)
	afterorder(rootall)
}

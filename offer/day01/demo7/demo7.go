package main

import "fmt"

type treeNode struct {
	value string
	left  *treeNode
	right *treeNode
}

// 给出二叉树的先序和中序，求后序
func getPostRead(pre string, mid string) {
	if len(pre) > 0 {
		root := pre[0]
		// 找到左子树和右子树
		i := 0
		for i = 0; i < len(mid); i++ {
			if mid[i] == root {
				break
			}
		}
		getPostRead(pre[1:i+1], mid[0:i])
		getPostRead(pre[i+1:], mid[i+1:])
		fmt.Println(string(root))
	}
}

// 根据先序和中序构建二叉树
func createTree(pre string, mid string) (root *treeNode) {
	if len(pre) <= 0 {
		return root
	}
	root = &treeNode{string(pre[0]), nil, nil}
	i := 0
	for i = 0; i < len(mid); i++ {
		if string(mid[i]) == root.value {
			break
		}
	}
	root.left = createTree(pre[1:i+1], mid[0:i])
	root.right = createTree(pre[i+1:], mid[i+1:])
	fmt.Println(root.value)
	return root
}

// 二叉树的先序遍历
func preRead(root *treeNode) {
	if root != nil {
		fmt.Print(root.value)
		preRead(root.left)
		preRead(root.right)
	}
}
func main() {
	pre := "GDAFEMHZ"
	mid := "ADEFGHMZ"
	// getPostRead(pre, mid)
	root := createTree(pre, mid)
	preRead(root)
}

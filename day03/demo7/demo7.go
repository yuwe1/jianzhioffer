package main

import "fmt"

// 判断B是否是A的子结构
func HasSubtree(pRoot1 *treeNode, pRoot2 *treeNode) bool {
	result := false
	if pRoot1 != nil && pRoot2 != nil {
		if pRoot1.Val == pRoot2.Val {
			result = DoesTreeHHaveTree2(pRoot1, pRoot2)
		}
		if !result {
			result = HasSubtree(pRoot1.Left, pRoot2)
		}
		if !result {
			result = HasSubtree(pRoot1.Right, pRoot2)
		}
	}
	return result
}
func DoesTreeHHaveTree2(pRoot1 *treeNode, pRoot2 *treeNode) bool {
	if pRoot2 == nil {
		return true
	}
	if pRoot1 == nil {
		return false
	}
	if pRoot1.Val != pRoot2.Val {
		return false
	}
	return DoesTreeHHaveTree2(pRoot1.Left, pRoot2.Left) && DoesTreeHHaveTree2(pRoot1.Right, pRoot2.Right)
}
func main() {
	root := NewBinaryTree()
	root.CreateTree(6)
	root.CreateTree(4)
	root.CreateTree(3)
	root.CreateTree(5)
	root.CreateTree(3)
	root.CreateTree(4)
	root.CreateTree(10)
	root.CreateTree(7)
	root.CreateTree(8)

	root2 := NewBinaryTree()
	root2.CreateTree(4)
	root2.CreateTree(3)
	root2.CreateTree(3)
	root2.CreateTree(4)
	root2.CreateTree(4)

	fmt.Println(HasSubtree(root.Root(), root2.Root()))
}

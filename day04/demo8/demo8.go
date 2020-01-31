package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func createTree(data []int, index int) (root *TreeNode) {
	if index >= len(data) {
		return
	}
	root = &TreeNode{data[index], nil, nil}
	root.left = createTree(data, 2*index+1)
	root.right = createTree(data, 2*index+2)
	return root
}

//
func FindPath(treeNode *TreeNode, expect int, path []int, current int) {
	current += treeNode.value
	path = append(path, treeNode.value)

	if treeNode.left == nil && treeNode.right == nil && expect == current {
		// 打印路径
		fmt.Println(path)
	}

	// 如果不是继续遍历子节点
	if treeNode.left != nil {
		FindPath(treeNode.left, expect, path, current)
	}
	if treeNode.right != nil {
		FindPath(treeNode.right, expect, path, current)
	}
	// 返回父节点之前需要将当前的值进行删除
	current = current - treeNode.value
	path = path[:len(path)-1]
}
func findPath(t *TreeNode, expectnum int) {
	if t == nil {
		return
	}
	path := []int{}
	FindPath(t, expectnum, path, 0)
}
func main() {
	data := []int{10, 5, 12, 4, 7}
	t := createTree(data, 0)
	findPath(t, 22)
}

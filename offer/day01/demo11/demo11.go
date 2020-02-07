package main

// 	求二叉树中序遍历任意结点的后继结点

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Father *TreeNode
}

func inorderSuccessor(p *TreeNode) *TreeNode {
	// 如果当前结点有右子树，则输出右子树的最左边结点
	if p.Right != nil {
		p = p.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	}
	// 如果当前结点没有右子树，且有父节点
	for p.Father != nil && p == p.Father.Right {
		p = p.Father
	}
	return p.Father
}
func main() {

}

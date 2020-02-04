### 验证二叉搜索树
对于二叉搜索树，每个结点的的左右子树都是一个二叉搜索树，因此可以使用递归的方式不停的判断左右子树。
### 思路
***递归的条件***

- 递归结束的条件为遍历的结点为空，可以直接返回true
- 当左子树的结点大于其根节点的值返回false
- 当右子树的结点小于其根节点的值返回true

对于上述思路，因为左子树的结点都是递减，因此对于左子树，根节点是最大值，同理，对于右子树，根节点是最小值。因此比较的时候，设置左子树比较的下边界，设置右子树的上边界。

初始化的上边界和下边界可以是max和min。当遍历左子树的时候，需要将左子树的结点和根节点比较，因此max为根节点；同理当遍历右子树的时候，需要将右子树的结点和根节点比较，此时min为根节点。

```go

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
``
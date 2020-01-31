### 找到和为定值的二叉树路径
路径必须满足根节点到叶子结点
如：找到10，5，12，4，7 中和为22的路径为10，5，7和10，12
```
     10
  5    12
4   7
```

### 思路
因为必须包含根节点，只有前序遍历才能经过根节点，因此解题思路应该是在前序遍历的基础上进行改进
当遍历到4时发现经过的路径中还未出现定值，此时递归返回到父节点，根据题目要求，可知应该将4进行
删除，即递归返回父节点之前，应该先将当前的结点进行删除，再继续进行相加判断的操作。

### 代码
```go
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
```
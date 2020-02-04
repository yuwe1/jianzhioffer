### 二叉树层次遍历存储
要求输出一个二维数组，每一层为一个一维数组进行存储，最后追加到二维数组中。

### 思路
***层数的+1可能性为遍历一个不为空的左子树或者一个不为空的右子树。***

但是有可能出现左子树和右子树处于同一层，这时只需要对层数进行一个判断，将相同层数的结点追加到一维数组中。初始化层数为0；

### 代码思路
- 定义一个二维数组
- 定义一个一维数组
- 当前的层数如果等于二维数组的长度，将二维数组的长度加一，追加值到当前层数的一维数组中。
- 如果left不为空，层数加一
- 如果right不为空，层数加一

```go
var levels [][]int

func helper(treenode *TreeNode, level int) {
	if len(levels) == level {
		levels = append(levels, []int{})
	}
	levels[level] = append(levels[level], treenode.Val)
	if treenode.Left != nil {
		helper(treenode.Left, level+1)
	}
	if treenode.Right != nil {
		helper(treenode.Right, level+1)
	}
}
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return levels
	}
	helper(root, 0)
	return levels
}
```


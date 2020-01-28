### 二叉树的镜像
        8
   6        10            
5    7   9      11

        8
   10        6            
11    9   7      5


### 思路
对于此类题，我们可以将跟结点的左右子树进行反转，再分别最各自的左右子树进行反转

```go
func MicrrorRecusively(root *treeNode) {
	if root == nil || (root.left == nil && root.right == nil) {
		return
	}
	root.left, root.right = root.right, root.left
	if root.left != nil {
		MicrrorRecusively(root.left)
	}
	if root.right != nil {
		MicrrorRecusively(root.right)
	}
}
```
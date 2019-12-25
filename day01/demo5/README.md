### 题目要求
根据前序遍历和中序遍历重建一个二叉树
```
//前序遍历
{1, 2, 4, 7, 3, 5, 6, 8}
//中序遍历
{4, 7, 2, 1, 5, 3, 8, 6}
```
### 解题思路
解答这个题之前，需要我们了解有关数据结构中二叉树树的相关知识，二叉树的结构如下图所示

![树的表现形式](./../../img/树的表现形式.jpg)

其次我们还需要知道二叉树的前序遍历和中序遍历的特点,对于前序遍历，根节点一定是遍历之后的第一个结果，对于中序遍历，根节点一定是在中间，且根节点的左侧都是左子树，右侧都是右子树

知道了上面的性质仍然不够，我们还需要理解在树中递归的思想。

举一个例子，对于一个二叉树，当我们进行前序遍历的时候，我们的代码通常用递归的方式完成，让我们来具体看看这个递归
```go
// 二叉树的前序遍历
func preorder(root *Node) {
	if root == nil {

		return
	}
	fmt.Print(root.data)
	preorder(root.left)
	preorder(root.right)
}
```
在这个函数中，我们不停的递归我们的树的子树，将子树当成新的树进行递归输出，那么当我们构建的时候，我们根节点的左子树也可以不断的赋值我们的根节点根据我们输出的序列。知道这个思路我们就能很轻松的写出代码了,因为我们分别找到了左、右子树的前序遍历序列和中序遍历序列，我们就可以用同样的方法分别去构建左右子树。换句话说，这是一个递归的过程。

### 代码思路
```
1、根据前序遍历的性质找到我们的根节点
2、根据根节点和中序遍历的结果比较出该树的左子树部分和右子树部分
3、对每一个子树进行上述操作赋值根节点
```
### 完整代码
```go
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
```
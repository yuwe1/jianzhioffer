### 题目要求
判断A，B两个二叉树，B是否是A的子树

### 题目分析
对于这个题，首先我们需要知道二叉树的创建，二叉树的种类有很多，这一题中我们先回顾一下二叉树的基本知识，以二叉查找树为例。

#### 二叉树
二叉树是一种存储数据的结构，有一个根节点，每个结点都有左右指针，指向一个新的结点，称根节点的子节点，而子节点也可以有子节点的树状结构。

#### 二叉查找树
***性质***
- 任意结点倘若它的左子树不为空，则左子树上的所有结点的值均小于它的根节点的值。
- 任意结点倘若它的右子树不为空，则右子树上的所有结点的值均大于它的根节点的值。

***二叉查找树的创建***
基于上述的性质，我们可以动态的创建一个查找树。步骤如下：
- 传入一个新结点到一颗二叉树。
-  如果该二叉树的根节点为空，则将此结点置为根节点。
- 如果二叉树根节点不为空，则将此节点将根节点的值进行比较。
- 如果结点的值小于根结点，则继续比较该结点的左结点，直到待比较的结点为空。
- 同理，如果结点的值大于根结点，则继续比较该结点的右结点，直到待比较的结点为空。

***定义一个结点类***
```go
type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}
```

***定义一个二叉树的类***
```go
// 定义一个二叉树类
type binaryTree struct {
	// 根结点
	root *treeNode
```

***创建***
```go
func (t *binaryTree) CreateTree(val int) {
	if t.Root() == nil {
		t.SetRoot(treeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		})
	} else {
		root := t.Root()
		for root != nil {
			if val <= root.Val && root.Left == nil{
				root.Left = &treeNode{val,nil,nil}
				break
			} else if val > root.Val && root.Right == nil {
				root.Right = &treeNode{val,nil,nil}
				break
			} else if val <= root.Val && root.Left != nil {
				root = root.Left
			} else if val > root.Val && root.Right != nil {
				root = root.Right
			}
		}
	}
}
```

接着我们回到题目分析的思路上，对于判断B树是否是A树的子树，首先应该判断B树的根节点是否存在于A树中，如果存在，则继续判断B树的左右子树，是否和找到的A树中相同根节点的左右子树相同。这是一个基本的思路，接着我们具体看一下步骤：
- 首先从A树的根节点开始遍历，如果根节点相同，则比较B树的左右结点和A树根结点的左右结点
- 如果不相同，则遍历A树的左右结点进行比较

***代码***
```go
func HasSubtree(pRoot1 *treeNode, pRoot2 *treeNode) bool {
	result := false
	if pRoot1 != nil && pRoot2 != nil {
        //找到可以开始比较的根节点
		if pRoot1.Val == pRoot2.Val {
            // 进行比较
			result = DoesTreeHHaveTree2(pRoot1, pRoot2)
        }
        // 开始遍历左节点
		if !result {
			result = HasSubtree(pRoot1.Left, pRoot2)
        }
        // 开始遍历右结点
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
    // 不断的递归左右结点进行比较
	return DoesTreeHHaveTree2(pRoot1.Left, pRoot2.Left) && DoesTreeHHaveTree2(pRoot1.Right, pRoot2.Right)
}
```
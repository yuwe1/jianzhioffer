### 判断一个数组是否可能是一颗二叉查找树的后序遍历
如5,7,6,9,11,10,8 <br />
```txt
      8
  6     10
5   7  9  11
```

### 思路
对于一个二叉查找树后序遍历有如下性质
- 输出的最后的值一定是二叉树的根节点
- 左子树一定都小于根节点，右子树一定都大于根节点。
- 左右子树的输出也满足以上性质

对于5，7，6，9，11，10，8 可以得出8为根节点，5，7，6为左子树，9，11，10为右子树，因此递归推断可知，每个子树都满足二叉查找树输出的性质即为true,否则为false

### 代码思路
- 找到根节点
- 找到左边比根节点小的序列，
- 判断右边剩余的序列是否均大根节点，如果大于，满足条件为true，否则返回false
- 继续判断左边的序列和右边的序列是否也满足二叉查找树的性质

```go
func IsAfter(data []int) bool {

	root := data[len(data)-1]
	i := 0
	for i = 0; i < len(data)-2; i++ {
		if data[i] > root {
			break
		}
	}
	j := i
	// 判断右边的数组是i否存在比根节点小的值
	for i := j; i < len(data)-2; i++ {
		if data[i] < root {
			return false
		}
	}
	right := true
	// 继续判断右边的数组
	if i > 0 {
		right = IsAfter(data[:j])
	}
	left := true
	if j < len(data)-1 {
		left = IsAfter(data[j : len(data)-1])
	}
	return left && right
}

```
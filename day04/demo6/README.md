### 二叉树从上到下的输出
如：
```txt
      8
  6     10
5   7  9  11
```
输出:
8，6，10，5，7，9，11

### 思路
对于输出8之后，我们需要记住该根节点的左右子节点，然后继续输出左右子节点，输出之前还要先记住第三层，由此思路可知，当我们输出值的时候，可以先把当前输出的结点的左右结点进行保存，然后当输出左右结点的时候，也要先保存各自的左右结点，不断的保存和输出，最终保存的容器的容量为0的时候，说明输出结束

### 代码思路
- 指针指向根节点进行保存到数组中
- 循环保存当前输出结点的左右结点，同时输出，并从数组中删除该结点【即第一位】
- 指针指向当前输出的结点。

```go
func printFromtop(root *treeNode) {
	// 将根节点进行入队
	data := []*treeNode{}
	data = append(data, root)
	p := root
	for len(data) > 0 {
		fmt.Println(data[0].value)
		p = data[0]
		data = data[1:]

		if p.left != nil {
			data = append(data, p.left)
		}
		if p.right != nil {
			data = append(data, p.right)
		}
	}
}
```

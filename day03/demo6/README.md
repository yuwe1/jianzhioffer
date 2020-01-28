## 反转链表
定义一个函数，输入一个链表的头结点，反转该链表并输出反转后链表的头结点。链表结点定义如下：
```go
type ListNode struct{
    m_nkey int
    m_pNext *ListNode
}
```

## 思路
假设有3个节点，a——>b——>c,现在要实现反转，第一步首先需要将b——>a,接着再把c——>b，因为在链表中的节点是由前一个结点找到的，因此应该实现找到b节点用一个临时变量存储，否则b节点在经历第一步后会发生改变。当反转结束的时候根据next为nil进行判断。

```go
func (l *ListNode) reverselist() {

	// 代表前指针
	var pre *Node
	// 当前指针
	p := l.head
	for p != nil {
		// 后指针
		temp := p.n_pNext
		if temp == nil {
			l.head = p
		}
		p.n_pNext = pre
		pre = p
		p = temp
	}
}
```
- [链表的倒数第K个结点——完整代码](https://gitee.com/yuweiwuyazi/jianzhioffer/tree/master/day03/demo6)
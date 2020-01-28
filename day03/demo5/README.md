### 链表的倒数第K个结点
输入一个链表。输出该链表的倒数第K个结点，即倒数第一个结点为尾结点。

### 解题思路
对于一个有头节点的链表，求倒数第k个结点，可用双指针法，如对于如下链表：
```
head——1——2——3——4——5——6
```
求倒数第一个，我们可以让双指针指向1，向后遍历，直到后指针.next域为nil
求倒数第二个，我们可以让前指针指向1，后指针指向2，直到后指针的.next域为nil
求倒数第三个，我们可以让前指针指向1，后指针指向3，直到后指针的.next域为nil
.....
可以得出，当求倒数第k个时，前后指针相差了k-1个位置，因此我们可以循环k-1次，找到后指针，然后再循环。

```go
func findKfromTail(head *ListNode, k int) {
	if head == nil || k == 0 {
		fmt.Println("请输入有效的k")
		return
	}
	phead := head.m_pNext
	for i := 0; i < k-1; i++ {
		if phead.m_pNext != nil {
			phead = phead.m_pNext
		} else {
			return
		}
	}
	pBhind := head.m_pNext
	for phead.m_pNext != nil {
		phead = phead.m_pNext
		pBhind = pBhind.m_pNext
	}
	fmt.Println(pBhind.m_nValue)
}
```

- [链表的倒数第K个结点——完整代码](https://gitee.com/yuweiwuyazi/jianzhioffer/tree/master/day03/demo5)
### 复制复杂链表——递归与非递归
定义树的结点如下，该结点不仅有一个next指针，而且还有一个指向任意结点的指针。实现一个clone函数。该结点形成的链表形式如下图：
<div style="align: center">
<img src="./../../img/album01.png"/>
</div>
<br/>

***结点***
```go
type ComListNode struct {
	value   int
	next    *ComListNode
	slibing *ComListNode
}
```
### 解题思路
在复杂链表的结点中除了next指针还有一个指向一个任意结点的指针，指向的结点有可能在该结点的前方，也有可能在该结点的后方，为了解决这个问题，我们可以将复制过程分成多个步骤，首先我们可以只复制每个结点，包含next指针，但是我们不能把该结点形成的链表单独拿出来，因为我们还要继续复制任意结点的指针，我们必须知道被复制的结点，比如上图中的A结点，复制成了一个A',那么必须能根据A'找到A的位置，才能迅速定义到C。因此我们可以复制成结点成如下链表。
<div style="align: center">
<img src="./../../img/album02.png"/>
</div>
<br/>
所以当我们遍历到A',A的slibing.next指针就是A'的slibing,最后我们仅仅分离出复制结点就行了。
<div style="align: center">
<img src="./../../img/album03.png"/>
</div>
<br/>
### 三步递归

***复制结点***
```go
func CopyNode(head *ComListNode) {
	copy := new(ComListNode)
	copy.value = head.value
	copy.next = head.next
	copy.slibing = nil
	head.next = copy
	CopyNode(copy.next)
}
```
***复制任意结点指针***
```go
func CopyComNode(head *ComListNode) {
	if head == nil {
		return
	}
	pclone := head.next
	if head.slibing != nil {
		pclone.slibing = head.slibing.next
	}
	CopyComNode(pclone.next)
}
```
***分离出链表***
```go
// 拆分成两个链表，返回复制的链表
// 将偶数结点相连起来
func getCopy(head *ComListNode) (pcopy *ComListNode) {
	if head == nil {
		return
	}

	pcopy = head.next
	pcopy.next = getCopy(pcopy.next)
	return pcopy
}
```

***关于非递归解法请查看github仓库***
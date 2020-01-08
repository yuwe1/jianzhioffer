### 在O(1)时间删除链表结点
给定单响链表的头指针和一个结点指针,定义一个函数在O(1)时间删除结点，链表结点与函数的定义如下：
```go
type ListNode struct {
	m_nValue int
	m_pNext  *ListNode
}
```
### 题目分析
这道题主要考察了我们对链表的基本操作，同时也考察了我们的思维灵活度。首先我们来看常规的链表操作应该怎么写？首先我们需要定义一个链表操作对象，用来存放链表的头指针和链表的数据长度如下：
```go
type ListN struct {
	head   *ListNode
	length int
}
```

接着我们需要获得一个实例化的对象用来操作链表,这里需要注意的是我们让链表的头节点指向了第一个结点，并且该结点目前是空值。
```go
func Init() *ListN {
	l := new(ListN)
	l.head = nil
	l.length = 0
	return l
}
```

我们试着添加一个结点的操作，添加结点我们需要注意的是当链表的长度为0的时候，我们将结点的数据存放在头节点中。如果不是0，需要遍历链表，将新结点追加到尾结点，遍历如下：
```go
head := l.head
for head.m_pNext != nil {
	head = head.m_pNext
}
head.m_pNext = &node
```
成功添加结点后，我们编写一个遍历链表的函数，用来测试我们的数据是否真正添加到链表中。这里需要注意的是遍历的时候应该从头结点遍历，因为我们的头节点也存在了数据，因此这里的循环和添加结点的循环不太一样：
```go
func (l *ListN) ReadNode() {
	head := l.head
	if head == nil {
		fmt.Println("该链表没有结点")
	}
	for head != nil {
		fmt.Println(head.m_nValue)
		head = head.m_pNext
	}
}
```

现在进入正题，题目要求删除结点要用O(1)的时间复杂度，回想一下我们的数据结构中删除链表结点时的思路，通常我们会准备两个指针指向链表，不停的移动指针，设p1、p2分别为前指针和后指针，那么当p2指向的结点需要被删除的时候，就是把p1指向结点的指针域指向p2结点的指针域指向的结点,有点绕，用代码表示就是
```go
p1->next = p2->next
delete p2
```
那么这就要求我们删除第n个结点，必须要至少遍历n-1次，时间复杂度为O(n)

那么有没有一种思路我们不需要知道待删除结点的前一个结点，就能将其删除呢？

其实我们可以试图去想如果我们把待删除的结点的值赋值成下一个结点，这时我们删除下一个结点，就能重新形成链表了。不过这里我们需要注意的时尾结点没有下一节点我们直接判断，还需注意的是待删除的结点的必须真实存在与链表中，值和next域均存在。
```go
func (l *ListN) DeleteNode_two(node ListNode) {
	if l.length == 0 {
		fmt.Println("该链表没有结点，无法进行删除操作")
		return
	}
	// 如果删除的结点是不是尾结点
	if node.m_pNext != nil {
		next := node.m_pNext
		node.m_nValue = next.m_nValue
		node.m_pNext = next.m_pNext
		next = nil
	} else if l.head == &node { //链表只有一个结点，删除头节点（也是尾结点）
		l.head = nil
	} else {
		// 链表中有多个结点，删除尾结点
		head := l.head
		for head.m_pNext != &node {
			head = head.m_pNext
		}
		head.m_pNext = nil
	}
}
```



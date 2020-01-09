package main

import (
	"fmt"
)

type ListNode struct {
	m_nValue int
	m_pNext  *ListNode
}

// 定义一个链表
type ListN struct {
	head   *ListNode
	length int
}

// 初始化一个链表
func Init() *ListN {
	l := new(ListN)
	l.head = nil
	l.length = 0
	return l
}

func (l *ListN) AddNode(node ListNode) {
	if l.length == 0 {
		l.head = &node
	} else {
		head := l.head
		for head.m_pNext != nil {
			head = head.m_pNext
		}
		head.m_pNext = &node
	}
	l.length++
}

// 遍历结点
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

func (l *ListN) FindKFromTail(k int) {
	if k <= 0 || k > l.length {
		fmt.Println("请输入有效的k")
		return
	}
	// 获取链表的长度
	length := l.length
	head := l.head
	for i := 0; i < length-k; i++ {
		head = head.m_pNext
	}
	fmt.Println(head.m_nValue)
}
func main() {
	// l := Init()
	// l.AddNode(ListNode{
	// 	m_nValue: 1,
	// 	m_pNext:  nil,
	// })
	// l.AddNode(ListNode{
	// 	m_nValue: 2,
	// 	m_pNext:  nil,
	// })
	// l.AddNode(ListNode{
	// 	m_nValue: 3,
	// 	m_pNext:  nil,
	// })
	// l.AddNode(ListNode{
	// 	m_nValue: 4,
	// 	m_pNext:  nil,
	// })
	// l.ReadNode()
	// fmt.Println("***************")
	// l.FindKFromTail(1)
	// l.FindKFromTail(0)
	// l.FindKFromTail(2)
	// l.FindKFromTail(3)
	// l.FindKFromTail(4)
	// l.FindKFromTail(5)
	// 定义一个头节点
	head := new(ListNode)
	addNode(head, ListNode{
		m_nValue: 1,
		m_pNext:  nil,
	})

	addNode(head, ListNode{
		m_nValue: 2,
		m_pNext:  nil,
	})
	addNode(head, ListNode{
		m_nValue: 3,
		m_pNext:  nil,
	})
	readNode(head)

	findKfromTail(head, 1)
	// head = nil
	findKfromTail(head, 2)
	findKfromTail(head, 2)
}

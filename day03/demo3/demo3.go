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

// 使用顺序查找删除结点,时间复杂度位O(n)
func (l *ListN) DeleteNode(node ListNode) {
	head := l.head
	// 删除的结点时头节点
	if head.m_nValue == node.m_nValue {
		l.head = head.m_pNext
		l.length--
		return
	}
	p1 := head
	p2 := head.m_pNext
	for p2 != nil {
		if p2.m_nValue == node.m_nValue {
			p1.m_pNext = p2.m_pNext
			l.length--
			break
		}
		p1 = p1.m_pNext
		p2 = p2.m_pNext
	}
	if p2 == nil {
		fmt.Println("未找到该结点")
	}
}

// 删除结点，前提是该节点必须是在链表中，值和next都有存在
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
func main() {
	l := Init()
	l.AddNode(ListNode{
		m_nValue: 1,
		m_pNext:  nil,
	})
	l.AddNode(ListNode{
		m_nValue: 2,
		m_pNext:  nil,
	})
	l.ReadNode()
	l.DeleteNode(ListNode{
		m_nValue: 2,
		m_pNext:  nil,
	})
	l.ReadNode()
	l.ReadNode()
}

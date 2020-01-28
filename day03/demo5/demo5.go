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

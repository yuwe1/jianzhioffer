package main

import "fmt"

type Node struct {
	m_nkey  int
	n_pNext *Node
}
type ListNode struct {
	head   *Node
	length int
}

func NewList() *ListNode {
	return &ListNode{
		nil, 0,
	}
}

// 添加一个结点
func (l *ListNode) addNode(node Node) {
	if l.length == 0 {
		l.head = &node
	} else {
		p := l.head
		for p.n_pNext != nil {
			p = p.n_pNext
		}
		p.n_pNext = &node
	}
	l.length++
}
func (l *ListNode) ReadNode() {
	p := l.head
	for p != nil {
		fmt.Println(p.m_nkey)
		p = p.n_pNext
	}
}

// 反转链表
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
func main() {
	l := NewList()
	l.addNode(Node{
		1, nil,
	})
	l.addNode(Node{
		2, nil,
	})
	l.ReadNode()
	l.reverselist()
	l.ReadNode()
}

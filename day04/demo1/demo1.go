package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func New() *ListNode {
	return new(ListNode)
}

// 添加一个结点
func Add(head *ListNode, Node ListNode) {
	if head.next == nil {
		head.next = &Node
		return
	}
	p := head.next
	for p.next != nil {
		p = p.next
	}
	p.next = &Node
}

// 遍历结点
func TravelNode(head ListNode) {
	p := head.next
	for p != nil {
		fmt.Print(p.value, " ")
		p = p.next
	}
	fmt.Println()
}

// 删除链表中的重复结点
func MoveRepeateNode(head *ListNode) {
	// 定义一个map
	m := map[interface{}]bool{}
	p1 := head
	p2 := p1.next
	for p2 != nil {
		if _, ok := m[p2.value]; ok {
			p1.next = p2.next
			p2 = p1.next
		} else {
			m[p2.value] = true
			p1 = p1.next
			p2 = p2.next
		}
	}
}

func main() {
	head := New()
	Add(head, ListNode{value: 1})
	Add(head, ListNode{value: 1})
	Add(head, ListNode{value: 2})
	Add(head, ListNode{value: 2})
	Add(head, ListNode{value: 1})
	TravelNode(*head)
	MoveRepeateNode(head)
	TravelNode(*head)
}

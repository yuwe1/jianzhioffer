package main

import "fmt"

func addNode(head *ListNode, node ListNode) {
	phead := head
	for phead.m_pNext != nil {
		phead = phead.m_pNext
	}
	phead.m_pNext = &node
}

func readNode(head *ListNode) {
	phead := head.m_pNext
	if phead == nil {
		fmt.Println("该链表没有结点")
		return
	}
	for phead != nil {
		fmt.Println(phead.m_nValue)
		phead = phead.m_pNext
	}
}

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
func main() {
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
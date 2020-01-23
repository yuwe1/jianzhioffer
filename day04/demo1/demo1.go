package main

import "fmt"

import "github.com/songtianyi/rrframework/structures/stack"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New() *ListNode {
	return new(ListNode)
}

// 添加一个结点
func Add(head *ListNode, Node ListNode) {
	if head.Next == nil {
		head.Next = &Node
		return
	}
	p := head.Next
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &Node
}

// 遍历结点
func TravelNode(head *ListNode) {
	p := head.Next
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	fmt.Println()
}

// 删除链表中的重复结点
func MoveRepeateNode(head *ListNode) {
	// 定义一个map
	m := map[interface{}]bool{}
	p1 := head
	p2 := p1.Next
	for p2 != nil {
		if _, ok := m[p2.Val]; ok {
			p1.Next = p2.Next
			p2 = p1.Next
		} else {
			m[p2.Val] = true
			p1 = p1.Next
			p2 = p2.Next
		}
	}
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head
	var pre *ListNode
	pre = nil
	p := p1.Next
	for p != nil {
		temp := p.Next
		if temp == nil {
			head.Next = p
		}
		p.Next = pre
		pre = p
		p = temp
	}
	return p1
}

//反转链表，递归解决
func reverseList_two(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l := reverseList_two(head.Next)
	head.Next.Next = head
	head.Next = nil
	return l
}

// 计算两个单链表所代表的数之和
func ListAdd(h1 *ListNode, h2 *ListNode) *ListNode {
	if h1 == nil || h1.Next == nil {
		return h2
	}
	if h2 == nil || h2.Next == nil {
		return h1
	}
	h1 = reverseList(h1)
	// TravelNode(*h1)
	h2 = reverseList(h2)
	// TravelNode(*h2)
	head := New()
	// 定义两个指针
	p1 := h1.Next
	p2 := h2.Next

	a := 0
	val := 0
	for p1 != nil && p2 != nil {
		val = (p1.Val + p2.Val + a)
		if val >= 10 {
			val = val % 10
		}
		Add(head, ListNode{Val: val})
		a = (p1.Val + p2.Val + a) / 10
		p1 = p1.Next
		p2 = p2.Next
	}
	for p1 != nil {
		val = p1.Val + a
		if val >= 10 {
			val = val % 10
		}
		Add(head, ListNode{Val: val})
		a = (p1.Val + a) / 10
		p1 = p1.Next
	}
	for p2 != nil {
		val = p2.Val + a
		if val >= 10 {
			val = val % 10
		}
		Add(head, ListNode{Val: val})
		a = (p2.Val + a) / 10
		p2 = p2.Next
	}
	if a > 0 {
		Add(head, ListNode{Val: a})
	}
	return reverseList(head)
}

// 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l1.Next == nil {
		return l2
	}
	if l2 == nil || l2.Next == nil {
		return l1
	}
	// 定义二个指针
	p1 := l1
	p11 := p1.Next
	p2 := l2.Next
	for p11 != nil && p2 != nil {
		for p11 != nil && p11.Val <= p2.Val {
			p1 = p1.Next
			p11 = p11.Next
		}
		if p11 != nil {
			temp := &ListNode{
				Val:  p2.Val,
				Next: p11,
			}
			p1.Next = temp
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	if p2 != nil {
		p1.Next = p2
	}
	return l1
}

// 判断链表是否是回文
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	// 链表反转
	l := reverseList(head)
	p1 := head.Next
	p2 := l.Next

	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			break
		}
		p1, p2 = p1.Next, p2.Next
	}
	if p1 != nil || p2 != nil {
		return false
	}
	return true
}

func main() {
	head1 := New()
	Add(head1, ListNode{Val: 1})
	Add(head1, ListNode{Val: 2})
	Add(head1, ListNode{Val: 1})
	TravelNode(head1)
	fmt.Println(isPalindrome(head1))

}

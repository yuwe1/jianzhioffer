package main

import "fmt"

type object interface{}

//定义节点
type Node struct {
	value object
	next  *Node
}

// 定义一个链表
type ListNode struct {
	size int
	head *Node
	tail *Node
}

// 初始化链表
func (l *ListNode) Init() {
	(*l).size = 0
	(*l).head = nil
	(*l).tail = nil
}

// 初始化链表
func (l *ListNode) Init_two() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

// 添加元素
func (l *ListNode) Append(node *Node) {
	if node == nil {
		return
	}
	(*node).next = nil
	if (*l).size == 0 {
		(*l).head = node
	} else {
		oldTail := (*l).tail
		(*oldTail).next = node
	}
	(*l).tail = node
	(*l).size++
}

// 插入元素
func (l *ListNode) Insert(node *Node, n int) {
	if node == nil || n > (*l).size+1 {
		return
	}
	if n == 1 {
		node.next = (*l).head
		(*l).head = node
		(*l).size++
		return
	}
	head := (*l).head
	rhead := head.next
	i := 1
	for i < n-1 {
		head = head.next
		rhead = head.next
		i++
	}

	head.next = node
	node.next = rhead
	if (*l).tail == head {
		(*l).tail = node
	}
	(*l).size++
}

// 遍历链表
func (l *ListNode) Print() {
	head := (*l).head
	for head != nil {
		fmt.Print(head.value, " ")
		head = head.next
	}
	fmt.Println()
}

// 打印尾节点
func (l *ListNode) TailPrint() {
	tail := (*l).tail
	fmt.Println(tail.value)

}

// 利用递归从后往前打印
func (l *ListNode) TailPrint_two(head *Node) {

	if head != nil {
		l.TailPrint_two(head.next)
		fmt.Print(head.value, " ")
	}
}
func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	l := new(ListNode)
	l.Init()
	for i := 0; i < len(data); i++ {

		l.Append(&Node{
			value: data[i],
			next:  nil,
		})
	}
	l.Print()
	l.Insert(&Node{
		value: 100,
		next:  nil,
	}, 1)
	l.Print()
	l.Insert(&Node{
		value: 102,
		next:  nil,
	}, 10)
	l.Print()
	l.TailPrint()
	head := l.head
	l.TailPrint_two(head)
}

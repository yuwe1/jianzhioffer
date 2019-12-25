package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

// 定义一个队列
type Queue struct {
	length int
	head   *Node //队列的头节点
	tail   *Node //队列的尾指针
}

// 定义一个栈
type Stack struct {
	// 栈顶指针
	top *Node
	//栈的深度
	h int
	// 定义两个队列
	q1 Queue
	q2 Queue
}

// 初始化队列
func InitQueue() Queue {
	return Queue{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

// 队列的基本操作
func (q *Queue) AppendTail(node Node) {
	if q.length == 0 {
		q.head = &node
	} else {
		q.tail.next = &node
	}
	q.tail = &node
	q.length++
}

// 队列的删除头节点
func (q *Queue) DeleteHead() (node Node) {
	if q.length == 0 {
		fmt.Println("队列为nil")
		return
	}
	node = *(q.head)
	q.head = q.head.next
	q.length--
	return node
}

// 实现栈的压入
func (s *Stack) Push(node Node) {
	s.q1.AppendTail(node)
	s.h++
	s.top = &node
}

// 实现栈的弹出
func (s *Stack) Pop() Node {
	if s.q2.length == 0 {
		for s.q1.length > 1 {
			s.q2.AppendTail(s.q1.DeleteHead())
		}
	}
	node := s.q1.DeleteHead()
	// 重新回答q1
	if s.q1.length == 0 {
		for s.q2.length > 0 {
			s.q1.AppendTail(s.q2.DeleteHead())
		}
	}
	s.h--
	return node
}

func main() {
	// 定义一个栈
	s := Stack{
		top: nil,
		h:   0,
		q1:  InitQueue(),
		q2:  InitQueue(),
	}
	// 原始数据
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range data {
		s.Push(Node{
			data: v,
			next: nil,
		})
	}
	// for s.q1.length > 0 {
	// 	fmt.Print(s.q1.DeleteHead())
	// }
	for s.h > 0 {
		fmt.Print(s.Pop())
	}
}

package main

import "fmt"

// 用两个栈实现队列
// 定义一个栈结点
type Node struct {
	data int
	next *Node
}

// 定义一个类
type Stack struct {
	// 定义栈顶指针
	top  *Node
	size int
}

// 定义一个队列
type Queue struct {
	s1     Stack
	s2     Stack
	length int
}

// 初始化一个栈
func Init() Stack {
	return Stack{
		top:  nil,
		size: 0,
	}
}
func (s *Stack) Push(data int) {
	node := &Node{
		data: data,
		next: s.top,
	}
	s.top = node
	s.size++

}
func (s *Stack) Pop() (node Node) {
	if s.size == 0 {
		return
	}
	node = *(s.top)
	s.top = s.top.next
	s.size--
	return node
}

//追加尾
func (q *Queue) appendTail(data int) {
	q.s1.Push(data)
	q.length++
}

// 删除头
func (q *Queue) deleteHead() (data int) {
	if q.s1.size == 0 {
		fmt.Println("暂无数据")
		return
	}
	if q.s2.size <= 0 {
		for q.s1.size > 0 {
			q.s2.Push(q.s1.Pop().data)
			q.length--
		}
	}
	data = q.s2.Pop().data
	// 重新压入数据到s1
	for q.s2.size > 0 {
		q.appendTail(q.s2.Pop().data)
	}
	return data
}
func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	q := &Queue{
		s1:     Init(),
		s2:     Init(),
		length: 0,
	}
	for _, v := range data {
		q.appendTail(v)
	}
	for q.length > 0 {
		fmt.Print(q.deleteHead())
	}
	fmt.Print(q.deleteHead())
}

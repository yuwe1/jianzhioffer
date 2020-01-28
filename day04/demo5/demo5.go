package main

import "fmt"

import "math"

// 定义一个结点
type Node struct {
	value int
	next  *Node
}

// 定义一个栈
type stack struct {
	top *Node //栈顶指针
	h   int   //栈的高度
}

var s = &stack{nil, 0}

// 压入栈
func (s *stack) push(value int) {
	top := s.top
	node := &Node{value, top}
	s.top = node
	s.h++
}

// 弹出
func (s *stack) pop() int {
	if s.top == nil {
		return math.MinInt64
	}
	value := s.top.value
	s.top = s.top.next
	return value
}

// 判断序列是否是另一个序列的弹出序列
func IsPopOrder(pushdata []int, popdata []int) bool {
	j := 0
	i := 0
	for j < len(popdata) {
		if i < len(pushdata) {
			s.push(pushdata[i])
		} else {
			break
		}
		for j < len(popdata) && s.top.value == popdata[j] {
			fmt.Println(s.pop())
			j++
		}
		i++
	}
	if s.top == nil {
		return true
	}
	return false
}
func main() {
	pushdata := []int{1, 2, 3, 4, 5}
	popdata := []int{4, 5, 3, 2, 1}
	pop2data := []int{4, 3, 5, 1, 2}
	fmt.Println(IsPopOrder(pushdata, popdata))
	fmt.Println(IsPopOrder(pushdata, pop2data))
}

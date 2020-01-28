package main

import "fmt"

// 定义一个Node
type Node struct {
	value int
	next  *Node
}

// 定义一个栈
type stack struct {
	top *Node //栈顶指针
	h   int   //高度
	min []int //辅助栈数组
}

var s = &stack{nil, 0, nil}

// 往栈中推入元素
func (k *stack) push(value int) {
	top := k.top
	node := &Node{value, nil}
	node.next = top
	k.top = node
	k.h++

	if k.h > 1 && value > s.min[len(s.min)-1] {
		s.min = append(s.min, s.min[len(s.min)-1])
	} else {
		s.min = append(s.min, value)
	}

}

// 弹出栈
func (k *stack) pop() int {
	if k.top == nil {
		fmt.Println("栈中没有元素")
		return 0
	}
	val := k.top.value
	k.top = k.top.next
	// 删除当前最小元素
	s.min = s.min[:len(s.min)-1]
	return val
}

// 获得栈中的最小元素
func (k *stack) minvalue() int {
	if len(k.min) == 0 {
		fmt.Println("没有最小元素")
		return 0
	}
	return k.min[len(k.min)-1]
}

//
func main() {
	s.pop()
	s.push(3)
	fmt.Println("最小元素：", s.minvalue())
	s.push(4)
	fmt.Println("最小元素：", s.minvalue())
	s.push(2)
	fmt.Println("最小元素：", s.minvalue())
	s.push(1)
	fmt.Println("最小元素：", s.minvalue())
	fmt.Println(" 开始弹出元素")
	s.pop()
	fmt.Println("最小元素：", s.minvalue())
	s.pop()
	fmt.Println("最小元素：", s.minvalue())
	s.pop()
	fmt.Println("最小元素：", s.minvalue())
	s.pop()
	fmt.Println("最小元素：", s.minvalue())

	s.pop()
}

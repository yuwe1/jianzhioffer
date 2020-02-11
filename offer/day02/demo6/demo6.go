package main

import "fmt"

// 重排链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func recorderList(head *ListNode) {
	// 快慢指针分离链表
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow
	left := head
	right := mid.Next
	mid.Next = nil

	// 反转后面部份的链表
	var preNode *ListNode
	curNode := right
	var nextNode *ListNode
	for curNode != nil {
		nextNode = curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = nextNode
	}

	right = preNode
	// 合并链表
	for left != nil && right != nil {
		// 保存next
		lefttemp := left.Next
		rightemp := right.Next

		left.Next = right
		right.Next = lefttemp

		left = lefttemp
		right = rightemp
	}

	Read(head)

}

// 创建一个链表
func createLink(data []int, index int) (head *ListNode) {
	if len(data) <= index {
		return
	}
	head = &ListNode{data[index], nil}
	head.Next = createLink(data, index+1)
	return head
}

// 遍历链表
func Read(head *ListNode) {
	if head != nil {
		fmt.Println(head.Val)
		Read(head.Next)
	}
}
func main() {
	data := []int{1, 2, 3, 4, 5}
	head := createLink(data, 0)
	recorderList(head)
}

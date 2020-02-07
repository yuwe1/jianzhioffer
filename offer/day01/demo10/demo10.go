package main

import "fmt"

// 从未到头打印链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func printListReversingly(head *ListNode) {
	if head != nil {
		printListReversingly(head.Next)
		fmt.Println(head.Val)
	}
}

func main() {

}

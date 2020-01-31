package main

import "fmt"

// 复杂链表的复制
type ComListNode struct {
	value   int
	next    *ComListNode
	slibing *ComListNode
}

func createlink(data []int, index int) (head *ComListNode) {
	if index >= len(data) {
		return
	}
	head = &ComListNode{data[index], nil, nil}
	head.next = createlink(data, index+1)
	return head
}

func addCom(head *ComListNode) {
	if head == nil || head.next.next == nil {
		return
	}
	head.slibing = head.next.next
	head.next.next.slibing = head.next
	addCom(head.next)
}

func read(head *ComListNode) {
	if head == nil {
		fmt.Println()
		return
	}
	if head.slibing != nil {
		fmt.Println(head.value, "指向了", head.slibing.value)
	}
	read(head.next)
}
func readNosblind(head *ComListNode) {
	if head == nil {
		fmt.Println()
		return
	}
	fmt.Print(head.value, " ")
	readNosblind(head.next)
}
func copyLink(head *ComListNode) *ComListNode {
	CopyNode(head)
	fmt.Println("第一步结果")
	readNosblind(head)
	CopyComNode(head)
	fmt.Println("第二步结果")
	read(head)
	copy := getCopy(head)
	fmt.Println("第三步结果")
	readNosblind(copy)
	read(copy)
	return copy
}

// 第一步复制每一个结点
func CopyNode(head *ComListNode) {
	if head == nil {
		return
	}
	// pNode := head
	// for pNode != nil {
	// 	copy := new(ComListNode)
	// 	copy.value = pNode.value
	// 	copy.next = pNode.next
	// 	copy.slibing = nil
	// 	pNode.next = copy
	// 	pNode = copy.next
	// }
	// 使用递归
	copy := new(ComListNode)
	copy.value = head.value
	copy.next = head.next
	copy.slibing = nil
	head.next = copy
	CopyNode(copy.next)
}

// 复制复杂的结点
func CopyComNode(head *ComListNode) {
	// pNode := head
	// for pNode != nil {
	// 	pclone := pNode.next
	// 	if pNode.slibing != nil {
	// 		pclone.slibing = pNode.slibing.next
	// 	}
	// 	pNode = pclone.next
	// }
	// 使用递归
	if head == nil {
		return
	}
	pclone := head.next
	if head.slibing != nil {
		pclone.slibing = head.slibing.next
	}
	CopyComNode(pclone.next)
}

// 拆分成两个链表，返回复制的链表
// 将偶数结点相连起来
func getCopy(head *ComListNode) (pcopy *ComListNode) {
	if head == nil {
		return
	}

	pcopy = head.next
	pcopy.next = getCopy(pcopy.next)
	return pcopy
}
func main() {
	data := []int{1, 2, 4, 5, 7, 9}
	head := createlink(data, 0)
	addCom(head)
	read(head)
	// read(head)
	fmt.Println("*********")
	copyLink(head)
}

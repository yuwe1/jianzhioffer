package main

// 两个数组实现队列

type MyQueue struct {
	s1 []int
	s2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for len(this.s1) != 0 {
		this.s2 = append(this.s2, this.s1[len(this.s1)-1])
		this.s1 = this.s1[:len(this.s1)-1]
	}
	this.s1 = append(this.s1, x)
	for len(this.s2) != 0 {
		this.s1 = append(this.s1, this.s2[len(this.s2)-1])
		this.s2 = this.s2[:len(this.s2)-1]
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	element := this.s1[len(this.s1)-1]
	this.s1 = this.s1[:len(this.s1)-1]
	return element
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.s1[len(this.s1)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0
}

func main() {

}

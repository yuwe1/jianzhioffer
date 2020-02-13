package main

import "fmt"

// 旋转数组
func roat(data []int, k int) {
	resver(data)
	resver(data[:k])
	resver(data[k:])
}

// 反转
func resver(data []int) {
	l := len(data) - 1
	for i := 0; i <= l/2; i++ {
		data[i], data[l-i] = data[l-i], data[i]
	}
}
func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	roat(data, 3)
	fmt.Println(data)
}

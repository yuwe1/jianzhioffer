package main

import "fmt"

// 递归解法
func Fib(n int) (sum int) {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// 循环解法
func Fib_two(n int) int {
	f1, f2 := 1, 1
	if n == 1 || n == 2 {
		return 1
	}
	f3 := 0
	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f1, f2 = f2, f3
	}
	return f3
}

// 跳台阶问题
func Fib_three(n int) int {
	f1, f2 := 1, 2
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	f3 := 0
	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f1, f2 = f2, f3
	}
	return f3
}
func main() {
	fmt.Println(Fib(10))
	fmt.Println(Fib_two(10))
}

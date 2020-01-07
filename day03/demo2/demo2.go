package main

import (
	"fmt"
	"strconv"
)

/*
1、当n为负数或0时，返回false
2、当n>0时,打印并返会true
*/
func printMaxN(n int64) bool {
	if n <= 0 {
		fmt.Println(0)
		return false
	}
	// 求出最大的n位数 ，若n为3，则9 *10 +9 = 99 ,99 *10 +9
	sum := int64(9)
	for i := int64(0); i < n-1; i++ {
		sum = sum*10 + 9
	}
	fmt.Println(sum)
	return true
}

// 思路2,有溢出的情况,打印出最大值
func printMaxN_two(n int64) bool {
	if n <= 0 {
		fmt.Println(0)
		return false
	}

	// 求出int64的最大值 max := math.MaxInt64
	// 求最大值用左移操作符，提升效率，64位成2^62
	max := int64(1<<63 - 1)

	// 求出ma下对应的最大位数,先将max转换成sting
	smax := strconv.FormatInt(max, 10)
	maxn := len(smax)

	// 溢出时打印最大值
	if n >= int64(maxn) {
		// 将发生溢出
		fmt.Println(smax)
		return false
	}

	// 未溢出正常计算
	sum := int64(9)
	for i := int64(0); i < n-1; i++ {
		sum = sum*10 + 9
	}
	fmt.Println(sum)
	return true
}

// 溢出情况下仍然正常打印，转字符串
func printMaxN_three(n int64) bool {
	// 判断n是否未负数
	if n <= 0 {
		fmt.Println(0)
		return false
	}

	// 这种情况无论是否溢出都进行打印字符串形式
	str := ""
	for i := int64(0); i < n; i++ {
		str += "9"
	}
	fmt.Println(str)
	return true
}

// 递归写法
func printMaxN_four(n uint64) (str string) {
	if n <= 0 {
		return ""
	} else {
		str += "9"
		n = n - 1
	}
	return str + printMaxN_four(n)
}

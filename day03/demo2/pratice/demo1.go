package main

import (
	"fmt"
	"strconv"
)

//实现两个数字的相加，需要考虑溢出问题,因此需要将数字当成字符串计算，
// 用字符串模拟数字的加法
/*
反转之后
1  9  9  1  1  1  9
2  2  3  1  4  1  9
3  11 12 2  5  2  18
*/
func add(num1, num2 int64) {
	str1 := strconv.FormatInt(num1, 10)
	str2 := strconv.FormatInt(num2, 10)

	len1 := len(str1)
	len2 := len(str2)
	if len1 < len2 {
		str1, str2 = str2, str1
		len1, len2 = len2, len1
	}
	rstr1 := []rune(str1)
	rstr2 := []rune(str2)
	fmt.Println(rstr1)
	fmt.Println(rstr2)
	for i := len1 - 1; i > len2-1; i-- {
		rstr2 = append(rstr2, '0')
	}
	// 旋转字符串
	for i := 0; i < len1/2; i++ {
		rstr1[i], rstr1[len1-i-1] = rstr1[len1-1-i], rstr1[i]
	}
	for i := 0; i < len2/2; i++ {
		rstr2[i], rstr2[len2-1-i] = rstr2[len2-1-i], rstr2[i]
	}

	// 字符相加
	s := make([]rune, len1)
	// 初始化进位
	pa := '0' - '0'
	for i := len1 - 1; i >= 0; i-- {
		n1 := rstr1[i] - '0'
		n2 := rstr2[i] - '0'

		s[i] = n1 + n2 + pa
		if s[i] > 10 {
			s[i] = s[i] - 10
			pa = '1' - '0'
		} else {
			pa = '0' - '0'
		}
		fmt.Println(n1, "+", n2, "=", s[i])
	}
	n1 := rstr1[len1-1] - '0'
	n2 := rstr2[len1-1] - '0'
	s[len1-1] = n1 + n2
	for i := 0; i <= len(s)-1; i++ {
		s[i] = s[i] + '0'
	}
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len1-i-1] = s[len1-i-1], s[i]
	}

	fmt.Println(string(s))

}
func main() {

	add(11, 222)
}

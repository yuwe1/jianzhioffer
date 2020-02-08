package main

import "fmt"

// 字符串实现减法
func q(str1 string, str2 string) {
	if len(str1) < len(str2) {
		// 相差的个数
		n := len(str2) - len(str1)
		temp := ""
		for i := 0; i < n; i++ {
			temp += "0"
		}
		str1 = temp + str1
	}
	r1 := []rune(str1)
	r2 := []rune(str2)
	result := make([]rune, len(r1))
	// 双指针
	i, j := len(r1)-1, len(r2)-1
	carry := '0' - '0'
	temp := '0' - '0'
	sign := 1
	for i >= 0 || j >= 0 {
		temp = (r1[i] - '0') - (r2[j] - '0')
		res := (temp + carry)
		sign = 1
		if temp < '0'-'0' {
			res = '0' - '0' - res
			carry = '0' - '1'
			sign = -1
		}
		result[i] = res
		i--
		j--
	}
	hh := ""
	for _, v := range result {
		hh += string(v + '0')
	}
	if sign == -1 {
		fmt.Println("-" + hh)
	}
	fmt.Println(hh)
}

func main() {
	str1 := "3456"
	str2 := "1934"
	q(str1, str2)
}

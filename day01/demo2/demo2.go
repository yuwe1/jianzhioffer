package main

import "fmt"

func ReplaceBank(str string) (r string) {
	count := 0
	// 计算字符串的空格个数
	for _, v := range str {
		if string(v) == " " {
			count++
		}
	}
	// 定义两个指针
	p1 := len(str) - 1
	p2 := len(str) - 1 + 2*count
	strs := []rune(str)
	for i := 0; i < count; i++ {
		temp := ' '
		strs = append(strs, temp, temp)
	}
	for p1 < p2 {
		if string(str[p1]) != " " {
			strs[p2], strs[p1] = strs[p1], strs[p2]
			p1--
			p2--
		} else {
			strs[p2] = '0'
			strs[p2-1] = '2'
			strs[p2-2] = '%'
			p2 -= 3
			p1--
		}
	}
	return string(strs)
}

func main() {
	str := "We are happy."
	// fmt.Println(string(str[2]), "hah")
	// fmt.Println(reflect.TypeOf(str[0]))
	fmt.Println(ReplaceBank(str))
}

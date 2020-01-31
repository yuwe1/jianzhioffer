package main

import "fmt"

// 字符串的排列
func CalcALLPermutation(str *string, from int, to int) {
	strr := []rune(*str)
	if to <= 1 {
		return
	}
	if from == to {
		fmt.Println("最终结果：", string(strr))
	} else {
		for j := from; j <= to; j++ {
			strr[j], strr[from] = strr[from], strr[j]
			*str = string(strr)
			CalcALLPermutation(str, from+1, to)
			fmt.Println(*str)
		}
	}
}
func main() {
	str := "abc"
	CalcALLPermutation(&str, 0, 2)
}

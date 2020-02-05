package main

import (
	"fmt"
	"strings"
)

// 最长无重复字串
func GetStr(str string) string {
	start := 0
	maxstr := ""
	currentstr := ""
	for i := 0; i < len(str); i++ {
		// 查看当前的字符串是否存在当前的游标
		isexist := strings.Index(str[start:i], string(str[i]))
		if isexist == -1 {
			//不存在
			currentstr = str[start : i+1]
			if len(maxstr) < len(currentstr) {
				maxstr = currentstr
			}
		} else {
			// 存在
			start = i
		}
	}
	return maxstr
}

// 使用falg进行标记
func GetStr_two(str string) string {
	start := 0
	maxstr := ""
	currentstr := ""
	flag := make(map[byte]bool, 128)
	key := 0
	for key < len(str) {
		if _, ok := flag[str[key]]; !ok {
			flag[str[key]] = true
			currentstr = str[start : key+1]
			if len(maxstr) < len(currentstr) {
				maxstr = currentstr
			}
			key++
		} else {
			start = key
			flag = make(map[byte]bool, 128)
		}
	}
	return maxstr
}
func main() {
	str := "acdefsfssasbjhkdjfhjkdhfuieuihrhdfhgf"
	fmt.Println(GetStr(str))
}

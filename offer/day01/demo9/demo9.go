package main

import "fmt"

// 空格替换
func replaceSpaces(str string) string {
	result := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			result += "%20"
		} else {
			result += string(str[i])
		}
	}
	return result
}
func main() {
	fmt.Println(replaceSpaces("we happpy a"))
}

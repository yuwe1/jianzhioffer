package main

import (
	"fmt"
)

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	misstr := strs[0]
	for _, k := range strs {
		misstr = helper(misstr, k)
		if misstr == "" {
			return ""
		}
	}
	return misstr
}
func helper(misstr, str string) string {
	l := len(misstr)
	if len(misstr) > len(str) {
		misstr, str = str, misstr
	}
	l = len(misstr)
	for l >= 0 {
		if misstr[:l] != str[:l] {
			l--
		} else {
			return misstr[:l]
		}
	}
	return ""
}
func main() {
	strs := []string{"flower", "flow", ":foverv"}
	fmt.Println(longestCommonPrefix(strs))
}

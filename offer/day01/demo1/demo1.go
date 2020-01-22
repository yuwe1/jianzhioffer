package main

import "fmt"

func StringErrorRecover(strs string) string {
	str := []rune(strs)
	i, j := 0, 2
	for j < len(str) {
		if str[i] == str[j] && str[i] == str[(i+j)/2] {
			str[i] = '#'

		} else if (j+1) < len(str) && str[i] != str[j] && str[i+1] == str[i] && str[j+1] == str[j] {
			str[j] = '#'
			str[i], str[j] = str[j], str[i]
		}
		i++
		j++
	}
	strs = ""
	for i := 0; i < len(str); i++ {
		if str[i] != '#' {
			strs += string(str[i])
		}
	}
	return strs
}
func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	str := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &str[i])
	}
	for i := 0; i < len(str); i++ {
		fmt.Println(StringErrorRecover(str[i]))
	}
}

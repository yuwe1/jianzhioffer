package main

/*
matrix=
[
  ["A","B","C","E"],
  ["S","F","C","S"],
  ["A","D","E","E"]
]

str="BCCE" , return "true"

str="ASAE" , return "false"
*/

import (
	"fmt"
)

func hasPath(matrix [][]byte, str string) bool {
	if len(matrix) == 1 && len(matrix[0]) == 1 && matrix[0][0] == str[0] {
		return true
	}
	i, j := 0, 0
	for i = 0; i < len(matrix); i++ {
		for j = 0; j < len(matrix[i]); j++ {
			if dfs(matrix, str, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(data [][]byte, str string, u, x, y int) bool {
	if u == len(str) {
		return true
	}
	if data[x][y] != str[u] {
		return false
	}
	// 左上右下
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}
	t := data[x][y]
	data[x][y] = '*'
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a >= 0 && a < len(data) && b >= 0 && b < len(data[a]) {
			if dfs(data, str, u+1, a, b) {
				return true
			}
		}
	}
	data[x][y] = t
	return false
}
func main() {
	data := [][]string{
		{"A"},
	}
	fmt.Println(hasPath(data, "A"))
}

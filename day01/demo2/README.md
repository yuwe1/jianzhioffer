## 字符串

## 面试题——字符串去除空格
请实现一个函数，将字符串中的每一个空格替换成"%20".例如输入"We are happy.",则输出"We%20are%20happy."

## 思路
(1)遍历字符串，统计出空格的个数num，用来计算最终转换成的新字符串的总长度，因为每次替换一个空格，字符串的长度为增加2，因此总共的长度为len(str)+2 * num
(2)定义两个指针P1,P2。P1指向原字符串的末尾,P2指向最终新字符串的末尾。
(3)同时移动P1和P2指针，并将P1指向的字符移动到P2的位置，如果P1遇到空格，则P2开始填补%20，当P1与P2相遇时则停止，输出最终结果。

## 完整代码
```go

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
```
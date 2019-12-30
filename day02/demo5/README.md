### 题目要求
在EXCEL中常常用英文字母表示列，如A为1，B为2.....Z为26，依次进行一个循环，AA为27，AB为28，按照这样的规律，要求输入一个字符串求出为第几列。

### 题目分析
这道题很容易得出是一道进制转换的问题，即26进制转换成10进制，因此我们只需要知道不同进制转换成10进制的计算规律即可。
举例"ABCD"
```txt
10进制(D)*26^0+10进制(C)*26^1+...
```
还不理解，可以假设ABCD为数组data，反转为DCBA，i为下标则：
D*26^i+C*26^(++i).....

### 代码思路
```
1、首先我们需要将A~Z和1~26进行绑定
2、遍历我们的字符串，倒叙进行求和
```
### 代码如下
```go
func init() {
	data = make(map[rune]int, 26)
	for i := 'A'; i <= 'Z'; i++ {
		data[i] = int(i) - 64
	}
}
func TwentysixToTen(str string) int {
	if len(str) <= 0 {
		return 0
	}
	sum := 0
	for i, _ := range str {
		j := len(str) - 1 - i
		m := Calaute(i)
		sum += (data[rune(str[j])] * m)
	}
	return sum
}

// 计算26的几次
func Calaute(n int) int {
	sum := 1
	for i := 0; i < n; i++ {
		sum *= 26
	}
	return sum
}
```
在上面的代码中我们由于使用了map，导致了我们的空间复杂度为O(n)，那么如何将空间复杂度降低到O(1)呢？我们可以巧妙的利用字符和数字的转换如下代码
```go
// 第二种解法降低空间复杂度O(1)
func TwentysixToTen_two(str string) int {
	sum := 0
	for i := len(str) - 1; i >= 0; i-- {
		// 确定了每一个字符所对应的整数
		p := int(str[i])
		if (p-64) > 26 || (p-64) < 1 {
			return 0
		}
		m := Calaute(len(str) - 1 - i)
		sum += (p - 64) * m
	}
	return sum
}
```
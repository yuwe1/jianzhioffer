### 题目要求
调整数组位置，使得数组顺序为奇数到偶数

### 题目分析
该题可以说是初级程序员的水平，然而却有很多程序员的解决思路并不是完美。现在一起看看不同程序员的解决思路吧~

### 初级程序员
这道题很简单，只需要遍历数组，判断每一个数字的奇数偶数的性质即可，因此需要准备两个临时数组用来存储，然后再合并即可。
```go
func ReOrderOddEven(data []int) []int {
	if len(data) == 0{
		return data
	}
	var temp1 []int
	var temp2 []int
	lengh := len(data)
	for i := 0; i < lengh; i++ {
		if data[i]&1 == 1 {
			temp1 = append(temp1, data[i])
		} else {
			temp2 = append(temp2, data[i])
		}
	}
	var result []int
	for i := 0; i < len(temp1); i++ {
		result = append(result, temp1[i])
	}
	for i := 0; i < len(temp2); i++ {
		result = append(result, temp2[i])
	}
	return result
}
```

### 中级程序员
这道题很简单，但一般考察这道题不会用那么明显的思路，会不会有更好的思路呢？于是该程序员想到了用两个指针，分别指向头和尾，那么思路就显而易见了

1、如果前指针和后指针指向的数组分别是一个偶数一个奇数，就进行交换位置
2、如果前后一奇数一偶数，则不需要变动，移动指针
3、如果两个都是奇数，只需要移动前指针找到偶数进行交换即可

```go
func ReOrderOddEven_two(data []int) []int {
	// 定义头尾指针
	i, j := 0, len(data)-1
	for i <= j {
		if data[i]&1 == 0 && data[j]&1 == 1 {
			data[i], data[j] = data[j], data[i]
			i++
			j--
		} else if data[i]&1 == 1 && data[j]&1 == 1 {
			i++
		} else if data[i]&1 == 0 && data[j]&1 == 0 {
			j--
		} else {
			i++
			j--
		}

	}
	return data
}
```

### 高级程序员
这道题过于简单，是否可以进行优化和扩展呢？假如我希望将负数排列在前面，正数排列在后面呢?只需要改动for循环中的判断条件，假如我需要将被3整除的放左，其余的放在右边呢？仍然需要改动for循环中的判断，很快，该程序员想到了模板，既然只需要改动部分代码，其实可以将需要改动的写成一个函数，下次想要扩展什么功能，也仅仅需要添加一个小函数，如下，假如仍然需要将奇数排列在前，偶数排列在后，那么我们就可以有这样的函数
```go
func JudgeOddAndEven(data int) bool {
	if data&1 == 1 {
		return true
	}
	return false
}
```

而我们的模板函数就变成了
```go
type F func(int) bool

func ReOrderOddEven_four(data []int, f F) []int {
	if len(data) <= 0 {
		return data
	}
	i, j := 0, len(data)-1
	for i <= j {
		for f(data[i]) && i <= j {
			i++
		}
		for !f(data[j]) && i <= j {
			j--
		}
		if i <= j {
			data[i], data[j] = data[j], data[i]
		}

	}
	return data
}
```

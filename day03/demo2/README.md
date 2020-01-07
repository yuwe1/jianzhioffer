### 题目要求
打印一个n位的最大整数，如3位999，4位9999


### 思路一
分析题目可知n位的最大数一定是n个9组成，举例
```
n = 1时 max = 9
n = 2时 max = 9*10+9
n = 3时 max = 99*10+9
...
```

因此我们可以初始化一个9，循环n次即可，当然这里需要注意的是n>0

**代码如下**
```go
func printMaxN(n int64) bool {
	if n <= 0 {
		fmt.Println(0)
		return false
	}
	// 求出最大的n位数 ，若n为3，则9 *10 +9 = 99 ,99 *10 +9
	sum := int64(9)
	for i := int64(0); i < n-1; i++ {
		sum = sum*10 + 9
	}
	fmt.Println(sum)
	return true
}
```

### 思路二
对于思路一，当我们试图让n = 19时，我们发现由于数据过大，产生了溢出，因此我们因该解决溢出，如果要求溢出后的结果打印出取值范围内的最大值，其中最大值为了提高效率，我们仍然采用之前学过的位操作符左移，判断是否移除，我们可以采用字符串长度的比较，我们可以对思路一修改代码如下;
```go
func printMaxN_two(n int64) bool {
	if n <= 0 {
		fmt.Println(0)
		return false
	}

	// 求出int64的最大值 max := math.MaxInt64
	// 求最大值用左移操作符，提升效率，64位成2^62
	max := int64(1<<63 - 1)

	// 求出ma下对应的最大位数,先将max转换成sting
	smax := strconv.FormatInt(max, 10)
	maxn := len(smax)

	// 溢出时打印最大值
	if n >= int64(maxn) {
		// 将发生溢出
		fmt.Println(smax)
		return false
	}

	// 未溢出正常计算
	sum := int64(9)
	for i := int64(0); i < n-1; i++ {
		sum = sum*10 + 9
	}
	fmt.Println(sum)
	return true
}

```

### 思路三
如果题目要求溢出情况仍然求出最大值呢？这是就需要我们用字符串进行输出，可以发现，无论是否溢出，都可用转换成字符串的思路，n位最大的数，即包含了n个9的字符串，因此仍然采用该循环的方式
```go
func printMaxN_three(n int64) bool {
	// 判断n是否未负数
	if n <= 0 {
		fmt.Println(0)
		return false
	}

	// 这种情况无论是否溢出都进行打印字符串形式
	str := ""
	for i := int64(0); i < n; i++ {
		str += "9"
	}
	fmt.Println(str)
	return true
}
```

### 思路四
基于上述的思路，基本都会用到循环，我们知道任何循环都可用递归的形式，递归结束的条件时n == 0,否则就在原始字符串的基础上追加9，代码如下：
```go
func printMaxN_four(n uint64) (str string) {
	if n <= 0 {
		return ""
	} else {
		str += "9"
		n = n - 1
	}
	return str + printMaxN_four(n)
}
```

### 思路五
对于上述的思路，有一个共性就是只能打印出最大值，如果题目要求在打印最大值的基础上模拟数字的加法来计算呢？即我们在输出最大值的时候，控制台会依次打印出1，2，3，4，5.....
对于这个思路我们不防换一种思路求解，对于n位数中的每一位，都是对1~9的全排列，因此我们可以利用全排列+递归的解法，之前的文章已经阐述过相关全排列，这里我们只需要打印出n位数的每一位的全排列的组合即可，那么我们需要一个数组存放我们的n个数，初始化0,因此也就要求我们在打印的时候前面的0不需要进行打印，接着来看具体步骤：

1、初始化数组，从下标0开始进入递归
2、打印的条件为下标等于n
3、当打印的数组中为0的部分不打印
代码如下：
```go
func Print1ToMaxOfDigits(n int) {
	if n <= 0 {
		return
	}
	number := make([]int, n)
	fmt.Println(number)
	for i := 0; i < 10; i++ {
		number[0] = i
		print1ToMaxOfDigitsRecursively(number, n, 0)
	}

}

func print1ToMaxOfDigitsRecursively(number []int, length int, index int) {
	if index == length-1 {
		printNumber(number)
		return
	}

	for i := 0; i < 10; i++ {
		number[index+1] = i
		print1ToMaxOfDigitsRecursively(number, length, index+1)
	}
}

func printNumber(number []int) {
	var isBeginning0 = true
	length := len(number)
	for i := 0; i < length; i++ {
		if isBeginning0 && number[i] != 0 {
			isBeginning0 = false
		}

		if !isBeginning0 {
			fmt.Printf("%d", number[i])
			if i == length-1 {
				fmt.Println()
			}
		}
	}
}

```
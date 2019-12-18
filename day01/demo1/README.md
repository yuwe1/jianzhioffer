## 数组
数组是一块连续的内存并按照顺序存储数据，使用数组必须分配内存，因此数组的空间效率差，经常会出现空闲的区域没有得到充分利用。
数组的的内存连续，根据下标在O(1)时间读/写任何元素，时间效率高。

## 面试题——二维数组的查找
在一个二维数组中，每一行都按照从左到右的递增顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数输入这样的一个二维
数组和整数，判断该整数是否在该二维数组中。如：
```txt
1  2  8  9
2  4  9  12
4  7  10 13
6  8  11 15
```  
## 解决思路
根据数组已经排好的序列，我们应该在比较的过程中，不断缩小比较的区域，通常的我们的思路是从1开始比较，我们可以尝试一下，比如
我们找7这个数字，1比7小，因此，7应该在1的下面或右边，这个时候出现了问题，我们是因该比较右边的还是下面呢？因此这个思路不行，
那如果我们从顶点9开始呢，发现9比7大，那7一定在9的右边，因此我们可以比较8，发现仍然比较大，因此我们比较2，发现比7小，那么
一定在下面，因此开始比较4，然后比较7，最终找到。很明显这个思路比较容易。
## 找规律
通过以上的思路我们可以发现一下规律
```go
a[i][j] > num 时，j--
a[i][j] < num时,i++
```
## 代码如下
```go
type S1 []int
type S2 []S1

func FindNum(data S2, row, col, num int) bool {
	if len(data) == 0 || row <= 0 || col <= 0 {
		return false
	}
	//
	i := 0
	j := col - 1
	for {
		if i < 0 || j < 0 {
			break
		}
		if data[i][j] > num {
			j--
		} else if data[i][j] < num {
			i++
		} else {
			return true
		}
	}
	return false
}
```
在上面的代码中我们需要注意的是，因为在go语言中无法对[][]int传参是，必须保证容量和传入的相同，因此这里我们借助下面的方式完成二维数组的创建。
```go
type S1 []int
type S2 []S1
```
## 测试案例
demo1是我传入的自定义的包
```
func TestFindNum(t *testing.T) {
	// 定义一个完整数组
	// var data [][4]int
	// data = make([][4]int, 4)
	// fmt.Println(data)[[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]
	// 初始化二维数组
	fmt.Println("测试用例1******")
	a := demo1.S2{
		demo1.S1{1, 2, 8, 9},
		demo1.S1{2, 4, 9, 12},
		demo1.S1{4, 7, 10, 13},
		demo1.S1{6, 8, 11, 15},
	}
	num1, num2 := 8, 3
	if demo1.FindNum(a, 4, 4, num1) == true && demo1.FindNum(a, 4, 4, num2) == false {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
	fmt.Println("测试用例2********")
	a = demo1.S2{}
	if demo1.FindNum(a, 4, 4, num1) == false {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}
```

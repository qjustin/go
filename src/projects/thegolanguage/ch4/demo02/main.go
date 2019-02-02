package main

import (
	"fmt"
)

/*
4.2. Slice (切片)
1. 变长的序列(动态数组), 元素类型相同
2. slice的语法和数组很像，只是没有固定长度而已
3. 切片操作：
	a. s[i:j]，其中0 ≤ i≤ j≤ cap(s)，用于创建一个新的slice，引用s的从第i个元素开始到第j-1个元素的子序列。
	b. 如果i位置的索引被省略的话将使用0代替
	c. 如果j位置的索引被省略的话将使用len(s)代替
	d. 如果j位置的索引被省略的话将使用len(s)代替
	e. months[1:13]切片操作将引用全部有效的月份 等价于 months[1:]
	f. months[:]切片操作则是引用整个数组

slic底层：
slice的底层确实引用一个数组对象。

一个slice由三个部分构成：指针、长度和容量。
指针：指向第一个slice元素对应的底层数组元素的地址，
长度：对应slice中元素的数目；长度不能超过容量。
容量：一般是从slice的开始位置到底层数据的结尾位置。

要注意的是slice的第一个元素并不一定就是数组的第一个元素，
内置的len和cap函数分别返回slice的长度和容量。

slice 不能使用==操作符来判断两个slice是否含有全部相等元素，但数组可以
安全的做法是直接禁止slice之间的比较操作。

slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断。
因为，也有非nil值的slice的长度和容量也是0的，例如[]int{}或make([]int, 3)[3:]。

内置的make函数创建一个指定元素类型、长度和容量的slice。容量部分可以省略，在这种情况下，容量将等于长度。
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]

总结：
		一定要注意：
		a := [...]int{0, 1, 2, 3, 4, 5}
		与
		s := []int{0, 1, 2, 3, 4, 5}
		之间的差异，前者是数组，后者是切片

		创建数组：如果“[]”中是 “数字” 或 “...” 定义的就是数组
		创建切片：如果“[]”里什么也没有或 “:” 定义的就是切片

		创建切片：
		s = []int{0, 1, 2, 3, 4, 5}
		s := make([]string, 0, len(ages))

*/

// 定义一个数组,注意这里是数组
var months = [...]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func main() {
	// 下面是对months数组进行切片，
	// 记住切片不是复制一份新的数组，切片由指针、长度和容量组成
	Q1 := months[1:4]
	Q2 := months[4:7]
	Q3 := months[7:10]
	Q4 := months[10:]
	summer := months[6:9]
	fmt.Println(Q1)
	fmt.Println(Q2) // ["April" "May" "June"]
	fmt.Println(Q3)
	fmt.Println(Q4)
	fmt.Println(summer) // ["June" "July" "August"]

	// 如果切片操作超出cap(s)的上限将导致一个panic异常，
	// 但是超出len(s)则是意味着扩展了slice，因为新slice的长度会变大：

	// 超出原有summer切片的cap(s)，则这里抛出异常
	//fmt.Println(summer[:20]) // panic: out of range

	// 超出原有summer切片的len(s)，则自动扩展
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"

	// arr 是数组不是切片
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr)
	// reverse修改副本，不影响参数arr本身
	reverse(arr)
	fmt.Println(arr)

	// 这里定义的是一个切片
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr2)
	// reverse2 直接修改底层数组
	reverse2(arr2)
	// 试图将数组传递给reverse2，编译将报错
	// reverse2(arr)
	fmt.Println(arr2)
}

// reverse 接收的是数组，数组传递将产生副本，reverse修改副本，不影响参数本身
func reverse(s [8]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// reverse2 接收的是切片, 对切片的操作直接影响到底层数组
func reverse2(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

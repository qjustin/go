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
*/


// 定义一个数组
var months = [...]string{
	1: "January", 
	2: "February",
	3: "March",
	4: "April",
	5: "May",
	6: "June",
	7: "July",
	8: "August",
	9: "September",
	10: "October",
	11: "November",
	12: "December",
}

func main()  {
	// 下面是对months数组进行切片，
	// 记住切片不是复制一份新的数组，切片由指针、长度和容量组成
	
	Q1 := months[1:4]
	Q2 := months[4:7]
	Q3 := months[7:10]
	Q4 := months[10:]
	summer := months[6:9]
	fmt.Println(Q1)
	fmt.Println(Q2)	// ["April" "May" "June"]
	fmt.Println(Q3)
	fmt.Println(Q4)
	fmt.Println(summer) // ["June" "July" "August"]


	// 如果切片操作超出cap(s)的上限将导致一个panic异常，
	// 但是超出len(s)则是意味着扩展了slice，因为新slice的长度会变大：

	// 超出原有summer切片的cap(s)，则这里抛出异常
	fmt.Println(summer[:20]) // panic: out of range
	// 超出原有summer切片的len(s)，则自动扩展
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"
}
package main

import "fmt"
type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)
func main() {
	// 数组赋值初始化
	var arr1 [1]int
	arr1[0] = 1

	var arr2 [3]int = [3]int{1,2,3}

	arr3 := [3]int{1,2,3}

	arr4 := [...]int{1,2,3}

	fmt.Printf("%d 的类型是: %T\n", arr1, arr1)
	fmt.Printf("%d 的类型是: %T\n", arr2, arr2)
	fmt.Printf("%d 的类型是: %T\n", arr3, arr3)
	fmt.Printf("%d 的类型是: %T\n", arr4, arr4)

	// 如果你觉得每次写 [3]int 有点麻烦，你可以为 [3]int 定义一个类型字面量，也就是别名类型。
	// 使用 type 关键字可以定义一个类型字面量，后面只要你想定义一个容器大小为3，元素类型为int的数组 ，都可以使用这个别名类型。
	type arr5 [3]int
	myarr := arr5{1,2,3}
	fmt.Printf("%d 的类型是: %T", myarr, myarr)
	// 给下标定义一个别名
	e := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, e[RMB])
	// 偷懒写法
	arr6:=[4]int{3:1}
	fmt.Printf("%d 的类型是: %T\n", arr6, arr6)

	// 切片
	// 与数组不同的是，无法通过切片类型来确定其值的长度。每个切片值都会将数组作为其底层数据结构， 切片是对数组的一个连续片段的引用，
	// 所以切片是一个引用类型，这个片段可以是整个数组， 也可以是由起始和终止索引标识的一些项的子集，需要注意的是，终止索引标识的项不
	// 包括在切片内（意思是这是个左闭右开的区间）
	// 例子已经展示：myarr[0:2]，0是索引起始值，2是索引终止值，区间左闭右开
	myarr = [...]int{1, 2, 3}
	fmt.Printf("%d 的类型是: %T", myarr[0:2], myarr[0:2])

	// 创建数组：如果“[]”中是 “数字” 或 “...” 定义的就是数组
	// 创建切片：如果“[]”里什么也没有或 “:” 定义的就是切片

	// 切片的构造，有四种方式

	// 方式1： 对数组进行片段截取
	// 如下这段代码所示，切片从索引值 1 开始，到原数组终止索引值5，中间还可以容纳4个元素，所以容量为 4，
	// 但是由于我们切片的时候只要求取到索引值2 （3-1），所以当我们对这个切片进行打印时，并不会打印索引值3，4，5 对应的元素值。
	myarr3 := [5]int{1,2,3,4,5}
	fmt.Printf("myarr 的长度为：%d，容量为：%d\n", len(myarr3), cap(myarr3))

	mysli := myarr[1:3]
	fmt.Printf("mysli 的长度为：%d，容量为：%d\n", len(mysli), cap(mysli))
	fmt.Println(mysli)

	// 方式2： 从头声明赋值
	// 声明字符串切片
	//var strList []string

	// 声明整型切片
	//var numList []int

	// 声明一个空切片
	var numListEmpty = []int{}

	// 方式3： 使用 make 函数构造，make 函数的格式：make( []Type, size, cap )
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
	fmt.Println(cap(a), cap(b))

	// 方式4： 使用和数组一样，偷懒的方法
	c := []int{4:2}
	fmt.Println(c)
	fmt.Println(len(c), cap(c))

	// 由于 切片是引用类型，所以你不对它进行赋值的话，它的零值（默认值）是 nil
	fmt.Println(nil == numListEmpty)
}
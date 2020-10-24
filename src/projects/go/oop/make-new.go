package main

import "fmt"
/*
new 函数
new 只能传递一个参数，该参数为一个任意类型，可以是Go语言内建的类型，也可以是你自定义的类型
new 函数到底做了哪些事呢
分配内存
设置零值
返回指针（重要）
 */

/**
make 函数
内建函数 make 用来为 slice，map 或 chan 类型（注意：也只能用在这三种类型上）分配内存和初始化一个对象
make 返回类型的本身而不是指针，而返回值也依赖于具体传入的类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了
注意，因为这三种类型是引用类型，所以必须得初始化（size和cap），但是不是置为零值，这个和new是不一样的。
 */

/**
总结
new：为所有的类型分配内存，并初始化为零值，返回指针。
make：只能为 slice，map，chan 分配内存，并初始化，返回的是类型。
另外，目前来看 new 函数并不常用，大家更喜欢使用短语句声明的方式。
 */

func main() {
	num := new(int)
	fmt.Println(*num) //打印零值：0

	// new 一个自定义类型
	//s := new(Student)
	//s.name = "wangbm"



	//切片
	a := make([]int, 2, 10)
	fmt.Println(a)
	// 字典
	b := make(map[string]int)
	fmt.Println(b)
	// 通道
	c := make(chan int, 10)
	fmt.Println(c)
}
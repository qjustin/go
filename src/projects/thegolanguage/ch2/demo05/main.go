package main

import "fmt"

/*
	2.3.3. new函数
	另一个创建变量的方法是调用内建的new函数。表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T。
*/
func main() {
	// p, *int 类型, 指向匿名的 int 变量
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	// 每次调用new函数都是返回一个新的变量的地址，因此下面两个地址是不同的：
	a := new(int)
	b := new(int)
	fmt.Println(a == b)

}

// 用new创建变量和普通变量声明语句方式创建变量没有什么区别，除了不需要声明一个临时变量的名字外，我们还可以在表达式中使用new(T)。
func newInt() *int {
	return new(int)
}

func newIntByAnd() *int {
	var dummy int
	return &dummy
}

// 由于new只是一个预定义的函数，它并不是一个关键字，因此我们可以将new名字重新定义为别的类型。
func delta(old, new int) int {
	return new - old
}

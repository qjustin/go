package main

import "fmt"

/*
	2.6.2. 包的初始化, 作用域

	首先是解决包级变量的依赖顺序，然后按照包级变量声明出现的顺序依次初始化

	包含有多个.go文件，根据文件名排序，之后依次编译

	init() 函数：
	init() 类似于构造函数
	每个文件都可以包含多个init初始化函数，init初始化函数除了不能被调用或引用外，其他行为和普通函数类似。在程序开始执行时按照它们声明的顺序被自动调用。
	每个包只会被初始化一次。初始化工作是自下而上进行的，main包最后被初始化。

	作用域：
	一个程序可能包含多个同名的声明，只要它们在不同的词法域就没有关系。例如，你可以声明一个局部变量，和包级的变量同名。
	当编译器遇到一个名字引用时，它会对其定义进行查找，查找过程从最内层的词法域向全局的作用域进行。如果查找失败，则报告“未声明的名字”这样的错误。
*/

var a = b + c // a 第三个初始化, 为 3
var b = f()   // b 第二个初始化, 为 2, 通过调用 f (依赖c)
var c = 1     // c 第一个初始化, 为 1

func main() {
	fmt.Printf("%d, %d, %d", a, b, c)

	// 在函数中词法域可以深度嵌套，因此内部的一个声明可能屏蔽外部的声明。
	// 下面的代码有三个不同的变量x，因为它们是定义在不同的词法域
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}

	// 第二个if语句嵌套在第一个内部，因此第一个if语句条件初始化词法域声明的变量在第二个if中也可以访问。
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}

	// compile error: x and y are not visible here
	//fmt.Println(x, y)
}

func f() int {
	return c + 1
}

func g(x int) int {
	return x + 2
}

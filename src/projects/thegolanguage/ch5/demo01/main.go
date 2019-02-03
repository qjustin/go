package main

import (
	"fmt"
)

/*
5.1. 函数声明

函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。
函数的签名：函数的类型被称为函数的签名。如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型或签名。
			两个函数得参数列表，返回值列表，参数个数，类型一一对应，那么这两个函数具有相同的签名

实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实参包括引用类型，
如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用被修改。

// 你可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以Go实现的。这样的声明定义了函数签名。
// func Sin(x float64) float

*/

func main() {
	fmt.Println("Hello world")

	/*
		 一组形参或返回值有相同的类型，我们不必为每个形参都写出参数类型
		func f(i, j, k int, s, t string)                 {  }
		等价于
		func f(i int, j int, k int,  s string, t string) {  }
	*/

	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"

}

// 每一次函数调用都必须按照声明顺序为所有参数提供实参（参数值）。在函数调用时，Go语言没有默认参数值，
//  add, sub, first, zero 四个函数具有相同的签名，_符号可以强调某个参数未被使用
func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

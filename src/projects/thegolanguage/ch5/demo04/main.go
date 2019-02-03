package main

import (
	"fmt"
	"strings"
)

/*
5.5. 函数值

1. 在Go中，函数被看作第一类值（first-class values）：函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。对函数值（function value）的调用类似函数调用。
2. 函数的类型是否相同，由函数的签名决定 square， negative 函数的签名是一样的，因此可以赋值，但product于square和negative签名不同
3. 函数类型的零值是nil。调用值为nil的函数值会引起panic错误：
4. 函数值可以与nil比较：
*/

func main() {
	fmt.Println("Hello world")

	f := square
	fmt.Println(f(3)) // "9" 对函数值（function value）的调用类似函数调用

	f = negative
	fmt.Println(f(3))     // "-3"
	fmt.Printf("%T\n", f) // "func(int) int"
	// 函数的类型是否相同，由函数的签名决定 square， negative 函数的签名是一样的，因此可以赋值，但product于square和negative签名不同
	// 会产生一个编译错误
	//f = product // compile error: can't assign func(int, int) int to func(int) int

	// 此处f的值为nil, 会引起panic错误
	//f(3)

	//函数值可以与nil比较：
	if function != nil {
		function(3)
	}

	//函数值使得我们不仅仅可以通过数据来参数化函数，亦可通过行为。即--函数作为参数
	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
}

var function func(int) int

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func add1(r rune) rune { return r + 1 }

// forEachNode针对每个结点x，都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前，pre被调用
// 遍历孩子结点之后，post被调用
// func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
// 	if pre != nil {
// 		pre(n)
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		forEachNode(c, pre, post)
// 	}
// 	if post != nil {
// 		post(n)
// 	}
// }

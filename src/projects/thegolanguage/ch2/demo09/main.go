package main

import (
	"fmt"
)

/*
	2.4 赋值
	2.4.1. 元组赋值：它允许同时更新多个变量的值。在赋值之前，赋值语句右边的所有表达式将会先进行求值，然后再统一更新左边对应变量的值。
	2.4.2. 可赋值性: 赋值语句是显式的赋值形式，但是程序中还有很多地方会发生隐式的赋值行为：函数调用会隐式地将调用参数的值赋值给函数的
			参数变量，一个返回语句会隐式地将返回操作的值赋值给结果变量，一个复合类型的字面量（§4.2）也会产生赋值行为。
*/

func main() {
	scale := 2
	// 命名变量的赋值
	var x = 1

	p := new(int)
	// 通过指针间接赋值
	*p = 2

	var count [2]int
	count[0] = 1
	count[1] = 2
	// 数组、slice或map的元素赋值
	count[x] = count[x] * scale
	count[x] *= scale
	v := 1
	v++
	v--

	fmt.Printf("%d, %d, %d, %d, %d", x, *p, count[x], v, scale)

	var y = 0
	// 元组赋值：它允许同时更新多个变量的值。在赋值之前，赋值语句右边的所有表达式将会先进行求值，然后再统一更新左边对应变量的值。
	// 交换x，y 得值
	x, y = y, x
	count[x], count[y] = count[y], count[x]
	var i, j, k int
	i, j, k = 2, 3, 5
	// 但如果表达式太复杂的话，应该尽量避免过度使用元组赋值；因为每个变量单独赋值语句的写法可读性会更好。
	fmt.Printf("%d, %d, %d, %d, %d, %d, %d, %d", x, *p, count[x], v, scale, i, j, k)

	// 通常，这类函数会用额外的返回值来表达某种错误类型，例如os.Open是用额外的返回值返回一个error类型的错误，
	// 还有一些是用来返回布尔值，通常被称为ok。在稍后我们将看到的三个操作都是类似的用法。如果map查找（§4.3）、
	// 类型断言（§7.10）或通道接收（§8.4.2）出现在赋值语句的右边，它们都可能会产生两个结果，有一个额外的布尔结果表示操作是否成功：
	var m = make(map[string]int)
	var ch = make(chan string)

	v1, ok1 := m["hello"] // map lookup
	v3, ok3 := <-ch       // channel receive

	v1 = m["hello"] // map查找，失败时返回零值
	v3 = <-ch       // 管道接收，失败时返回零值（阻塞不算是失败）

	_, ok1 = m["hello"]   // map返回2个值
	_, ok1 = m[""], false // map返回1个值
	_ = m[""]             // map返回1个值

	medals := []string{"gold", "silver", "bronze"}
	fmt.Printf("%d, %d, %d, %d, %d, %d, %d, %d, %d, %t, %s, %t, %s", x, *p, count[x], v, scale, i, j, k, v1, ok1, v3, ok3, medals[0])
}

// 计算两个整数值的的最大公约数（GCD）
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

//计算斐波纳契数列（Fibonacci）的第N个数
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

package main

import (
	"fmt"
)

/*

	4.4.1. 结构体字面值
 	1. 这个结构体包含两个可到处的成员，未导出的成员不能在包外赋值
	2. 结构体可以作为函数的参数和返回值
		记住：在Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。
			

 */
type Point struct {
	X, Y int
}



func main() {
	// 用结构体定义一个结构体变量
	var a Point
	// 结构体变量赋值
	// 方式1：这种写法，要求以结构体成员定义的顺序为每个结构体成员指定一个字面值。
	// 缺点是：结构体成员有细微的调整就可能导致上述代码不能编译。
	// 使用场景：只在定义结构体的包内部使用，或者是在较小的结构体中使用，
	// 这些结构体的成员排列比较规则。例如： color.RGBA{red, green, blue, alpha}
	b := Point{1, 2}
	// 方式2：因为提供了成员的名字，所以成员出现的顺序并不重要
	c := Point{X:3, Y:4}

	/*
	未导出的成员不能在包外赋值：

	package p
	type T struct{ a, b int } // a and b are not exported

	package q
	import "p"
	var _ = p.T{a: 1, b: 2} // compile error: can't reference a, b
	var _ = p.T{1, 2}       // compile error: can't reference a, b
	*/

	//因为结构体通常通过指针处理，可以用下面的写法来创建并初始化一个结构体变量，并返回结构体的地址：
	d := &Point{1, 2}
	fmt.Println("%d %d %d %d", a, b, c, d)
}

// 结构体可以作为函数的参数和返回值， 
// 这种传递方式p是一个副本，副本的修改不会影响参数的内容
// 这种效率低下
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// 这种传递方式p是一个指针，指向参数的实际地址，直接修改参数内容
// 这种方式效率高
func ScalePoint(p *Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}
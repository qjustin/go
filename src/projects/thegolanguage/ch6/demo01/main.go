package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

/*
6.1. 方法声明
方法声明：在函数声明时，在其名字之前放上一个变量，即是一个方法。
这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义
了一个独占的方法。
*/
func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call
}



// 定义一个传统的函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

/* 
	定义一个方法
	附加的参数p，叫做方法的接收器（receiver），早期的面向对象语言留下的遗产将
	调用一个方法称为“向一个对象发送消息”。在Go语言中，我们并不会像其它语言那样
	用this或者self作为接收器；我们可以任意的选择接收器的名字。由于接收器的名字
	经常会被使用到，所以保持其在方法间传递时的一致性和简短性是不错的主意。这里的
	建议是可以使用其类型的第一个字母，比如这里使用了Point的首字母p。
*/
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}


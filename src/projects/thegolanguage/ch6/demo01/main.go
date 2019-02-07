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
/*
	第一个Distance的调用实际上用的是包级别的函数geometry.Distance
	第二个则是使用刚刚声明的Point，调用的是Point类下声明的Point.Distance方法。
*/
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call



	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
/*
	对于一个给定的类型，其内部的方法都必须有唯一的方法名，但是不同的类型却可以有同样
	的方法名，比如我们这里Point和Path就都有Distance这个名字的方法；所以我们没有必
	要非在方法名之前加类型名来消除歧义，比如PathDistance。这里我们已经看到了方法比
	之函数的一些好处：方法名可以简短。当我们在包外调用的时候这种好处就会被放大，因为
	我们可以使用这个短名字，而可以省略掉包的名字
*/
// 	import "gopl.io/ch6/geometry"

// perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
// fmt.Println(geometry.PathDistance(perim)) // "12", standalone function
// fmt.Println(perim.Distance())             // "12", method of geometry.Path
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

/*
Path是一个命名的slice类型，而不是Point那样的struct类型，然而我们依然可以为它定义方法。在能够给任意类型定义方法这一点上，

Go和很多其它的面向对象的语言不太一样。因此在Go语言里，我们为一些简单的数值、
字符串、slice、map来定义一些附加行为很方便。我们可以给同一个包内的任意命名
类型定义方法，只要这个命名类型的底层类型不是指针或者interface。
（译注：这个例子里，底层类型是指[]Point这个slice，Path就是命名类型）
*/
type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

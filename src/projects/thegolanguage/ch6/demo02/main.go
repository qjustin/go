package main

import (
	"fmt"
)

type Point struct{ X, Y float64 }

/*
6.2. 基于指针对象的方法

1. 函数参数传递时值传递(拷贝一个副本)，对应到方法更新接收器的对象的方法，
当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法

2. 约定：在现实的程序里，一般会约定如果Point这个类有一个指针作为接收器的方法，
那么所有Point的方法都必须有一个指针接收器，即使是那些并不需要这个指针接
收器的函数。

3. 声明方法时，如果一个类型名本身是一个指针的话，是不允许其出现在接收器中
	type P *int
	func (P) f() {} //compile error: invalid receiver type

4. 调用方式4:如果接收器u是一个Point类型的变量，但是ScaleBy方法需要一个Point
	指针作为接收器，编译器会隐式地帮我们用&u去调用ScaleBy这个方法。这种简写
	的前提是u是一个Point类型的变量。这种简写方法只适用于“变量”，包括struct里
	的字段比如p.X，以及array和slice内的元素比如perim[0]。

译注： 作者这里说的比较绕，其实有两点：

1. 不管你的method的receiver是指针类型还是非指针类型，都是可以通过
	指针/非指针类型进行调用的，编译器会帮你做类型转换。
2. 在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的因素，
	第一方面是这个对象本身是不是特别大，如果声明为非指针变量时，调用会产生一次拷贝；
	第二方面是如果你用指针类型作为receiver，那么你一定要注意，这种指针类型指向的始终
		是一块内存地址，就算你对其进行了拷贝。熟悉C或者C++的人这里应该很快能明白。
*/

func main() {
	// ScaleBy 接收器类型时一个point 指针，
	// 调用方式1
	p := &Point {1, 2}
	p.ScaleByPointPtr(2.0)
	fmt.Println(p)
	// 调用方式2
	q := Point{2, 4}
	qptr := &q
	qptr.ScaleByPointPtr(2.0)
	fmt.Println(qptr) 
	// 调用方式3
	t := Point{6, 8}
	(&t).ScaleByPointPtr(2)
	fmt.Println(t) 
	// 调用方式4
	u := Point{6, 8} // 声明一个Point类型的变量u
	u.ScaleByPointPtr(2)
	fmt.Println(u) 
	// 我们不能通过一个无法取到地址的接收器来调用指针
	// 方法，比如临时变量的内存地址就无法获取得到
	// Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal
/*
	不管你的method的receiver是指针类型还是非指针类型，都是可以通过
	指针/非指针类型进行调用的，编译器会帮你做类型转换。
*/

a := Point{1, 2}

b := &Point{1, 2}

// 接收器实参是类型*T，形参是类型T。编译器会隐式地为我们解引用，取到指针指向的实际变量
// 接收器实参是类型T，但接收器形参是类型*T，这种情况下编译器会隐式地为我们取变量的地址
a.ScaleByPoint(2)  			//接收器实参是类型T，形参是类型T
a.ScaleByPointPtr(2)		//接收器实参是类型T，形参是类型*T  // implicit（&a）
(&a).ScaleByPoint(2)  		//接收器实参是类型*T，形参是类型T
(&a).ScaleByPointPtr(2)		//接收器实参是类型*T，形参是类型*T
//(*a).ScaleByPoint(2)  
//(*a).ScaleByPointPtr(2)
b.ScaleByPoint(2)    		//接收器实参是类型*T，形参是类型T // implicit (*b)
b.ScaleByPointPtr(2)		//接收器实参是类型*T，形参是类型*T
//(&b).ScaleByPoint(2)    
//(&b).ScaleByPointPtr(2)
(*b).ScaleByPoint(2)    	//接收器实参是类型*T，形参是类型T
(*b).ScaleByPointPtr(2)		//接收器实参是类型*T，形参是类型*T

Point{1, 2}.ScaleByPoint(2) //接收器实参是类型T，形参是类型T
(&Point{1, 2}).ScaleByPoint(2) //接收器实参是类型*T，形参是类型T
// Point{1, 2}.ScaleByPointPtr(2)
(&Point{1, 2}).ScaleByPointPtr(2) //接收器实参是类型*T，形参是类型*T

}

/*
	1. 方法ScaleBy 更新 接收器 p 的值，方法默认也是值传递，如果p对象太大
	   拷贝过程将影响运行效率，我们可以用其指针而不是对象来声明方法。

	   这个方法的名字是(*Point).ScaleBy。这里的括号是必须的；没有括号的
	   话这个表达式可能会被理解为*(Point.ScaleBy)
*/

func (p *Point) ScaleByPointPtr(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) ScaleByPoint(factor float64) {
	p.X *= factor
	p.Y *= factor
}

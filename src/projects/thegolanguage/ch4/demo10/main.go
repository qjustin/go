package main

import (
	"fmt"
)

/*
	4.4.3. 结构体嵌入和匿名成员
	1. 一个命名的结构体可以包含另一个结构体类型，
	2. 一个命名的结构体可以包含另一个结构体类型的匿名成员
	3. 因为匿名成员也有一个隐式的名字，因此不能同时包含两个类型相同的匿名成员，这会导致名字冲突。

	匿名成员：结构体中声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员。
	
	到目前为止，我们看到匿名成员特性只是对访问嵌套成员的点运算符提供了简短的语法糖。
	稍后，我们将会看到匿名成员并不要求是结构体类型；其实任何命名的类型都可以作为结构体
	的匿名成员。
*/

func main() {
	// 缺点：这样改动之后结构体类型变的清晰了，但是这种修改同时也导致了访问每个成员变得繁琐
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20

	// 现在使用匿名成员来改进上面的结构，我们可以直接访问叶子属性而不需要给出完整的路径
	// 注意：在右边的注释中给出的显式形式访问这些叶子成员的语法依然有效，
	var w1 Wheel1
	w1.X = 8      // equivalent to w1.Circle1.Point1.X = 8
	w1.Y = 8      // equivalent to w1.Circle1.Point1.Y = 8
	w1.Radius = 5 // equivalent to w1.Circle1.Radius1 = 5
	w1.Spokes = 20

	// 结构体嵌套的赋值
	// 结构体字面值并没有简短表示匿名成员的语法， 因此下面的语句都不能编译通过：
	// w2 = Wheel1{8, 8, 5, 20}                       // compile error: unknown fields
	// w2 = Wheel1{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields

	// 结构体字面值必须遵循形状类型声明时的结构，所以我们只能用下面的两种语法，它们彼此是等价的：
	w3 := Wheel1{Circle1{Point1{8, 8}, 5}, 20}
	w4 := Wheel1{
		Circle1: Circle1{
			Point1: Point1{
				X: 8, Y: 8,
			},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%d, %d, %d, %d", w, w1, w3, w4)
}

/*
	结构可以包含另一个结构，因此，点，圆形，轮子 声明的形式如下
	圆包含点，轮子包含园

	缺点：这样改动之后结构体类型变的清晰了，但是这种修改同时也导致了访问每个成员变得繁琐
*/
type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

/*
	现在使用匿名成员来改进上面的结构
*/
type Point1 struct {
	X, Y int
}

type Circle1 struct {
	Point1
	Radius int
}

type Wheel1 struct {
	Circle1
	Spokes int
}

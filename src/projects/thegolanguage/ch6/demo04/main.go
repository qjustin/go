package main

import (
	"fmt"
	"image/color"
	"math"
	"time"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

/*
6.4. 方法值和方法表达式

我们经常选择一个方法，并且在同一个表达式里执行，比如常见的p.Distance()
形式，实际上将其分成两步来执行也是可能的。p.Distance叫作“选择器”，选择
器会返回一个方法“值”->一个将方法（Point.Distance）绑定到特定接收器变
量的函数。这个函数可以不通过指定其接收器即可被调用；即调用时不需要指定接
收器（译注：因为已经在前文中指定过了），只要传入函数的参数即可
*/
func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance        // method value
	fmt.Println(distanceFromP(q))      // "5"
	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", sqrt(5)

	scaleP := p.ScaleBy // method value
	scaleP(2)           // p becomes (2, 4)
	scaleP(3)           //      then (6, 12)
	scaleP(10)          //      then (60, 120)

	r := new(Rocket)
	time.AfterFunc(10*time.Second, func() { r.Launch() })
	// 直接用方法“值”传入AfterFunc的话可以更为简短：(省掉了上面那个例子里的匿名函数)
	time.AfterFunc(10*time.Second, r.Launch)

	// 当调用一个方法时，与调用一个普通的函数相比，我们必须要用选择器（p.Distance）
	// 语法来指定方法的接收器。

	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)            // "{2 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"

	// 译注：这个Distance实际上是指定了Point对象为接收器的一个方法func (p Point) Distance()，
	// 但通过Point.Distance得到的函数需要比实际的Distance方法多一个参数，
	// 即其需要用第一个额外参数指定接收器，后面排列Distance方法的参数。
	// 看起来本书中函数和方法的区别是指有没有接收器，而不像其他语言那样是指有没有返回值。
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// 等价于：
func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}

func (p *ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}

type Rocket struct{}

func (r *Rocket) Launch() {}

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)

	}
}

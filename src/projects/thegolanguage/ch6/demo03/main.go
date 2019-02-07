package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

type ColoredPoint1 struct {
	*Point
	Color color.RGBA
}

/*
6.3. 通过嵌入结构体来扩展类型



*/
func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	/*
		可以把ColoredPoint类型当作接收器来调用Point里的方法，
		即使ColoredPoint里没有声明这些方法,因为ColoredPoint
		包含Point

		需要注意的是：ColoredPoint 与 Point 不是 is a的关系
		从Distance参数可以看出，需要指出明确的类型q.Point
	*/
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"

	/*
		在类型中内嵌的匿名字段也可能是一个命名类型的指针，这种情况下字段和方法会被
		间接地引入到当前的类型中（译注：访问需要通过该指针指向的对象去取）。添加这
		一层间接关系让我们可以共享通用的结构并动态地改变对象之间的关系。下面这个
		ColoredPoint的声明内嵌了一个*Point的指针。
	*/
	p1 := ColoredPoint1{&Point{1, 1}, red}
	q1 := ColoredPoint1{&Point{5, 4}, blue}
	fmt.Println(p1.Distance(*q1.Point)) // "5"
	q1.Point = p1.Point                 // p and q now share the same Point
	p1.ScaleBy(2)
	fmt.Println(*p1.Point, *q1.Point) // "{2 2} {2 2}"
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

var (
	mu      sync.Mutex // guards mapping
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

/*
下面这个版本在功能上是一致的，但将两个包级别的变量放在了cache这个struct一组内

我们给新的变量起了一个更具表达性的名字：cache。因为sync.Mutex字段也被嵌入到了
这个struct里，其Lock和Unlock方法也就都被引入到了这个匿名结构中了，这让我们能
够以一个简单明了的语法来对其进行加锁解锁操作
*/
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func LookupV2(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

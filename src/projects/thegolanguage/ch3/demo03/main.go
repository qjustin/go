package main

import (
	"fmt"
	"math"
)

/*
3.6.2. 无类型常量

多常量并没有一个明确的基础类型。编译器为这些没有明确基础类型的数字常量提供比基础类型更高精度的算术运算；你可以认为至少有256bit的运算精度。

Go有六种未明确类型的常量类型，分别是：无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。

无类型的常量不仅可以提供更高的运算精度，而且可以直接用于更多的表达式而不需要显式的类型转换。

0、0.0、0i、\u0000、false、"0" 虽然有着相同的常量值，但是它们分别对应无类型的整数、无类型的浮点数、无类型的复数、无类型的字符、无类型的布尔类型， 无类型的字符串类型等不同的常量类型。
*/
// 这里math.Pi 是一个无类型常量，因此不需要转换
var x float32 = math.Pi
var y float64 = math.Pi
var z complex128 = math.Pi

const (
	deadbeef = 0xdeadbeef        // untyped int with value 3735928559
	a1       = uint32(deadbeef)  // uint32 with value 3735928559
	b1       = float32(deadbeef) // float32 with value 3735928576 (rounded up)
	c1       = float64(deadbeef) // float64 with value 3735928559 (exact)
	// d1       = int32(deadbeef)   // compile error: constant overflows int32
	// e1       = float64(1e309)    // compile error: constant overflows float64
	// f1       = uint(-1)          // compile error: constant underflows uint
)

func main() {
	// ZiB和YiB的值已经超出任何Go语言中整数类型能表达的范围，但是它们依然是合法的常量，而且像下面的常量表达式依然有效
	// fmt.Println(YiB / ZiB) // "1024"

	// 这里由于math.Pi被确定为特定类型，因此需要转换
	const Pi64 float64 = math.Pi
	var x float32 = float32(Pi64)
	var y float64 = Pi64
	var z complex128 = complex128(Pi64)

	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f - 32))     // "0";   5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float

	// 只有常量可以是无类型的。当一个无类型的常量被赋值给一个变量的时候，就像下面的第一行语句，或者出现在有明确类型的变量声明的右边，如下面的其余三行语句，无类型的常量将会被隐式转换为对应的类型，如果转换合法的话。
	var o float64 = 3 + 0i // untyped complex -> float64
	o = 2                  // untyped integer -> float64
	o = 1e123              // untyped floating-point -> float64
	o = 'a'                // untyped rune ->

	// 等价于
	var p float64 = float64(3 + 0i)
	p = float64(2)
	p = float64(1e123)
	p = float64('a')

	// 对于一个没有显式类型的变量声明（包括简短变量声明），常量的形式将隐式决定变量的默认类型，就像下面的例子：
	q := 0      // untyped integer;        implicit int(0)
	r := '\000' // untyped rune;           implicit rune('\000')
	s := 0.0    // untyped floating-point; implicit float64(0.0)
	t := 0i     // untyped complex;        implicit complex128(0i)

	// Go语言不存在整型类似的不确定内存大小的浮点数和复数类型。
	// 如果要给变量一个不同的类型，我们必须显式地将无类型的常量转化为所需的类型，或给声明的变量指定明确的类型，
	var i = int8(0) // 等价于 	var i int8 = 0

	// 当尝试将这些无类型的常量转为一个接口值时（见第7章），这些默认类型将显得尤为重要，因为要靠它们明确接口对应的动态类型。
	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)

	fmt.Printf("%d, %d, %d, %d, %d, %d, %f, %f, %f, %f, %d", x, y, z, f, o, p, q, r, s, t, i)
}

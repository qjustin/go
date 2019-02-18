package tempconv

import "fmt"

/*
	2.5. 类型
	即使类型的底层数据结构相同，它们仍是两个不同的类型。

	每个类型T都有一个对应的类型转换操作T(x)，用于将x转为T类型
	如果T是指针类型，可能会需要用小括弧包装T，比如(*int)(0),
	转换的前提是：只有当两个类型的底层基础类型相同时，才允许这种转型操作

	将x类型转换成T类型的时候，x类型的值不变

*/
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func PrintTemp() {
	// 数值类型之间也可以相互转换
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	// 将一个浮点数转为整数将丢弃小数部分，BoilingC的值不会被改变
	boilingF := CToF(BoilingC)
	// 底层数据类型相同因此可以做减法
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	// 底层数据类型不同因此不能做算数运算， 这里报错
	// fmt.Printf("%g\n", boilingF-FreezingC) // compile error: type mismatch

	// 如果类型底层数据类型相同可以使用==和< 或者类型的底层数据类型与未命名的值类型相同也可以进行比较。
	// 总之一句话，相同类型才能比较，不管是自定义类型，还是值类型，只要底层类型与值类型的类型相同则可以比较。
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	//fmt.Println(c == f)          // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // "true"!

	// 许多类型都会定义一个String方法，因为当使用fmt包的打印方法时，将会优先使用该类型对应的String方法返回的结果打印
	d := FToC(212.0)
	//fmt.Println(d.String()) // "100°C"
	fmt.Printf("%v\n", d)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", d)   // "100°C"
	fmt.Println(d)          // "100°C"
	fmt.Printf("%g\n", d)   // "100"; does not call String
	fmt.Println(float64(d)) // "100"; does not call String
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

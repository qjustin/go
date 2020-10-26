package main

import "fmt"

func main() {
	// 所谓的静态类型（即 static type），就是变量声明的时候的类型。
	var age int     // int 是静态类型
	var name string // string 也是静态类型
	fmt.Println(age)
	fmt.Println(name)

	// 所谓的 动态类型（即 concrete type，也叫具体类型）是 程序运行时系统才能看见的类型。
	var i interface{} // 第一行：我们在给 i 声明了 interface{} 类型，所以 i 的静态类型就是 interface{}
	i = 18       // 第二行：当我们给变量 i 赋一个 int 类型的值时，它的静态类型还是 interface{}，这是不会变的，但是它的动态类型此时变成了 int 类型。
	i = "Go编程时光" // 第三行：当我们给变量 i 赋一个 string 类型的值时，它的静态类型还是 interface{}，它还是不会变，但是它的动态类型此时又变成了 string 类型。
	fmt.Println(i)
	// 不管是 i=18 ，还是 i="Go编程时光"，都是当程序运行到这里时，变量的类型，才发生了改变，这就是我们最开始所说的 动态类型是程序运行时系统才能看见的类型。

	// 3. 接口组成
	// 每个接口变量，实际上都是由一 pair 对（type 和 data）组合而成，pair 对中记录着实际变量的值和类型。
	var age1 int = 25 // 等价于 age1 := (int)(25)
	fmt.Println(age1)

	// 4. 接口细分
	// 第一种：iface，表示带有一组方法的接口。
	// 第二种：eface，表示不带有方法的接口

	// iface 类型指针指向的的数据结构包含 静态类型和动态类型两个指针，静态指向定义时的类型，动态指向赋值时的实例类型值（实例类型）
	// eface 类型指针直接指向实例的值（实例类型）

	// 5. 反射的必要性
	// 由于动态类型的存在，在一个函数中接收的参数的类型有可能无法预先知晓，此时我们就要对参数进行反射，然后根据不同的类型做不同的处理。
}

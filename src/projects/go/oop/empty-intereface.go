package main

import "fmt"

/**
空接口是特殊形式的接口类型，普通的接口都有方法，而空接口没有定义任何方法口，也因此，我们可以说所有类型都至少实现了空接口。
*/

func main() {
	// 1. 每一个接口都包含两个属性，一个是值，一个是类型。而对于空接口来说，这两者都是 nil，可以使用 fmt 来验证一下
	var i1 interface{}
	fmt.Printf("type: %T, value: %v", i1, i1)

	// 2. 如何使用空接口？
	// 2.1 第一，通常我们会直接使用 interface{} 作为类型声明一个实例，而这个实例可以承载任意类型的值。

	var i2 interface{}
	// 存 int 没有问题
	i2 = 1
	fmt.Println(i2)
	// 存字符串也没有问题
	i2 = "hello"
	fmt.Println(i2)
	// 存布尔值也没有问题
	i2 = false
	fmt.Println(i2)

	// 2.2 第二，如果想让你的函数可以接收任意类型的值 ，也可以使用空接口
	a := 10
	b := "hello"
	c := true

	myfunc1(a)
	myfunc1(b)
	myfunc1(c)

	myfunc2(a, b, c)

	// 2.3 第三，你也定义一个可以接收任意类型的 array、slice、map、strcut，例如这边定义一个切片
	any := make([]interface{}, 5)
	any[0] = 11
	any[1] = "hello world"
	any[2] = []int{11, 22, 33, 44}
	for _, value := range any {
		fmt.Println(value)
	}

	// 3. 空接口几个要注意的坑
	// 坑1：空接口可以承载任意值，但不代表任意类型就可以承接空接口类型的值
	// 从实现的角度看，任何类型的值都满足空接口。因此空接口类型可以保存任何值，也可以从空接口中取出原值。
	// 但要是你把一个空接口类型的对象，再赋值给一个固定类型（比如 int, string等类型）的对象赋值，是会报错的。
	// 示例：1
	// 声明a变量, 类型int, 初始值为1
	// var a int = 1
	// 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
	// var i interface{} = a
	// 声明b变量, 尝试赋值i
	// var b int = i
	// 这个报错，它就好比可以放进行礼箱的东西，肯定能放到集装箱里，但是反过来，能放到集装箱的东西就不一定能放到行礼箱了，在 Go 里就直接禁止了这种反向操作。（声明：底层原理肯定还另有其因，但对于新手来说，这样解释也许会容易理解一些。）

	// 坑2：：当空接口承载数组和切片后，该对象无法再进行切片
	// 示例：2
	//sli := []int{2, 3, 5, 7, 11, 13}
	//var i interface{}
	//i = sli
	//g := i[1:3]
	//fmt.Println(g)

	// 坑3：当你使用空接口来接收任意类型的参数时，它的静态类型是 interface{}，但动态类型（是 int，string 还是其他类型）我们并不知道，因此需要使用类型断言。
	a1 := 10
	b1 := "hello"
	myfunc(a1)
	myfunc(b1)
}
func myfunc(i interface{})  {
	switch i.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}

func myfunc1(iface interface{}) {
	fmt.Println(iface)
}

func myfunc2(ifaces ...interface{}) {
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}

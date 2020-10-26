package main

import "fmt"

type Phone1 interface {
	call()
}

type iPhone struct {
	name string
}

func (phone iPhone)call()  {
	fmt.Println("Hello, iPhone.")
}

func (phone iPhone)send_wechat()  {
	fmt.Println("Hello, Wechat.")
}
// ---------------------
func printType(i interface{})  {

	switch i.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}


func main() {
	// 接口的三个“潜规则”
	// 1. 对方法的调用限制:接口是一组固定的方法集，由于静态类型的限制，接口变量有时仅能调用其中特定的一些方法。
	// 当静态类型时一个接口类型时，只能调用接口中的方法。不能调用动态类型中的方法
	var phone Phone
	phone = iPhone{name:"ming's iphone"}
	phone.call()
	// Phone1 j接口不包含send_wechat() 因为我们的phone对象显示声明为 Phone1 接口类型，因此 phone调用的方法会受到此接口的限制。
	// phone.send_wechat()

	// 那么如何让 phone 可以调用 send_wechat 方法呢？案是可以不显示的声明为 Phone接口类型 ，
	// 但要清楚 phone 对象实际上是隐式的实现了 Phone 接口，如此一来，方法的调用就不会受到接口类型的约束。
	var phone1 iPhone
	phone1 = iPhone{name:"ming's iphone"}
	phone1.call()
	phone1.send_wechat()

	// 1. 调用函数时的隐式转换： Go 语言中的函数调用都是值传递的，变量会在方法调用前进行类型转换。
	a := 10
	printType(a) // 一切都很正常

	// 但是如果你把函数内的内容搬到到外面来， 就会有意想不到的结果，居然报错了。
	// 原因其实很简单。
	//
	//当一个函数接口 interface{} 空接口类型时，我们说它可以接收什么任意类型的参数（江湖上称之为无招胜有招）。
	//当你使用这种写法时，Go 会默默地为我们做一件事，就是把传入函数的参数值（注意：Go 语言中的函数调用都是值传递的）的类型隐式的转换成 interface{} 类型。
	a1 := 10

	switch a1.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}


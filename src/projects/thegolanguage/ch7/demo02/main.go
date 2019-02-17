package main

import (
	"io"
	"os"
	"bytes"
	"fmt"
)

/*
	7.3. 实现接口的条件

	1. 一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。(接口与实现只依赖于判断两个类型的方法)
	2. 接口指定的规则非常简单：表达一个类型属于某个接口只要这个类型实现这个接口
	3. 接口类型封装和隐藏具体类型和它的值。即使具体类型有其它的方法，也只有接口类型暴露出来的方法会被调用到

	6.2章中，对于些方法的接收者是类型T本身然而另一些则是一个*T的指针，那么
	在T类型的参数上调用一个*T的方法是合法的，只要这个参数是一个变量，编译器
	隐式的获取了它的地址

	type IntSet struct {  }
	func (*IntSet) String() string
	var _ = IntSet{}.String() // compile error: String requires *IntSet receiver
	但是我们可以在一个IntSet变量上调用这个方法：
	var s IntSet
	var _ = s.String() // OK: s is a variable and &s has a String method

	因此：T类型的值不拥有所有*T指针的方法，这样它就可能只实现了更少的接口。

	空接口类型：
	interface{}被称为空接口类型是不可或缺的。因为空接口类型对实现它的类型没有要求，
	所以我们可以将任意一个值赋给空接口类型。
	
	不像基于类的语言，他们一个类实现的接口集合需要进行显式的定义，在Go语言中我们
	可以在需要的时候定义一个新的抽象或者特定特点的组，而不需要修改具体类型的定义。
	当具体的类型来自不同的作者时这种方式会特别有用。当然也确实没有必要在具体的类型
	中指出这些共性。
	*/

func main() {
	//3. 接口类型封装和隐藏具体类型和它的值。即使具体类型有其它的方法，也只有接口类型暴露出来的方法会被调用到
	os.Stdout.Write([]byte("hello")) // OK: *os.File has Write method
	os.Stdout.Close()                // OK: *os.File has Close method

	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello")) // OK: io.Writer has Write method
	// w.Close()                // compile error: io.Writer lacks Close method

	// 空接口类型
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	fmt.Println(any)
}
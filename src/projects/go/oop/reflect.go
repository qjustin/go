package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 2. 两种类型：reflect.Type 和 reflect.Value 是整个反射的核心
	// reflect.Type 是以一个接口的形式存在的
	// reflect.Value 是以一个结构体的形式存在

	// 为了区分反射前后的变量值类型，我将反射前环境称为 真实世界，而将反射后的环境称为 反射世界。
	// 真实世界里，type 和 data 是合并在一起组成 接口变量的
	// 反射世界里，type 和 data 却是分开的，他们分别由 reflect.Type 和 reflect.Value 来表现

	// 反射可以将接口类型变量 转换为“反射类型对象”；
	// 反射可以将 “反射类型对象” 转换为 接口类型变量；
	// 如果要修改 “反射类型对象” 其类型必须是可写的；

	// reflect.TypeOf(i) ：获得接口值的类型 	Type Object
	// reflect.ValueOf(i)：获得接口值的值		Value Object
}

// 反射可以将接口类型变量 转换为“反射类型对象”；
func demo01() {
	var age interface{} = 25 // 上面我们定义的 age 不是 int 类型的吗？第一法则里怎么会说是接口类型的呢？
	// 由于 TypeOf 和 ValueOf 两个函数接收的是 interface{} 空接口类型，而 Go 语言函数都是值传递，因此Go语言会将我们的类型隐式地转换成接口类型。
	// TypeOf returns the reflection Type of the value in the interface{}.TypeOf returns nil.
	// func TypeOf(i interface{}) Type

	// ValueOf returns a new Value initialized to the concrete value stored in the interface i. ValueOf(nil) returns the zero Value.
	// func ValueOf(i interface{}) Value

	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)

	// 原始接口变量的类型为 int，值为 25
	// 从接口变量到反射对象：Type对象的类型为 *reflect.rtype
	// 从接口变量到反射对象：Value对象的类型为 reflect.Value

}

/*
总结：

第一定律：反射可以将接口类型变量 转换为“反射类型对象”
我们可以通过反射将一个接口类型变量（接口类型变量可以接收任何类型）转换为“反射类型对象” 即：reflect.Type 对象和 reflect.Value 对象

反射类型有两个：reflect.Type , reflect.Value。
reflect类型提供了两个方法，用于从“接口类型变量”中提取 reflect.Type 对象和 reflect.Value 对象
 func TypeOf(i interface{}) Type ：获得接口值的类型 	Type Object
 func ValueOf(i interface{}) Value ：获得接口值的值	Value Object

第二定律：反射可以将 “反射类型对象” 转换为 接口类型变量；
reflect.Value 对象的 Interface() 方法 可以还原接口类型变量
通过源码可知， reflect.Value 的结构体会接收 Interface 方法，返回了一个 interface{} 类型的变量
（注意：只有 Value 才能逆向转换，而 Type 则不行，这也很容易理解，如果 Type 能逆向，那么逆向成什么呢？）

第三定律：如果要修改 “反射类型对象” 其类型必须是可写的
不是接收变量指针创建的反射对象，是不具备『可写性』的
是否具备『可写性』，可使用 CanSet() 来获取得知
对不具备『可写性』的对象进行修改，是没有意义的，也认为是不合法的，因此会报错。
 */
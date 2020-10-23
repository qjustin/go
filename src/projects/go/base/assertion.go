package main

import "fmt"

func main() {
	// Type Assertion（中文名叫：类型断言），通过它可以做到以下几件事情
	// 检查 i 是否为 nil
	// 检查 i 存储的值是否为某个类型
    // 接口对象.(类型)

	// 第一种 t := i.(T) ch发 panic
	// 这个表达式可以断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T，如果断言成功，
	// 就会返回值给 t，如果断言失败，就会触发 panic。
	var i interface{} = 10
	t1 := i.(int)
	fmt.Println(t1)

	fmt.Println("=====分隔线=====")

	t2 := i.(string)
	fmt.Println(t2)
	// 如果要断言的接口值是 nil，那我们来看看也是不是也如预期一样会触发panic
	var j interface{} // nil
	var _ = j.(interface{})

	// 第二种 t, ok:= i.(T)
	// 这个表达式也是可以断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T，如果断言成功，
	// 就会返回其类型给 t，并且此时 ok 的值 为 true，表示断言成功。
	var l interface{} = 10
	t11, ok := l.(int)
	fmt.Printf("%d-%t\n", t11, ok)

	fmt.Println("=====分隔线1=====")

	t21, ok := l.(string)
	fmt.Printf("%s-%t\n", t21, ok)

	fmt.Println("=====分隔线2=====")

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok)

	fmt.Println("=====分隔线3=====")
	k = 10
	t4, ok := k.(interface{})
	fmt.Printf("%d-%t\n", t4, ok)

	t5, ok := k.(int)
	fmt.Printf("%d-%t\n", t5, ok)

	// Type Switch
	// 如果需要区分多种类型，可以使用 type switch 断言，这个将会比一个一个进行类型断言更简单、直接、高效。

	// 额外说明一下：
	//如果你的值是 nil，那么匹配的是 case nil
	//如果你的值在 switch-case 里并没有匹配对应的类型，那么走的是 default 分支
	//此外，还有两点需要你格外注意
	//
	//类型断言，仅能对静态类型为空接口（interface{}）的对象进行断言，否则会抛出错误，具体内容可以参考：关于接口的三个“潜规则”
	//类型断言完成后，实际上会返回静态类型为你断言的类型的对象，而要清楚原来的静态类型为空接口类型（interface{}），这是 Go 的隐式转换。

	findType(10)      // int
	findType("hello") // string

	var r interface{} // nil
	findType(r)

	findType(10.23) //float64
}
func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	default:
		fmt.Println(x, "not type matched")
	}
}

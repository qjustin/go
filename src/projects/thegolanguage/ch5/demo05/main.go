package main

import (
	"fmt"
	"sort"
)

/*
5.6. 匿名函数
0. 拥有函数名的函数只能在包级语法块中被声明，匿名函数可以绕过这个限制
1. 没有名字的函数，匿名函数可以在表达式中声明，或函数中
2. 匿名函数，可以访问完整的词法环境
3. 匿名函数的递归调用方式
*/

func main() {
	fmt.Println("Hello world")
	/*
		函数squares返回另一个类型为 func() int 的函数。对squares的
		一次调用会生成一个局部变量x并返回一个匿名函数。每次调用匿名函数时，
		该函数都会先使x的值加1，再返回x的平方。
		
		squares的例子证明，函数值不仅仅是一串代码，还记录了状态(变量x)。

		对squares的第一次调用会生成一个局部变量x并返回一个匿名函数。
		注意：这里变量x定义在func()的外部,x相对于func()来说是全局的
		因此，每次调用f(), x的值会被累加。

		这个例子说明了三点：
		1. squares的例子证明，函数值不仅仅是一串代码，还记录了状态。
		2. 函数值属于引用类型和函数值不可比较的原因。Go程序员也把函数值叫做闭包。
		3. 变量的生命周期不由它的作用域决定：squares返回后，变量x仍然隐式的存在于f中。
	*/

	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
 // 函数squares返回另一个类型为 func() int 的函数。
 func squares() func() int {
	 var x int
	 return func() int {
		 x++
		 return x * x
	 }
 }


 func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	/*
		当匿名函数需要被递归调用时，我们必须首先声明一个变量（在上面的例子中，我们首先声明了 visitAll），
		再将匿名函数赋值给这个变量。如果不分成两步，函数字面量无法与visitAll绑定，我们也无法递归调用该匿名函数。
		下面的代码用深度优先搜索了整张图
	*/
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}


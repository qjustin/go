package main

import "fmt"

func main() {
	// Go编译器，对于 { 和 } 的位置有严格的要求，它要求 else if （或 else）和 两边的花括号，必须在同一行。
	age := 20
	// 2. 单分支判断
	if age > 10 {
		fmt.Println(age)
	}

	// 在 if 里可以允许先运行一个表达式，取得变量后，再对其进行判断，比如第一个例子里代码也可以写成这样
	if age1 := 20; age1 > 18 {
		fmt.Println("已经成年了")
	}
}
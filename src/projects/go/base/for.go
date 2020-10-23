package main

import "fmt"

func main() {
	// 1. 接一个条件表达式
	a := 1
	for a < 100 {
		fmt.Println(a)
		a++
	}

	// 2. 接三个表达式
	for b := 2; b < 100; b++ {
		fmt.Println(b)
	}

	// 3. 不接表达式：无限循环
	//for ; ; {
	//	fmt.Println("hello")
	//}
	//// 等价于
	//for {
	//	fmt.Println("hello")
	//}

	// 4. 接 for-range 语句
	myarr := [...]string{"world", "python", "go"}
	for _, item := range myarr {
		fmt.Printf("hello, %s\n", item)
	}
}

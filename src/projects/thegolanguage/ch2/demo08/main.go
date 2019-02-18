package main

import (
	"fmt"
	"os"
)

/*
2.4. 赋值

 在包级别声明的变量会在main入口函数执行前完成初始化，局部变量将在声明语句被执行到的时候完成初始化。
*/
func main() {
	var str string
	fmt.Println(str)

	// int, int, int
	var i, j, k int
	// bool, float64, string
	var b, f, s = true, 2.3, "hello"

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(b)
	fmt.Println(f)
	fmt.Println(s)

	// 一组变量也可以通过调用一个函数，由函数返回的多个返回值初始化
	ref, err := os.Open("c:\text.txt")
	if err != nil {
		fmt.Println(err)
	}

	ref.Close()
}

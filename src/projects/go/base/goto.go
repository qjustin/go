package main

import "fmt"

func main() {
	i := 1
flag:
	if i <= 5 {
		fmt.Println(i)
		i++
		goto flag
	}

	// goto语句与标签之间不能有变量声明，否则编译错误。
//	fmt.Println("start")
//	goto flag1
//	var say = "hello oldboy"
//	fmt.Println(say)
//flag1:
//	fmt.Println("end")
}

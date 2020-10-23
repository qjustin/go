package main

import (
	"fmt"
	"time"
)

func main() {
	// 异常机制：panic 和 recover
	// 经过检查判断，当前环境无法达到我们程序进行的预期条件时（比如一个服务指定监听端口被其他程序占用），可以手动触发 panic，让程序退出停止运行。


	// 1. 触发panic
	// panic("test panic")

	// 2. 捕获 panic : recover
	// 发生了异常，有时候就得捕获，就像 Python 中的except 一样，那 Golang 中是如何做到的呢？
	// 这就不得不引出另外一个内建函数 – recover，它可以让程序在发生宕机后起生回生。
	// 但是 recover 的使用，有一个条件，就是它必须在 defer 函数中才能生效，其他作用域下，它是不工作的。
	set_data(20)
	// 如果能执行到这句，说明panic被捕获了
	// 后续的程序能继续运行
	fmt.Println("everything is ok")

	// 3. 无法跨协程
	// 但是这个 defer 在多个协程之间是没有效果，在子协程里触发 panic，只能触发自己协程内的 defer，
	// 而不能调用 main 协程里的 defer 函数的。

	// 这个 defer 并不会执行
	defer fmt.Println("in main")

	go func() {
		defer println("in goroutine")
		panic("")
	}()

	time.Sleep(2 * time.Second)
}

func set_data(x int) {
	// recover 它必须在 defer 函数中才能生效，其他作用域下，它是不工作的。
	defer func() {
		// recover() 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 故意制造数组越界，触发 panic
	var arr [10]int
	arr[x] = 88
}
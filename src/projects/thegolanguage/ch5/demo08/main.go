package main

import (
	"fmt"
	"os"
)

/*
	5.8. Deferred函数

	你只需要在调用普通函数或方法前加上关键字defer，就完成了defer所需要的语法。
	当执行到该条语句时，函数和参数表达式得到计算，但直到包含该defer语句的函数执
	行完毕时，defer后的函数才会被执行，不论包含defer语句的函数是通过return正
	常结束，还是由于panic导致的异常结束。你可以在一个函数中执行多条defer语句，
	它们的执行顺序与声明顺序相反。

	defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放
	锁。通过defer机制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。
	释放资源的defer应该直接跟在请求资源的语句后。
*/
func main() {
	fmt.Println("Hello world")
}
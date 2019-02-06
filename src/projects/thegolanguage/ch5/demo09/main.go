package main

import (
	"fmt"
	"runtime"
	"os"
)

/*
	5.9. Panic异常

	当panic异常发生时，程序会中断运行，并立即执行在该goroutine
	（可以先理解成线程，在第8章会详细介绍）中被延迟的函数（defer 机制）。
	随后，程序崩溃并输出日志信息。日志信息包括panic value和函数调用的堆栈跟踪信息。

	Go的panic机制类似于其他语言的异常，但panic的适用场景有一些不同。由于panic
	会引起程序的崩溃，因此panic一般用于严重错误，如程序内部的逻辑不一致。

	1. 直接调用内置的panic函数也会引发panic异常；panic函数接受任何值作为参数。
	当某些不应该发生的场景发生时，我们就应该调用panic。比如，当程序到达了某条
	逻辑上不可能到达的路径：

	2. 当输入导致程序无法执行下去时，可以抛出panic异常

	3. 用runtime 包打印堆栈信息,Go的panic机制中，延迟函数的调用在释放堆栈信息之前。
*/
func main() {
	defer printStack()
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

/*
	为了方便诊断问题，runtime包允许程序员输出堆栈信息。
	在下面的例子中，我们通过在main函数中延迟调用printStack输出堆栈信息。
*/
func printStack() {
	var buf [4096]byte
	/*
		将panic机制类比其他语言异常机制的读者可能会惊讶，
		runtime.Stack为何能输出已经被释放函数的信息？
		在Go的panic机制中，延迟函数的调用在释放堆栈信息之前。
	*/
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
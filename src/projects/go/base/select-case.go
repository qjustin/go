package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	6. 总结一下

	select 与 switch 原理很相似，但它的使用场景更特殊，学习了本篇文章，你需要知道如下几点区别：

	select 只能用于 channel 的操作(写入/读出)，而 switch 则更通用一些；
	select 的 case 是随机的，而 switch 里的 case 是顺序执行；
	select 要注意避免出现死锁，同时也可以自行实现超时机制；
	select 里没有类似 switch 里的 fallthrough 的用法；
	select 不能像 switch 一样接函数或其他表达式。

	 */

	// 1. 最简单的例子
	// 在运行 select 时，会遍历所有（如果有机会的话）的 case 表达式，只要有"一个"信道有接收到数据，那么 select 就结束，所以输出如下
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	c2 <- "hello"

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	default:
		fmt.Println("No data received.")
	}

	// 2. 避免造成死锁
	// select 在执行过程中，必须命中其中的某一分支。如果在遍历完所有的 case 后，若没有命中
	// 任何一个 case 表达式，就会进入 default 里的代码分支。
	c3 := make(chan string, 1)
	c4 := make(chan string, 2)

	// c4 <- "hello"

	select {
	case msg3 := <-c3:
		fmt.Println("c3 received: ", msg3)
	case msg4 := <-c4:
		fmt.Println("c4 received: ", msg4)
	default:
		fmt.Println("No data received.")
	}

	// 但如果你没有写 default 分支，select 就会阻塞，直到有某个 case 可以命中，
	// 而如果一直没有命中，select 就会抛出 deadlock 的错误，就像下面这样子。
	//select {
	//case msg3 := <-c3:
	//	fmt.Println("c3 received: ", msg3)
	//case msg4 := <-c4:
	//	fmt.Println("c4 received: ", msg4)
	//	//default:
	//	//	fmt.Println("No data received.")
	//}

	// 3. 解决死锁：解决这个问题的方法有两种
	// 		1. 在 select 的时候，也写好 default 分支代码，尽管你 default 下没有写任何代码。
	// 		2. 其中某一个信道可以接收到数据

	// 开启一个协程，可以发送数据到信道
	c5 := make(chan string, 1)
	c6 := make(chan string, 1)

	// 开启一个协程，可以发送数据到信道
	go func() {
		time.Sleep(time.Second * 1)
		c6 <- "hello"
	}()

	select {
	case msg5 := <-c5:
		fmt.Println("c5 received: ", msg5)
	case msg6 := <-c6:
		fmt.Println("c6 received: ", msg6)
	}

	// 4. select 随机性
	c7 := make(chan string, 1)
	c8 := make(chan string, 1)

	c7 <- "matrix"
	c8 <- "tony"

	select {
	case msg7 := <-c7:
		fmt.Println("c7 received: ", msg7)
	case msg8 := <-c8:
		fmt.Println("c8 received: ", msg8)
	default:
	}

	// 5. select 的超时设置
	c10 := make(chan string, 1)
	c11 := make(chan string, 1)
	timeout := make(chan bool, 1)

	go makeTimeout(timeout, 2)

	select {
	case msg10 := <-c10:
		fmt.Println("c10 received: ", msg10)
	case msg11 := <-c11:
		fmt.Println("c11 received: ", msg11)
	case <-timeout:
		fmt.Println("Timeout, exit.")
	}

	// 6. 读取/写入都可以
	// 上面例子里的 case，好像都只从信道中读取数据，但实际上，select
	//里的 case 表达式只要求你是对信道的操作即可，不管你是往信道写入数据，还是从信道读出数据。
	c12 := make(chan int, 1)
	c12 <- 2

	select {
	case c12 <- 4:
		fmt.Println("c12 received: ", <-c12)
		fmt.Println("c12 received: ", <-c12)
	default:
		fmt.Println("channel blocking")
	}
}

func makeTimeout(ch chan bool, t int) {
	time.Sleep(time.Second * time.Duration(t))
	ch <- true
}

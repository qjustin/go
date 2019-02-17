package main

import (
	"bytes"
	"fmt"
	"io"
)

/*
	7.5.1. 警告：一个包含nil指针的接口不是nil接口
	一个不包含任何值的nil接口值和一个刚好包含nil指针的接口值是不同的。
	这个细微区别产生了一个容易绊倒每个Go程序员的陷阱。
*/
const debug = true

func main() {
	fmt.Println("hello world")

	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}

	/*
		 当main函数调用函数f时，它给f函数的out参数赋了一个*bytes.Buffer的空指针，
		 所以out的动态值是nil。然而，它的动态类型是*bytes.Buffer，意思就是out变量
		 是一个包含空指针值的非空接口（如图）。
		 		    w
			   ---------------
		type  |	*bytes.Buffer |
		      |---------------|
		value |    *          |
			   ---------------
		所以所以防御性检查out!=nil的结果依然是true

		动态分配机制依然决定(*bytes.Buffer).Write的方法会被调用，但是这次的接收者
		的值是nil。对于一些如*os.File的类型，nil是一个有效的接收者（§6.2.1），但是
		*bytes.Buffer类型不在这些种类中。这个方法会被调用，但是当它尝试去获取缓冲区
		时会发生panic。
	*/
	f(buf) // NOTE: subtly incorrect!
	if debug {
		fmt.Println(buf)
	}

	/*
		问题在于尽管一个nil的*bytes.Buffer指针有实现这个接口的方法，它也不满足这
		个接口具体的行为上的要求。特别是这个调用违反了(*bytes.Buffer).Write方法
		的接收者非空的隐含先觉条件，所以将nil指针赋给这个接口是错误的。解决方案就是
		将main函数中的变量buf的类型改为io.Writer，因此可以避免一开始就将一个不完
		整的值赋值给这个接口

		var buf io.Writer
		if debug {
			buf = new(bytes.Buffer) // enable collection of output
		}
		f(buf) // OK
	*/
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

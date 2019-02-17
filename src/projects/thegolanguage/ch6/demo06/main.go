package main

import (
	"fmt"
)

/*
	6.6. 封装
	1. 一个对象的变量或者方法如果对调用方是不可见的话，一般就被定义为“封装”。

	Go语言只有一种控制可见性的手段：大写首字母的标识符会从定义它们的包中被
	导出，小写字母的则不会。这种限制包内成员的方式同样适用于struct或者一个
	类型的方法。
	
	封装的三个优点：

	第一，因为调用方不能直接修改对象的变量值，其只需要关注少量的语句并且只要
	弄懂少量变量的可能的值即可。

	第二，隐藏实现的细节，可以防止调用方依赖那些可能变化的具体实现，这样使设
	计包的程序员在不破坏对外的api情况下能得到更大的自由。

	第三，也是最重要的优点，是阻止了外部调用方对对象内部的值任意地进行修改。
	因为对象内部变量只可以被同一个包内的函数修改，所以包的作者可以让这些函数确保对象内部的一些值的不变性。

*/

type IntSet struct {
	// 用无符号数的slice来，表示bit数组
	words []uint64
}

type Buffer struct {
    buf     []byte
    initial [64]byte
    /* ... */
}

func main() {
	fmt.Println("hello world")
}

func (b *Buffer) Grow(n int) {
    if b.buf == nil {
        b.buf = b.initial[:0] // use preallocated space initially
    }
    if len(b.buf)+n > cap(b.buf) {
        buf := make([]byte, b.Len(), 2*cap(b.buf) + n)
        copy(buf, b.buf)
        b.buf = buf
    }
}
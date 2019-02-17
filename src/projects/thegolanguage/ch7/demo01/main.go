package main

import (
	"fmt"
)

/*
7.2. 接口类型
*/
func main() {
	fmt.Println("hello world")
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}
/*
	新的接口类型通过组合已有的接口来定义,语法和结构内嵌相似，我们可以用
	这种方式以一个简写命名一个接口，而不用声明它所有的方法。这种方式称为
	接口内嵌。
*/
type ReaderWriter interface {
	Reader
	Writer
}
/*
等价于
type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

等价于 混合风格
type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

3种定义方式都是一样的效果
*/
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
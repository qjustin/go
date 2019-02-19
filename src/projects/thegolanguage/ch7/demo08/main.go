package main

import (
	"errors"
	"fmt"
	"syscall"
)

/*
	7.8. error接口

	1. 创建一个error最简单的方法就是调用errors.New函数，它会根据传入的错误信息返回一个新的error。
	2. 调用errors.New函数是非常稀少的，因为有一个方便的封装函数fmt.Errorf
	3. Errno是一个系统调用错误的高效表示方式，它通过一个有限的集合进行描述，并且它满足标准的错误接口。
			 		    w
			   ---------------
		type  |	*bytes.Buffer |
		      |---------------|
		value |    *          |
			   ---------------
*/

func main() {
	fmt.Println(errors.New("EOF") == errors.New("EOF")) // "false"

	var err error = syscall.Errno(2)
	fmt.Println(err.Error()) // "no such file or directory"
	fmt.Println(err)         // "no such file or directory"
}

package main

import (
	"fmt"
	"os"
)

/*
	5.6.1. 警告：捕获迭代变量

	本节，将介绍Go词法作用域的一个陷阱。请务必仔细的阅读，弄清楚发生问题的原因。
	即使是经验丰富的程序员也会在这个问题上犯错误。

	
*/



func main() {
	fmt.Println("Hello world")
	
}
/*
	首先创建一些目录，再将目录删除
*/
func mkRmDirs() {
	/* 
		这段代码是错误的
		
		 循环变量dir在这个词法块中被声明。函数值func中记录的是循环变量的内存地址，
		 而不是循环变量某一时刻的值。以dir为例，后续的迭代会不断更新dir的值，
		 当删除操作执行时，for循环已完成，dir中存储的值等于最后一次迭代的值。
		 这意味着，每次对os.RemoveAll的调用删除的都是相同的目录。

		 我的理解：
		上一节中告诉我们，对于匿名函数func,变量dir不受作用域的影响，推出循环后变量
		dir仍然在匿名函数func中。换句话说func仍就引用着变量dir。
		func 引用了变量dir的地址而不是dir的值，循环推出后，dir是最后一次迭代的值。

		var rmdirs []func()
		for _, dir := range tempDirs() {
			os.MkdirAll(dir, 0755)
			rmdirs = append(rmdirs, func() {
				os.RemoveAll(dir) // NOTE: incorrect!
			})
		}

		// 这个问题不仅存在基于range的循环，在下面的例子中，对循环变量i的使用也存在同样的问题：
		var rmdirs []func()
		dirs := tempDirs()
		for i := 0; i < len(dirs); i++ {
			os.MkdirAll(dirs[i], 0755) // OK
			rmdirs = append(rmdirs, func() {
				os.RemoveAll(dirs[i]) // NOTE: incorrect!
			})
		}

		// go语句（第八章）或者defer语句（5.8节）会经常遇到此类问题。这不是go或defer本身导致的，而是因为它们都会等待循环结束后，再执行函数值。
	*/
	
	/* 
		这里是正确的代码
	*/

	var rmdirs []func()
	// for _, dir := range tempDirs()
	for _, dir := range [...]string {"abc", "def"} {
		/*
		 原因上面已经说明了，因此这一步是必须的
		 为了解决这个问题，我们会引入一个与循环变量同名的局部变量，作为循环变量的副本。
		 比如下面的变量dir，虽然这看起来很奇怪，但却很有用。
		*/
		dir := dir 
		os.MkdirAll(dir, 0755)

		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}
}
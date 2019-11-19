package main

import (
	"log"
	"os"

/**
	1. 同一个目录下的所有代码文件，必须使用相同的包名
	2. _ 包前的下划线，作用是对包做初始化，但并不使用包里的标识符。
		Go 不允许导入未使用的包，下划线让编译器接收这类导入，并调用
		包内所有代码文件里定义的init函数
	3. init 在 main 之前调用
 */
	_ "projects/goinaction/chapter2/sample/matchers"
	"projects/goinaction/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}

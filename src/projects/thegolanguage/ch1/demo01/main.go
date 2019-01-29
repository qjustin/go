package main

/**
* 缺少了必要的包或者导入了不需要的包，程序都无法编译通过
* Go语言在代码格式上采取了很强硬的态度。gofmt工具把代码格式化为标准格式
* Go语言不允许使用无用的局部变量（local variables），因为这会导致编译错误
 */
import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}

	fmt.Println(s)
	fmt.Println("for1")
}

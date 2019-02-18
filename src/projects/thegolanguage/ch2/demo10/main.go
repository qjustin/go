package main

/*
2.6. 包和文件
2.6.1. 导入包

1. 一个包的源代码保存在一个或多个以.go为文件后缀名的源文件中
2. 一个包的名字和包的导入路径的最后一个字段相同(也就是说：包名称应该与所在得目录名称一致)
3. 如果导入了一个包，但是又没有使用该包将被当作一个编译错误处理。
4. 导入声明10.4 可以给包取别名
*/
import (
	"fmt"
	"os"
	"projects/thegolanguage/ch2/tempconv"
	"strconv"
	//D:\Dev\sourcecode\GITHUB\go\src\projects\thegolanguage\ch2
	// 按照惯例，一个包的名字和包的导入路径的最后一个字段相同
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		// 通过该短小的名字就可以引用包中导出的全部内容。
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

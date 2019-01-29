package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建一个Map key是string类型，value是int类型， 类似java的HashMap
	counts := make(map[string]int)
	// bufio包，它使处理输入和输出方便又高效。Scanner类型是该包最有用的特性之一
	input := bufio.NewScanner(os.Stdin)
	// input.Scan()，即读入下一行，并移除行末的换行符， 读到一行时返回true， 不再有输入时返回false
	for input.Scan() {
		// 读取一行输入，作为Key，value 自增1 等价于：
		// line := input.Text()
		// counts[line] = counts[line] + 1
		// map中不含某个键时不用担心，首次读到新行时，等号右边的表达式counts[line]的值将被计算为其类型的零值，对于int即0。
		// map顺序是随机的。无须的。
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			// 按照惯例，以字母f结尾的格式化函数
			//	%d          十进制整数
			//	%x, %o, %b  十六进制，八进制，二进制整数。
			//	%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
			//	%t          布尔：true或false
			//	%c          字符（rune） (Unicode码点)
			//	%s          字符串
			//	%q          带双引号的字符串"abc"或带单引号的字符'c'
			//	%v          变量的自然形式（natural format）
			//	%T          变量的类型
			//	%%          字面上的百分号标志（无操作数）
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

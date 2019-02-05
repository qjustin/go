package main

import (
	"fmt"
	"os"
)

/*
	5.7. 可变参数
	1. 声明可变参数：在参数列表的最后一个参数类型之前加上省略符号“...”
		func sum(vals...int) int
	2. 可变参数，可以接受一个数组或切片
		数组：sum(1, 2, 3, 4)
		切片：sum(values...)
	3. 可变参数函数和以切片作为参数的函数是不同的
		签名：
			func(...int)
			func([]int)
	4. 可变参数函数经常被用于格式化字符串。
*/
func main() {
	fmt.Println("Hello world")

	/*
		调用者隐式的创建一个数组，并将原始参数复制到数组中，
		再把数组的一个切片作为参数传给被调用函数。
	*/
	fmt.Println(sum())           // "0"
	fmt.Println(sum(3))          // "3"
	fmt.Println(sum(1, 2, 3, 4)) // "10"

	/*
		如果原始参数已经是切片类型，我们该如何传递给sum？
		只需在最后一个参数后加上省略符。
	*/
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"

	/*
		...int 型参数的行为看起来很像切片类型，但实际上，
		可变参数函数和以切片作为参数的函数是不同的。
	*/
	fmt.Printf("%T\n", f) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"

	/*
		4. 可变参数函数经常被用于格式化字符串。
	*/
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name) // "Line 12: undefined: count"
}

/*
	在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上
	省略符号“...”，这表示该函数会接收任意数量的该类型参数。
*/
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}

	return total
}

func f(...int) {
	fmt.Println("f(...int)")
}
func g([]int) {
	fmt.Println("g([]int)")
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 因为uint8 和 uint32 ，直观上让人以为这是一个数值，但是实际上，它也可以表示一个字符，
	// 所以为了消除这种直观错觉，就诞生了 byte 和 rune 这两个别名类型。

	// byte，占用1个节字，就 8 个比特位，所以它和 uint8 类型本质上没有区别
	var a byte = 'A'
	var b uint8 = 'B'
	fmt.Printf("a 的值: %c \nb 的值: %c\n", a, b)

	var c byte = 65
	var d uint8 = 66
	fmt.Printf("c 的值: %c \nd 的值: %c\n", c, d)
	// rune，占用4个字节，共32位比特位，所以它和 uint32 本质上也没有区别。它表示的是一个 Unicode字符
	var e byte = 'A'
	var f rune = 'B'
	fmt.Printf("e 占用 %d 个字节数\nf 占用 %d 个字节数", unsafe.Sizeof(e), unsafe.Sizeof(f))

	var mystr01 string = "\\r\\n"
	var mystr02 string = `\r\n`
	fmt.Println(mystr01)
	fmt.Println(mystr02)

	// 你可以使用 fmt 的 %q 来还原一下。
	var mystr03 string = `\r\n`
	fmt.Println(`\r\n`)
	fmt.Printf("的解释型字符串是： %q", mystr03)
}
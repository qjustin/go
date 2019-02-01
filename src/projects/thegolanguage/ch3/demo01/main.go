package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
3.5.4. 字符串和Byte切片
标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。
strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。
bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。因为字符串是只读的，因此逐步构建字符串会导致很多分配和复制。在这种情况下，使用bytes.Buffer类型将会更有效
strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。
unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。每个函数有一个单一的rune类型的参数，然后返回一个布尔值。而像ToUpper和ToLower之类的转换函数将用于rune字符的大小写转换。

3.5.5. 字符串和数字的转换
字符串和数值之间的转换也比较常见。由strconv包提供这类转换功能。
*/

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("a.bc.go"))
	fmt.Println(basename("abc"))
	bToS()
	sToI()
}

func basename(s string) string {
	return basenamev1(s)
}
func basenamev1(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basenamev2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// 一个字符串是包含只读字节的数组，一旦创建，是不可变的。相比之下，一个字节slice的元素则可以自由地修改。
func bToS() {
	// 字符串和字节slice之间可以相互转换
	s := "abc"
	// 一个[]byte(s)转换是分配了一个新的字节数组用于保存字符串数据的拷贝，然后引用这个底层的字节数组。
	b := []byte(s)
	// 一个字节slice转换到字符串的string(b)操作则是构造一个字符串拷贝，以确保s2字符串是只读的
	s2 := string(b)

	fmt.Print(s2)
}

func sToI() {
	// 将一个整数转为字符串，一种方法是用fmt.Sprintf返回一个格式化的字符串；另一个方法是用strconv.Itoa(“整数到ASCII”)：
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	// FormatInt和FormatUint函数可以用不同的进制来格式化数字
	fmt.Println(strconv.FormatInt(int64(x), 2))

	// fmt.Printf函数的%b、%d、%o和%x等参数提供功能往往比strconv包的Format函数方便很多，特别是在需要包含有附加额外信息的时候：
	s := fmt.Sprintf("x=%b", x)

	// 将一个字符串解析为整数，可以使用strconv包的Atoi或ParseInt函数
	a, err := strconv.Atoi("123") // x is an int
	// ParseInt函数的第三个参数是用于指定整型数的大小；例如16表示int16，0则表示int。
	// 在任何情况下，返回的结果y总是int64类型，你可以通过强制类型转换将它转为更小的整数类型。
	b, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits

	fmt.Printf("%s, %d, %d, %s", s, a, b, err)
}

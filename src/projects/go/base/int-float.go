package main

import "fmt"

func main() {
	// int 并没有指定它的位数，说明它的大小，是可以变化的
	var c1 int = 0b1100 // 二进制 0b 1.13 版本才支持
	var c2 int = 0o14	// 八进制 0o 1.13 版本才支持
	var c3 int = 0xC	// 16进制

	fmt.Printf("2进制数 %b 表示的是: %d \n", c1, c1)
	fmt.Printf("8进制数 %o 表示的是: %d \n", c2, c2)
	fmt.Printf("16进制数 %X 表示的是: %d \n", c3, c3)

	// 3.7E-2表示浮点数0.037。又比如，3.7E+1表示浮点数37
	// 37.0可以被简化为37。又比如，0.037可以被简化为.037

	// float32 和 float64
	// float32，也即我们常说的单精度，存储占用4个字节，也即4*8=32位，其中1位用来符号，8位用来指数，剩下的23位表示尾数
	// float64，也即我们熟悉的双精度，存储占用8个字节，也即8*8=64位，其中1位用来符号，11位用来指数，剩下的52位表示尾数
	// float32 小数部分只能精确到后面6位，加上小数点前的一位，即有效数字为7位。
	// float64 小数部分只能精确到小数点后15位，加上小数点前的一位，有效位数为16位

	// 出现这种很怪异的现象，myfloat == myfloat +1 会返回 true

	// 10000018 科学计数法就是 1.0000018 * 10^7， 能精确到小数点后面6位
	// 刚好满足精度要求数据的临界情况
	var myfloat float32 = 10000018
	fmt.Println("myfloat: ", myfloat)
	fmt.Println("myfloat: ", myfloat+1)
	// 刚好不满足精度要求的例子
	var myfloat01 float32 = 100000182
	var myfloat02 float32 = 100000187
	fmt.Println("myfloat: ", myfloat01)
	fmt.Println("myfloat: ", myfloat01+5)
	fmt.Println(myfloat02 == myfloat01+5)
}
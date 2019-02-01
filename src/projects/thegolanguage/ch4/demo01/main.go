package main

import (
	"fmt"
	"crypto/sha256"
)
/*
	4.1. 数组
	1.数组的长度是固定的
	2.数组元素的类型是一致的
	3.数组定义时必须确定数组长度。一旦确定不允许改变(数组的长度需要在编译阶段确定)
	4.数组下标从0开始

	当传递的参数是一个数组时，应当传递数组地址，避免赋值
*/
func main() {
	// 数组初始化方式
	var a [3]int = [3]int{1, 2, 3}
	var b [3]int = [3]int{1, 2}
	fmt.Println(b[2]) // "0"
	// 如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算。
	c := [...]int{1, 2, 3}
	 d := [3]int{1, 2, 3}
	// .数组定义时必须确定数组长度。一旦确定不允许改变
	//d = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int

	// 指定一个索引和对应值列表的方式初始化数组
	e := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, e[RMB])
	// 定义了一个含有100个元素的数组r，最后一个元素被初始化为-1，其它元素都是用0初始化。
	f := [...]int{99: -1}

	// 数据元素类型可以使用==来比较，那么数组也可以直接通过==来比较，素有元素相等时，返回true
	g := [2]int{1, 2}
	h := [...]int{1, 2}
	i := [2]int{1, 3}
	fmt.Println(g == h, g == i, h == i) // "true false false"
	j := [3]int{1, 2}
	fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int

	// crypto/sha256包的Sum256函数对一个任意的字节slice类型的数据生成一个对应的消息摘要。
	k := sha256.Sum256([]byte("x"))
	l := sha256.Sum256([]byte("X"))
	// %x 指定以十六进制的格式打印数组或slice全部的元素 %t 表示布尔值
	fmt.Printf("%x\n%x\n%t\n%T\n", k, l, k == l, k)

	fmt.Printf("%d, %d, %d, %d, %s, %d, %d", a[0], b[0], c[0], d[0], e[0], f[0], j[0])


}
	/* 
	调用一个函数时，函数的每个调用参数将会被赋值给函数内部的参数变量
	所以函数参数变量接收的是一个复制的副本，并不是原始调用的变量，这样的
	参数传递的机制导致传递大的数组类型将是低效的，参数传递的机制导致传递
	大的数组类型将是低效的。

	我们可以显式地传入一个数组指针，那样的话函数通过指针对数组的任何修改都可以直接反馈到调用者。
	*/

func zero(ptr *[32]byte) {
	// 给数组清零
	for i := range ptr {
		ptr[i] = 0
	}

	// 等价于，数组字面值[32]byte{}就可以生成一个32字节的数组。
	// 而且每个数组的元素都是零值初始化，也就是0
	
	// 给数组清零
	*ptr = [32]byte{}
}


type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)
package main

import (
	"fmt"
)

/*
	7.13. 类型分支
	接口被以两种不同的方式使用:

	1. 一个接口的方法表达了实现这个接口的具体类型间的相似性，但是隐藏了代码的细节和这些具体类型本身的操作。
	重点在于方法上，而不是具体的类型上。

	2. 利用一个接口值可以持有各种具体类型值的能力，将这个接口认为是这些类型的联合。类型断言用来动态地区别
	   这些类型，使得对每一种情况都不一样。在这个方式中，重点在于具体的类型满足这个接口，而不在于接口的方
	   法（如果它确实有一些的话），并且没有任何的信息隐藏。我们将以这种方式使用的接口描述为
	   discriminated unions（可辨识联合）。
*/

func main() {
	fmt.Println("hello world")
}

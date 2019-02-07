

package main

import (
	"fmt"
)

type Point struct{ X, Y float64 }

/*
6.6. 封装

一个对象的变量或者方法如果对调用方是不可见的话，一般就被定义为“封装”。


Go语言只有一种控制可见性的手段：大写首字母的标识符会从定义它们的包中被导出，小写字母的则不会。
这种限制包内成员的方式同样适用于struct或者一个类型的方法。
*/
func main() {
	fmt.Println("hello world")
}
package main

import "fmt"

func main() {
	// 变量声明后必须使用否则报错
	// 变量/常量都只能声明一次，声明多次，编译就会报错。

	// 第一种 ：一行声明一个变量
	var name string = "hello"
	fmt.Println(name)

	// 第二种：多个变量一起声明
	var (
		str    string
		age    int
		gender string
	)
	str = "str1"
	age = 0
	gender = "F"
	fmt.Println(str)
	fmt.Println(age)
	fmt.Println(gender)

	// 第三种：声明和初始化一个变量
	// 使用 := （推导声明写法或者短类型声明法：编译器会自动根据右值类型推断出左值的对应类型。）
	// := 限制 只能用于函数内部
	c1 := "hello"
	fmt.Println(c1)

	// 第四种：声明和初始化多个变量 (这种方法，也经常用于变量的交换)
	c2, c3 := 2, 3
	fmt.Println(c2)
	fmt.Println(c3)
	var c4 int = 1
	var c5 int = 2
	c4, c5 = c5, c4
	fmt.Println(c4)
	fmt.Println(c5)

	// 第五种：new 函数声明一个指针变量
	var c6 int = 6
	var c7 = &c6 // &后面接变量名，表示取出该变量的内存地址
	fmt.Println(c6)
	fmt.Println(c7)

	//return new(int)
	//等价于
	//var dummy int
	//return &dummy
	c8 := new(int) //new(Type) 创建一个Type类型的匿名变量，初始化为Type类型的零值，然后返回变量地址，返回的指针类型为*Type
	fmt.Println(c8)

	// 第六种：“_" 下划线称为,匿名变量，也称作占位符 ：通常我们用匿名接收必须接收，但是又不会用到的值
	// 匿名变量，优点有三：
	// 不分配内存，不占用内存空间
	// 不需要你为命名无用的变量名而纠结
	// 多次声明不会有任何问题
	//func GetData() (int, int) {
	//	return 100, 200
	//}
	//func main(){
	//	a, _ := GetData()
	//	_, b := GetData()
	//	fmt.Println(a, b)
	//}
}

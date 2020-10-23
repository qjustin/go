package main

import "fmt"

func main() {

	// 根据变量指向的值，是否是内存地址，我把变量分为两种：
	// & ：从一个普通变量中取得内存地址
	// * ：当*在赋值操作值的右边，是从一个指针变量中取得变量值，当*在赋值操作值的左边，是指该指针指向的变量
	// 普通变量：存数据值本身
	// 指针变量：存值的内存地址

	// 指针创建有三种方法
	// 方法1
	a := 1
	ptr := &a
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", ptr)

	// 方法2
	b := new(string)
	*b = "hello"
	fmt.Printf("%s\n", *b)

	// 方法3
	c := 2
	var ptrc *int
	ptrc = &c
	fmt.Printf("%d\n", *ptrc)

	// 2. 指针的类型
	astr := "hello"
	aint := 1
	abool := false
	arune := 'a'
	afloat := 1.2

	fmt.Printf("astr 指针类型是：%T\n", &astr)
	fmt.Printf("aint 指针类型是：%T\n", &aint)
	fmt.Printf("abool 指针类型是：%T\n", &abool)
	fmt.Printf("arune 指针类型是：%T\n", &arune)
	fmt.Printf("afloat 指针类型是：%T\n", &afloat)

	// 这个 只能接收一个*int类型的指针
	// mytest(&afloat) // 因此报错

	// 3. 指针的零值 为 nil

	// 4. 指针与切片
	//切片与指针一样，都是引用类型。
	//如果我们想通过一个函数改变一个数组的值，有两种方法
	//将这个数组的切片做为参数传给函数
	//将这个数组的指针做为参数传给函数
	arr9 := [3]int{89, 90, 91}
	modifysli(arr9[:])
	fmt.Println(arr9)

	modifyarr(&arr9)
	fmt.Println(arr9)
}

func modifyarr(arr *[3]int) {
	(*arr)[0] = 90
}

func modifysli(sls []int) {
	sls[0] = 90
}

// 所以若我们定义一个只接收指针类型的参数的函数，可以这么写
func mytest(ptr *int) {
	fmt.Println(*ptr)
}

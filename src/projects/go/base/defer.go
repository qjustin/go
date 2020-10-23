package main

import "fmt"

func main() {
	// 1. 延迟调用
	// defer 的用法很简单，只要在后面跟一个函数的调用，就能实现将这个 xxx 函数的调用延迟到当前函数执行完后再执行。
	defer fmt.Println("B")
	fmt.Println("A")

	// 2. 即时求值的变量快照
	name := "go"
	defer fmt.Println(name) // 输出: go

	name = "python"
	fmt.Println(name)      // 输出: python

	// 3. 多个defer 反序调用
	// 越早定义的defer越晚执行

	// 4. defer 与 return 孰先孰后（return 先执行）
	// 那就是 defer 是return 后才调用的。所以在执行 defer 前，myname 已经被赋值成 go 了。
	myname := myfunc()
	fmt.Printf("main 函数里的name: %s\n", name)
	fmt.Println("main 函数里的myname: ", myname)

	// 4. 为什么要有 defer？
	// 固然可以，但是当一个函数里有多个 return 时，你得多调用好多次这个函数，代码就臃肿起来了。
	//若是没有 defer，你可以写出这样的代码
	//func f() {
	//	r := getResource()  //0，获取资源
	//	......
	//	if ... {
	//		r.release()  //1，释放资源
	//		return
	//	}
	//	......
	//	if ... {
	//		r.release()  //2，释放资源
	//		return
	//	}
	//	......
	//	if ... {
	//		r.release()  //3，释放资源
	//		return
	//	}
	//	......
	//	r.release()     //4，释放资源
	//	return
	//}
	// 使用了 defer 后，代码就显得简单直接，不管你在何处 return，都会执行 defer 后的函数。
	// func f() {
	//    r := getResource()  //0，获取资源
	//
	//    defer r.release()  //1，释放资源
	//    ......
	//    if ... {
	//        ...
	//        return
	//    }
	//    ......
	//    if ... {
	//        ...
	//        return
	//    }
	//    ......
	//    if ... {
	//        ...
	//        return
	//    }
	//    ......
	//    return
	//}
}

var name string = "go"

func myfunc() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myfunc 函数里的name：%s\n", name)
	return name
}
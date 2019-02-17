package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
)

/*
7.5. 接口值

概念上讲一个接口的值，接口值，由两个部分组成，一个具体的类型和那个类型的值。
它们被称为接口的动态类型和动态值。Go语言这种静态类型的语言，类型是编译期的
概念；因此一个类型不是一个值.在我们的概念模型中，一些提供每个类型信息的值被
称为类型描述符，比如类型的名称和方法。在一个接口值中，类型部分代表与之相关
类型的描述符。

一个接口值可以持有任意大的动态值。例如，表示时间实例的time.Time类型，这个
类型有几个对外不公开的字段。我们从它上面创建一个接口值：
var x interface{} = time.Now()

*/
func main() {
	fmt.Println("hello world")

	// 变量w得到了3个不同的值。（开始和最后的值是相同的）
	
	/*
	第一句
	定义了变量w,变量总是被一个定义明确的值初始化，即使接口类型也不例外。
	对于一个接口的零值就是它的类型和值的部分都是nil
		     w
		   ------
	type  |	nill |
	      |------|
	value | nill |
		   ------
	
	空指针异常
	w.Write([]byte("hello")) // panic: nil pointer dereference
	*/
	var w io.Writer
	fmt.Printf("%T\n", w) // "<nil>"

	/*
	第二句
	这个赋值过程调用了一个具体类型到接口类型的隐式转换，这和显式的使用
	io.Writer(os.Stdout)是等价的。这个接口值的动态类型被设为*os.File
	指针的类型描述符，它的动态值持有os.Stdout的拷贝；这是一个代表处理
	标准输出的os.File类型变量的指针
		       w
		   ----------
	type  |	*os.File |			  os.File
	      |----------| 		-------------------
	value |    *     | --> | fd int = 1(stdout)|
		   ----------	 	-------------------

	调用一个包含*os.File类型指针的接口值的Write方法，使得(*os.File).Write
	方法被调用。这个调用输出“hello”。
	w.Write([]byte("hello")) // "hello"

	通常在编译期，我们不知道接口值的动态类型是什么，所以一个接口上的调用必须使用
	动态分配。因为不是直接进行调用，所以编译器必须把代码生成在类型描述符的方法
	Write上，然后间接调用那个地址。这个调用的接收者是一个接口动态值的拷贝，
	os.Stdout。效果和下面这个直接调用一样：

	os.Stdout.Write([]byte("hello"))
	*/
	w = os.Stdout
	fmt.Printf("%T\n", w) // "*os.File"

	/*
	第三句
	现在动态类型是*bytes.Buffer并且动态值是一个指向新分配的缓冲区的指针
		       w
		   ---------------
	type  |	*bytes.Buffer |		  bytes.Buffer
	      |---------------| 	 --------------
	value |    *          | --> | data []bytes |
		   ---------------	 	 --------------

	这次类型描述符是*bytes.Buffer，所以调用了(*bytes.Buffer).Write方法，
	并且接收者是该缓冲区的地址。这个调用把字符串“hello”添加到缓冲区中。

	w.Write([]byte("hello")) // writes "hello" to the bytes.Buffers
	*/
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) // "*bytes.Buffer"

	/*
	第四句
	语句将nil赋给了接口值，把变量w恢复到和它之前定义时相同的状态
		     w
		   ------
	type  |	nill |
	      |------|
	value | nill |
		   ------
	*/
	w = nil
	fmt.Printf("%T\n", w) // "<nil>"

	/*
		 接口类型是非常与众不同的。其它类型要么是安全的可比较类型（如基本类型和指针）
		 要么是完全不可比较的类型（如切片，映射类型，和函数）

		接口值可以使用==和!＝来进行比较。两个接口值相等仅当它们都是nil值，或者它们的
		动态类型相同并且动态值也根据这个动态类型的==操作相等。因为接口值是可比较的，
		所以它们可以用在map的键或者作为switch语句的操作数。

		然而，如果两个接口值的动态类型相同，但是这个动态类型是不可比较的
	*/
	var x interface{} = []int{1,2,3}
	//fmt.Println(x == x) // painc 
	fmt.Println(x)
	fmt.Printf("%T\n", w) // 打印接口值
	fmt.Println(w)
}


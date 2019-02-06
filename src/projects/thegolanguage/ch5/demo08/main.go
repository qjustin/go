package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*
	5.8. Deferred函数

	你只需要在调用普通函数或方法前加上关键字defer，就完成了defer所需要的语法。
	当执行到该条语句时，函数和参数表达式得到计算，但直到包含该defer语句的函数执
	行完毕时，defer后的函数才会被执行，不论包含defer语句的函数是通过return正
	常结束，还是由于panic导致的异常结束。你可以在一个函数中执行多条defer语句，
	它们的执行顺序与声明顺序相反。

	defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放
	锁。通过defer机制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。
	释放资源的defer应该直接跟在请求资源的语句后。

	defer语句中的函数会在return语句更新返回值变量后再执行，又因为在函数中定义的
	匿名函数可以访问该函数包括返回值变量在内的所有变量，所以，对匿名函数采用defer
	机制，可以使其观察函数的返回值。

	理解：
	defer 修饰的语句，在函数执行之前和之后一定会被调用，类似java中的finally机制
	1. 调试复杂程序时，defer机制也常被用于记录何时进入和退出函数。
	2. 利用defer机制观察匿名函数的返回值
	3. defer配合匿名函数可以修改函数返回值
	4. 在循环体中的defer语句需要特别注意，因为只有在函数执行完毕后，这些被延迟的函数才会执行。
*/
func main() {
	bigSlowOperation()

	_ = double(4)

	fmt.Println(triple(4)) // "12"

}

func bigSlowOperation() {
	// 需要注意一点：不要忘记defer语句后的圆括号，否则本该在进入时执行的操作会在
	// 退出时执行，而本该在退出时执行的，永远不会被执行。
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

/*
	defer语句中的函数会在return语句更新返回值变量后再执行，又因为在函数中定义
	的匿名函数可以访问该函数包括返回值变量在内的所有变量，所以，对匿名函数采用
	defer机制，可以使其观察函数的返回值。
*/
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

/*
	被延迟执行的匿名函数甚至可以修改函数返回给调用者的返回值
*/
func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func openFile() (str string, err error) {

	/*
		在循环体中的defer语句需要特别注意，因为只有在函数执行完毕后，这些被延迟
		的函数才会执行。下面的代码会导致系统的文件描述符耗尽，因为在所有文件都被
		处理之前，没有文件会被关闭。
	*/
	for _, filename := range [...]string{"a.txt", "b.txt"} {
		f, err := os.Open(filename)
		if err != nil {
			return "a", err
		}
		defer f.Close() // NOTE: risky; could run out of file descriptors
		// ...process f…
	}

	return "a", nil
}

func openFileV2() (str string, err error) {
	/*
		一种解决方法是将循环体中的defer语句移至另外一个函数。在每次循环时，调用这个函数。
	*/
	for _, filename := range [...]string{"", "", ""} {
		if err := doFile(filename); err != nil {
			return "error", err
		}
	}

	return "a", nil
}

// 一种解决方法是将循环体中的defer语句移至另外一个函数。在每次循环时，调用这个函数。
func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}
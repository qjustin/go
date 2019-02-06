package main

import (
	"fmt"
)

/*
	5.10. Recover捕获异常
	
	1. 通常来说，不应该对panic异常做任何处理，但有时需要从异常中恢复
	2. 如何捕捉函数的异常？
		1. 在函数A内部定义一个 defer匿名函数，
		2. 在defer匿名函数中调用 recover() 函数
		3. recover() 函数会从painc中恢复，并返回panic value
		4. 导致panic异常的函数不会继续运行，但能正常返回。 
		5. 在未发生panic时调用recover，recover会返回nil。
	3. 公有的API应该将函数的运行失败作为error返回
	4. 你也不应该恢复一个由他人开发的函数引起的panic，比如说调用者
	   传入的回调函数，因为你无法确保这样做是安全的。
	5. 为了标识某个panic是否应该被恢复，我们可以将panic value设置成特殊类型。
*/
func main() {
	fmt.Println("hello world")
}

// func Parse(input string) (s *Syntax, err error) {
// 	defer func() {
// 		if p := recover(); p != nil {
// 			err = fmt.Errorf("internal error: %v", p)
// 		}
// 	}()
// 	// ...parser...
// }
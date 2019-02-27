package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// 用来等待程序完成
	var wg sync.WaitGroup
	// 2表示等待两个goroutine完成
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	// 等待goroutine结束
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

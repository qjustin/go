package main

import (
	"fmt"
	"sync"
	"time"
)

// https://www.cnblogs.com/netuml/p/9063301.html
func main() {

}

func SyncWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 1; i < 5; i++ {
			time.Sleep(2 * time.Second)
			fmt.Println("func1")

		}

		wg.Done()
	}()

	go func() {
		for i := 1; i < 5; i++ {
			time.Sleep(2 * time.Second)
			fmt.Println("func2")
		}

		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All Done")
}

func SyncChanSelect() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

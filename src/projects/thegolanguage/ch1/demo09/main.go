package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// main函数本身也运行在一个goroutine中， gouroutine 之间通过channel来传递参数
// goroutine 尝试在一个channel上做send(ch <- expression)或者receive(<- ch)操作时
func main() {
	start := time.Now()
	// 创建了一个传递string类型参数的channel
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// 表示创建一个新的goroutine，并在这个新的goroutine中执行fetch函数。
		go fetch(url, ch)
	}

	// 这个程序中我们用main函数来接收所有fetch函数传回的字符串，可以避免在goroutine异步执行还没有完成时main函数提前退出。
	for range os.Args[1:] {
		// 主函数(主线程)负责接收Channel中传递过来的值(<-ch)
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		//往Channel中发送一个err
		ch <- fmt.Sprint(err)
		return
	}
	// io.Copy会把响应的Body内容拷贝到ioutil.Discard输出流中
	// 可以把这个ioutil.Discard变量看作一个垃圾桶，可以向里面写一些不需要的数据
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		// 往通道中发送(写入) 格式化后的值
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

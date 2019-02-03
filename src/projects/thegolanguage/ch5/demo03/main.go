package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

/*
5.4. 错误
1. panic是来自被调用函数的信号，表示发生了某个已知的bug。
2. 如果导致失败的原因只有一个，额外的返回值可以是一个布尔值，通常被命名为ok。
3. 通常，导致失败的原因不止一种，额外的返回值不再是简单的布尔类型，而是error类型,error类型可能是nil或者non-nil,nil意味着函数运行成功，non-nil表示失败。
4. 通常，当函数返回non-nil的error时，其他的返回值是未定义的（undefined），这些未定义的返回值应该被忽略。

5.4.1. 错误处理策略
当一次函数调用返回错误时，调用者应该选择合适的方式处理错误。有很多处理方式，让我们来看看常用的五种方式。
1. 直接将错误返回给调用者或构建一个新的包含更多错误信息的error，然后返回给调用者。
	注意错误信息中应该避免大写于换行符，错误信息应该尽可能的详尽。
2. 如果错误的发生是偶然性的，或由不可预知的问题导致的。一个明智的选择是重新尝试失败的操作。
	在重试时，我们需要限制重试的时间间隔或重试的次数，防止无限制的重试。
3. 如果错误发生后，程序无法继续运行，我们就可以采用第三种策略：输出错误信息并结束程序。
	这种策略只应在main中执行。对库函数而言，应仅向上传播错误，除非该错误意味着程序内部包含不一致性，即遇到了bug，才能在库函数中结束程序。
4. 第四种策略：有时，我们只需要输出错误信息就足够了，不需要中断程序的运行。我们可以通过log包提供函数
5. 第五种，也是最后一种策略：我们可以直接忽略掉错误。


5.4.2. 文件结尾错误（EOF）
io包保证任何由文件结束引起的读取失败都返回同一个错误——io.EOF, 因为文件结束这种错误不需要更多的描述，所以io.EOF有固定的错误信息——“EOF”。
*/

func main() {

	// value, ok := cache.Lookup(key)
	// if !ok {
	// 	fmt.Println(value)
	// }

	// 3. 在主函数中，输出错误信息并结束程序。
	url := "www.163.com"
	if err := WaitForServer(url); err != nil {
		// 写log
		log.Fatalf("Site is down: %v\n", err)
		// 标准错误流输出错误信息。
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}

	// 4. 不需要中断程序的运行。接写log了事
	url1 := "www.163.com"
	if err := WaitForServer(url1); err != nil {
		// 写log
		log.Fatalf("Site is down: %v\n", err)
		// 标准错误流输出错误信息。
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)

	}

}

// 2. 重试策略
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s);retrying…", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func IoRead() error {
	// 调用者只需通过简单的比较，就可以检测出这个错误。下面的例子展示了如何从标准输入中读取字符，以及判断文件结束。
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break // finished reading
		}
		if err != nil {
			return fmt.Errorf("read failed:%v", err)
		}

		fmt.Println(r)
	}

	return nil
}

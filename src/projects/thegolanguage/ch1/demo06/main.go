package main

import (
	"bufio"
	"fmt"
	"os"
)

// go run main.go C:\Users\depl_\Desktop\qianyi.txt
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// map作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本），
		// 被调用函数对map底层数据结构的任何修改，调用者函数都可以通过持有的map引用看到。
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// os.Open函数返回两个值, 一个是：os.Open函数返回两个值。 另一个是：内置error类型的值，err != nil 表示文件被打开。
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

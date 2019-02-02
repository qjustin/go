package main

import "fmt"

/*
4.2.1. append函数
append函数对于理解slice底层是如何工作的非常重要，所以让我们仔细查看究竟是发生了什么。

*/

func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
}

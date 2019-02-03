package main

import (
	"fmt"
)

/*
5.2. 递归
函数可以是递归的，这意味着函数可以直接或间接的调用自身。
略....

5.3. 多返回值
1. 调用多返回值函数时，返回给调用者的是一组值，调用者必须显式的将这些值分配给变量:
2. 如果一个函数所有的返回值都有显式的变量名，那么该函数的return语句可以省略操作数。
3. 一个函数内部可以将另一个有多返回值的函数调用作为返回值
4. 当你调用接受多参数的函数时，可以将一个返回多参数的函数调用作为该函数的参数。
*/

func main() {
	fmt.Println("Hello world")
}

/*
1. 如果一个函数所有的返回值都有显式的变量名，那么该函数的return语句可以省略操作数。
按照返回值列表的次序，返回所有的返回值，在上面的例子中，每一个return语句等价于：return words, images, err
*/

// func CountWordsAndImages(url string) (words, images int, err error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return
// 	}
// 	doc, err := html.Parse(resp.Body)
// 	resp.Body.Close()
// 	if err != nil {
// 		err = fmt.Errorf("parsing HTML: %s", err)
// 		return
// 	}
// 	words, images = countWordsAndImages(doc)
// 	return
// }
// func countWordsAndImages(n *html.Node) (words, images int) {
// 	fmt.Println(*n)
// 	return 1, 2
// }

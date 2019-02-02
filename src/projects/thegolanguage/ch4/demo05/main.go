package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	reverse(&arr, len(arr))
	fmt.Println(arr)

}

// 练习 4.3： 重写reverse函数，使用数组指针代替slice。
func reverse(ptr *[6]int, len int) {
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

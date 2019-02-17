package main

import (
	"sort"
	"projects/thegolanguage/ch7/demo06/stringslice"
)

/*
	7.6. sort.Interface接口

	Go语言的sort.Sort函数不会对具体的序列和它的元素做任何假设。相反，它使用了一个接口类型
	sort.Interface来指定通用的排序算法和可能被排序到的序列类型之间的约定。这个接口的实现
	由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。

	type Interface interface {
		// 一个内置的排序算法需要知道三个东
		Len() int	// 序列的长度
		Less(i, j int) bool // 表示两个元素比较的结果
		Swap(i, j int) // 一种交换两个元素的方式
	}
*/

var StringSlice stringslice
func main() {
	sort.Sort(StringSlice([...] names {"justin","emma","mark","lucas","harry"}))
}
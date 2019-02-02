package main

import "fmt"

/*
4.2.2. Slice内存技巧

*/
func main() {
	// 返回不包含空字符串的列表
	out1 := nonempty([]string{"abc", "def", "", "hij", ""})
	out2 := nonempty2([]string{"abc", "def", "", "hij", ""})

	fmt.Println(out1)
	fmt.Println(out2)

	// 一个slice可以用来模拟一个stack
	// 初始化
	stack := []string{"abc", "def", "hij"}
	stack = append(stack, "hello") // push
	fmt.Println(stack)
	top := stack[len(stack)-1] // top of stack
	fmt.Println(top)
	stack = stack[:len(stack)-1] // pop
	fmt.Println(stack)

	// 要删除slice中间的某个元素并保存原有的元素顺序，可以通过内置的copy函数将后面的子slice向前依次移动一位完成：
	out3 := []int{1, 2, 3, 4, 4, 5, 6}
	out4 := remove(out3, 3)
	fmt.Println(out4)
}

// 微妙的地方是，输入的slice和输出的slice共享一个底层数组。这可以避免分配另一个数组，不过原来的数据将可能会被覆盖
func nonempty(s []string) []string {
	i := 0
	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}

	return s[:i]
}

func nonempty2(s []string) []string {
	//
	out := []string{}
	// 等价于
	//out := s[:0]
	for _, v := range s {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

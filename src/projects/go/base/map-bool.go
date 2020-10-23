package main

import "fmt"

func main() {
	// 三种声明并初始化字典的方法

	// 方法1
	var map1 map[string]int = map[string]int{"hello": 1, "map": 2}
	fmt.Printf("%d", map1)
	// 方法2
	map2 := map[string]int{"hello": 1, "map": 2}
	fmt.Printf("%d", map2)
	// 方法2
	map3 := make(map[string]int)
	map3["hello"] = 3
	map3["say"] = 4

	// 第一种方法如果拆分成多步（声明、初始化、再赋值），和其他两种有很大的不一样了，相对会比较麻烦。
	var map4 map[string]int
	map4 = make(map[string]int)
	map4["say"] = 4
	map4["say2"] = 5
	fmt.Println(map3)
	fmt.Println(map4)

	// 删除元素
	delete(map4, "say2")
	fmt.Println(map4)

	// 判断 key 是否存在
	fmt.Println(map4["say2"]) //返回0值，因此不能通过此方法判断
	// 正确判断方式
	if value, ok := map4["say"]; ok {
		fmt.Println(value);
	} else {
		fmt.Println("the key say not exists")
	}
	if value, ok := map4["say2"]; ok {
		fmt.Println(value);
	} else {
		fmt.Println("the key say2 not exists")
	}

	// 遍历字典
	scores := map[string]int {"english":80, "math":100}
	for key, value := range scores {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
	// 遍历只获取key，这里注意不用占用符。
	for key := range scores {
		fmt.Printf("key: %s\n", key)
	}

	// 只获取 value，用一个占位符替代。
	for _, value := range scores {
		fmt.Printf("value: %d\n", value)
	}

	// boolean
	// 而在 Go 中，真值用 true 表示，不但不与 1 相等，并且更加严格，不同类型无法进行比较，而假值用 false 表示，同样与 0 无法比较。
	// bool 转 int
}

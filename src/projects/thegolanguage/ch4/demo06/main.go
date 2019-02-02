package main

import (
	"fmt"
	"sort"
)

/*
4.3. Map

map类型可以写为map[K]V，其中K和V分别对应key和value。map中所有的key都有相同的类型，所有的value也有着相同的类型，
但是key和value之间可以是不同的数据类型。其中K对应的key必须是支持==比较运算符的数据类型，所以map可以通过测试key是否相等来判断是否已经存在。
*/

func main() {
	// 创建map
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	// 创建map并初始化
	address := map[string]string{
		"justin": "shanghai",
		"kevin":  "beijing",
	}

	//创建map
	statures := map[string]float32{}
	fmt.Println(ages)
	fmt.Println(address)
	fmt.Println(statures)

	delete(ages, "alice")
	ages["bob"] = ages["bob"] + 1
	ages["bob"]++
	// map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作
	// _ = &ages["bob"] // compile error: cannot take address of map element

	/*
	 遍历map
	 1. 遍历的顺序是随机的，遍历的顺序是随机的
	*/
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// 2. 如果要按顺序遍历key/value对，我们必须显式地对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序。
	var names []string
	// 从ages map中取出所有name
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	// map类型的零值是nil，也就是没有引用任何哈希表。大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上, 但赋值不行
	var agess map[string]int
	fmt.Println(agess == nil)    // "true"
	fmt.Println(len(agess) == 0) // "true"

	// 判断一个key是否存在
	if _, ok := ages["bob"]; !ok {
		fmt.Println("the key not exists")
	}

	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
}

// 和slice一样，map之间也不能进行相等比较；唯一的例外是和nil进行比较。要判断两个map是否包含相同的key和value，我们必须通过一个循环实现：
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

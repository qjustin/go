package main

import (
	"fmt"
	"time"
)

/*
4.4. 结构体
结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。

1. 结构体成员的输入顺序也有重要的意义。字段顺序不一样，就是两个不同得结构体
2. 成员名字以大写开头，就是导出成员，结构体成员的输入顺序也有重要的意义。
3. 一个命名为S的结构体类型将不能再包含S类型的成员，但是S类型的结构体可以包含*S指针类型的成员
	但是S类型的结构体可以包含*S指针类型的成员
4. 结构体类型的零值是每个成员都是零值。

完整的结构体写法通常只在类型声明语句的地方出现，就像Employee类型声明语句那样。
*/

// 定义结构
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	dilbert.Salary -= 5000
	position := &dilbert.Position
	*position = "Senior " + *position

	// employeeOfTheMonth指针，指向一个Employee类型的结构
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	// 等价于
	(*employeeOfTheMonth).Position += " (proactive team player)"

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"

	id := dilbert.ID
	// 如果将EmployeeByID函数的返回值从*Employee指针类型改为Employee值类型，那么更新语句将不能编译通过，
	// 因为在赋值语句的左边并不确定是一个变量（译注：调用函数返回的是值，并不是一个可取地址的变量）。
	EmployeeByID(id).Salary = 0 // fired for... no real reason

	// 如果结构体没有任何成员的话就是空结构体，写作struct{}。它的大小为0，也不包含任何信息，但是有时候依然是有价值的。
	// 有些Go语言程序员用map来模拟set数据结构时，用它来代替map中布尔类型的value，只是强调key的重要性，但是因为节约的
	// 空间有限，而且语法比较复杂，所以我们通常会避免这样的用法。
	seen := make(map[string]struct{}) // set of strings
	if _, ok := seen["hello"]; !ok {
		seen["hello"] = struct{}{}

	}
}

func EmployeeByID(id int) *Employee {
	return &dilbert
}

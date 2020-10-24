package main

import "fmt"

type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针
}
/**
1 两种定义方法的方式：
1.1. 以值做为方法接收者
1.2. 以指针做为方法接收者

处于性能，代码一致性，建议都使用指针做为接收者。
 */

// 1.1 定义方法：答案是可以使用组合函数的方式来定义结构体方法
// 其中FmtProfile 是方法名，
// (person Profile) ：表示将 FmtProfile 方法与 Profile 的实例绑定。
// 我们把 Profile 称为方法的接收者，而 person 表示实例本身，在方法内可以使用 person.属性名 的方法来访问实例属性。
// 以值做为方法接收者（修改副本）
func (person Profile) FmtProfile() {
	person.age += 1
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%s\n", person.gender)
}
// 1.2 你想要在方法内改变实例(myself.IncAge() 的 myself对象)的属性的时候，必须使用指针做为方法的接收者。
// 以指针做为方法接收者（修改对象本身）
func (person *Profile) IncAge() {
	person.age += 1
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%s\n", person.gender)
}

/**
	2. 结构体实现 “继承”，
	Go 语言本身并不支持继承。但我们可以使用组合的方法，实现类似继承的效果。
 */

type company struct {
	companyName string
	companyAddr string
}

type staff struct {
	name string
	age int
	gender string
	position string
	company
}

/**
	3. 内部方法与外部方法
	当方法的首字母为大写时，这个方法对于所有包都是Public，其他包可以随意调用
	当方法的首字母为小写时，这个方法是Private，其他包是无法访问的。
 */
func main() {
	// Go 语言中没有没有 class 类的概念，只有 struct 结构体的概念，因此也没有继承, 无法在结构体内定义方法
	myself := Profile{name: "tony", age: 24, gender: "male"}
	myself.FmtProfile()
	fmt.Printf("年龄：%d\n", myself.age)
	myself.IncAge()
	fmt.Printf("年龄：%d\n", myself.age)

	myCom := company{
		companyName: "Tencent",
		companyAddr: "深圳市南山区",
	}
	staffInfo := staff{
		name:     "小明",
		age:      28,
		gender:   "男",
		position: "云计算开发工程师",
		company: myCom,
	}

	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyName)
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyName)
}

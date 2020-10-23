package main

import "fmt"

func main() {
	education := "本科"

	// 1. 最简单的示例
	switch education {
	case "博士":
		fmt.Println("我是博士")
	case "研究生":
		fmt.Println("我是研究生")
	case "本科":
		fmt.Println("我是本科生")
	case "大专":
		fmt.Println("我是大专生")
	case "高中":
		fmt.Println("我是高中生")
	default:
		fmt.Println("学历未达标..")
	}

	// 2. 一个 case 多个条件（多个条件之间是 或 的关系）
	month := 2
	switch month {
	case 3, 4, 5:
		fmt.Println("春天")
	case 6, 7, 8:
		fmt.Println("夏天")
	case 9, 10, 11:
		fmt.Println("秋天")
	case 12, 1, 2:
		fmt.Println("冬天")
	default:
		fmt.Println("输入有误...")
	}

	// 3. case 条件常量不能重复
	//gender := "male"
	//
	//switch gender {
	//case "male":
	//	fmt.Println("男性")
	//// 与上面重复
	//case "male":
	//	fmt.Println("男性")
	//case "female":
	//	fmt.Println("女性")
	//}

	// 4. switch 后可接函数
	chinese := 80
	english := 50
	math := 100

	switch getResult(chinese, english, math) {
	// case 后也必须 是布尔类型
	case true:
		fmt.Println("该同学所有成绩都合格")
	case false:
		fmt.Println("该同学有挂科记录")
	}

	// 5. switch 可不接表达式
	// switch 后可以不接任何变量、表达式、函数。
	// 当不接任何东西时，switch - case 就相当于 if - elseif - else
	score := 30

	switch {
	case score >= 95 && score <= 100:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("合格")
	case score >= 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入有误...")
	}

	// 6. switch 的穿透能力:fallthrough
	// 正常情况下 switch - case 的执行顺序是：只要有一个 case 满足条件，
	// 就会直接退出 switch - case ，如果 一个都没有满足，才会执行 default 的代码块。
	s := "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
		// 需要注意的是，fallthrough 只能穿透一层，意思是它让你直接执行下一个case的语句，而且"不需要判断条件"。
	case s == "world":
		fmt.Println("world")
	}
}

func getResult(args ...int) bool {
	for _, i := range args {
		if i < 60 {
			return false
		}
	}
	return true
}

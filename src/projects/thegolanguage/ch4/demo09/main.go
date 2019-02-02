package main

import (
	"fmt"
)
/*
	4.4.2. 结构体比较

	1. 用同一个结构体声明的两个变量，如果结构体的全部成员都是可以比较的，
		那么这两个结构体变量也是可以比较的
	2. 结构体的所有成员相等，结构体才相等
	3. 不同的结构体变量不能直接比较
	4. 结构体可用于map的key类型
	
*/

type Point struct{ X, Y int }
type PointB  struct{ X, Y int }

type address struct {
	hostname string
	port     int
}
func main() {
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q)                   // "false"
	// 结构体的所有成员相等，结构体才相等
	p1 := Point{1, 2}
	q1 := Point{1, 2}
	fmt.Println(p1.X == q1.X && p1.Y == q1.Y) 	// "true"
	fmt.Println(p1 == q1)                   	// "true"

	/* 
		不同的结构体变量不能直接比较
		p2 := Point{1, 2}
		q2 := PointB{1, 2}
		fmt.Println(p2 == q2) // 编译错误
	*/

	// 结构体可用于map的key类型
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}
package main

import (
	"fmt"
	"log"
	"encoding/json"
)

/*

Go语言对于这些标准格式的编码和解码都有良好的支持，由标准库中的
encoding/json、
encoding/xml、
encoding/asn1等包提供支持
（译注：Protocol Buffers的支持由 github.com/golang/protobuf 包提供）

一个JSON数组是一个有序的值序列，写在一个方括号中并以逗号分隔；
一个JSON数组可以用于编码Go语言的数组和slice。
一个JSON对象是一个字符串到值的映射，写成一系列的name:value对形式，用花括号包含并以逗号分隔；
JSON的对象类型可以用于编码Go语言的map类型（key类型是字符串）和结构体。


MarshalIndent方法底层通过reflect反射技术来转换

*/

// 只有导出的结构体成员才会被编码，这也就是我们为什么选择用大写字母开头的成员名称。
type Movie struct {
	Title  string
	// json开头键名对应的值用于控制encoding/json包的编码和解码的行为，并且encoding/...下面其它的包也遵循这个约定。
	// 当用json编码时 Year字段 输出为release
	Year   int  `json:"released"`   
	/*
	当用json编码时 Color字段 输出为color，并且如果color值未false 或 0 时不输出该字段
	Color成员的Tag还带了一个额外的omitempty选项，表示当Go语言结构体成员为空或零值时不生成该JSON对象（这里false为零值）。果然，Casablanca是一个黑白电影，并没有输出Color成员。Color成员的Tag还带了一个额外的omitempty选项，表示当Go语言结构体成员为空或零值时不生成该JSON对象（这里false为零值）。果然，Casablanca是一个黑白电影，并没有输出Color成员。
	*/
	Color  bool `json:"color,omitempty"`  
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	// 将结构体转换成JSON

	// Marshal 输出紧凑格式的JSON 不易于阅读
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	// MarshalIndent 输出带缩进格式的JSON。参数用于表示每一行输出的前缀和每一个层级的缩进
	data1, err1 := json.MarshalIndent(movies, "", "    ")
	if err1 != nil {
		log.Fatalf("JSON marshaling failed: %s", err1)
	}
	fmt.Printf("%s\n", data1)


	// 将JSON逆编码为结构体
	// 定义一个结构体切片，结构体中只有Title成员
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
}
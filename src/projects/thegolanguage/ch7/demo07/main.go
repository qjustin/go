package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
	7.7. http.Handler接口

	1. 单路Handler接口： ListenAndServe 函数监听请求，并将所有请求转发给 handler来处理
	package http
	type Handler interface {
		ServeHTTP(w ResponseWriter, r *Request)
	}
	func ListenAndServe(address string, h Handler) error

	2. 多路器ServeMux： net/http包提供了一个请求多路器ServeMux来简化URL和handlers的联系。 一个ServeMux将一批
						http.Handler聚集到一个单一的http.Handler中。再一次，我们可以看到满足同一接口的不同类型
						是可替换的：web服务器将请求指派给任意的http.Handler 而不需要考虑它后面的具体类型。

	3. 全局多路器ServeMux： net/http包提供了一个全局的ServeMux实例DefaultServerMux和包级别的http.Handle和
						   http.HandleFunc函数。现在，为了使用DefaultServeMux作为服务器的主handler，我们
						   不需要将它传给ListenAndServe函数；nil值就可以工作。
*/

/*
	1. 单路Handler接口
	dbColumns 类型的底层数据结构是一个map,key 是列名，value列存储类型,
	dbColumns 类型实现了Handler接口
*/
type dbColumns map[string]string

func (cols dbColumns) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for colName, typeName := range cols {
		fmt.Printf("%s: %s\n", colName, typeName)
	}
}

/*
	2. 多路器ServeMux
	dbColumns 类型的底层数据结构是一个map,key 是列名，value列存储类型,
	dbColumns 类型实现了Handler接口
*/

func (cols dbColumns) printColName(w http.ResponseWriter, req *http.Request) {
	for colName, _ := range cols {
		fmt.Print(colName)
	}
}
func (cols dbColumns) printColType(w http.ResponseWriter, req *http.Request) {
	for _, typeName := range cols {
		fmt.Print(typeName)
	}
}

/*
	3. 全局多路器ServeMux
	dbColumns 类型的底层数据结构是一个map,key 是列名，value列存储类型,
	dbColumns 类型实现了Handler接口
*/

func main() {
	// 1. 单路Handler接口
	cols := dbColumns{"username": "string", "userid": "bigint"}
	log.Fatal(http.ListenAndServe("localhost:8000", cols))

	// 2. 多路器ServeMux
	mux := http.NewServeMux()
	mux.Handle("/name", http.HandlerFunc(cols.printColName))
	mux.Handle("/type", http.HandlerFunc(cols.printColType))
	log.Fatal(http.ListenAndServe("localhost:8001", mux))

	// 3. 全局多路器ServeMux
	http.HandleFunc("/name", cols.printColName)
	http.HandleFunc("/type", cols.printColType)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

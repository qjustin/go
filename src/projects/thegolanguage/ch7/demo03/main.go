package main

import (
	"fmt"
	"flag"
	"time"
)

/*
7.4. flag.Value接口
1. 因为时间周期标记值非常的有用，所以这个特性被构建到了flag包中
2. 为自己的数据类型定义新的标记符号是简单容易的。只需要定义一个实现
flag.Value接口的类型. 参考：flag.go

flag包实现了命令行参数的解析。每个参数认为一条记录，根据实际进行定义，到一个set集合。每条都有各自的状态参数。

在使用flag时正常流程： 

1. 通过flag.String(), flag.Bool(), flag.Int()等函数来定义命令行中需要使用的参数。

2. 在定义完flag后，通过调用flag.Parse()来进行对命令行参数的解析。

3. 获取flag.String(), flag.Bool(), flag.Int()等方法的返回值,即对应用户输入的参数.

　注意的是flag.Xxx()返回的值是变量的内存地址,要获取值时要通过在变量前加*(星号)获取.
*/

//时间周期标记值 go run main.go -period 50ms
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()

	printFlag()
}


func printFlag() {
	// golang的flag包的一些基本使用方法

    // 待使用的变量
    var id int
    var name string
    var male bool

    // 是否已经解析
    fmt.Println("parsed? = ", flag.Parsed())

    // 设置flag参数 (变量指针，参数名，默认值，帮助信息)
    // 也可以用以下带返回值的方法代替，不过他们返回的是指针，比较麻烦点
    // Int(name string, value int, usage string) *int
    // String(name string, value string, usage string) *string
    // Bool(name string, value bool, usage string) *bool
    flag.IntVar(&id, "id", 123, "help msg for id")
    flag.StringVar(&name, "name", "default name", "help msg for name")
    flag.BoolVar(&male, "male", false, "help msg for male")

    // 解析
    flag.Parse()

    // 是否已经解析
    fmt.Println("parsed? = ", flag.Parsed())

    //

    // 获取非flag参数
    fmt.Println(flag.NArg())
    fmt.Println("------ Args start ------")
    fmt.Println(flag.Args())
    for i, v := range flag.Args() {
        fmt.Printf("arg[%d] = (%s).\n", i, v)
    }
    fmt.Println("------ Args end ------")
    //
    // visit只包含已经设置了的flag
    fmt.Println("------ visit flag start ------")
    flag.Visit(func(f *flag.Flag) {
        fmt.Println(f.Name, f.Value, f.Usage, f.DefValue)
    })
    fmt.Println("------ visit flag end ------")
    //
    // visitAll只包含所有的flag(包括未设置的)
    fmt.Println("------ visitAll flag start ------")
    flag.VisitAll(func(f *flag.Flag) {
        fmt.Println(f.Name, f.Value, f.Usage, f.DefValue)

    })
    fmt.Println("------ visitAll flag end ------")
    //
    // flag参数
    fmt.Printf("id = %d\n", id)
    fmt.Printf("name = %s\n", name)
    fmt.Printf("male = %t\n", male)

    // flag参数默认值
    fmt.Println("------ PrintDefaults start ------")
    flag.PrintDefaults()
    fmt.Println("------ PrintDefaults end ------")

     //非flag参数个数
    fmt.Printf("NArg = %d\n", flag.NArg())
    // 已设置的flag参数个数
    fmt.Printf("NFlag = %d\n", flag.NFlag())
}
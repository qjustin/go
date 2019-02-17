package customflag

import (
	"flag"
    "fmt"
    "strings"
)

/*
https://www.jianshu.com/p/f9cf46a4de0e

1.2 flag 包概述
flag 包实现了命令行参数的解析。
1.2.1 定义 flags 有两种方式
1）flag.Xxx()，其中 Xxx 可以是 Int、String，Bool 等；返回一个相应类型的指针，如：
var ip = flag.Int("flagname", 1234, "help message for flagname")

第一个参数 ：flag名称为flagname
第二个参数 ：flagname默认值为1234
第三个参数 ：flagname的提示信息

返回的ip是指针类型，所以这种方式获取ip的值应该fmt.Println(*ip)
2）flag.XxxVar()，将 flag 绑定到一个变量上，如：
var flagValue int
flag.IntVar(&flagValue, "flagname", 1234, "help message for flagname")


第一个参数 ：接收flagname的实际值的
第二个参数 ：flag名称为flagname
第三个参数 ：flagname默认值为1234
第四个参数 ：flagname的提示信息
这种方式获取ip的值fmt.Println(ip)就可以了：

*/

/*
1.2.2 自定义 Value
另外，还可以创建自定义 flag，只要实现 flag.Value 接口即可（要求 receiver 是指针），这时候可以通过如下方式定义该 flag：
type Value interface {
	String() string
	Set(string) error
}

我们希望直接解析到 slice 中，我们可以定义如下 sliceValue类型，然后实现Value接口
*/

//定义一个类型，用于增加该类型方法
type sliceValue []string

//new一个存放命令行参数值的slice
func newSliceValue(vals []string, p *[]string) *sliceValue {
    *p = vals
    return (*sliceValue)(p)
}

/*
Value接口：
type Value interface {
    String() string
    Set(string) error
}
实现flag包中的Value接口，将命令行接收到的值用,分隔存到slice里
*/
func (s *sliceValue) Set(val string) error {
    *s = sliceValue(strings.Split(val, ","))
    return nil
}

//flag为slice的默认值default is me,和return返回值没有关系
func (s *sliceValue) String() string {
    *s = sliceValue(strings.Split("default is me", ","))
    return "It's none of my business"
}

/*
可执行文件名 -slice="java,go"  最后将输出[java,go]
可执行文件名 最后将输出[default is me]
 */
func main(){
    var languages []string
    flag.Var(newSliceValue([]string{}, &languages), "slice", "I like programming `languages`")
    flag.Parse()

    //打印结果slice接收到的值
    fmt.Println(languages)
}
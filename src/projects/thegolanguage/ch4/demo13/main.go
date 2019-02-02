package main

import (
	"html/template"
	"log"
	"os"
	"projects/thegolanguage/ch4/github"
	"time"
)

/*
4.6. 文本和HTML模板
text/template和html/template等模板包，它们提供了一个将变量值填充到一个文本或HTML格式的模板的机制

模板：
模板是一个字符串或一个文件，里面包含了一个或多个由双花括号包含的{{action}}对象
大部分的字符串只是按字面值打印，但是对于actions部分将触发其它的行为

每个actions都包含了一个用模板语言书写的表达式

这个表达式可以是以下几种行为：
1. 选择结构体成员
2. 调用函数或方法
3. if else 或 range循环语句等
4. 其它实例化模板

*/
const templ = `{{.TotalCount}} issues:
					{{range .Items}}----------------------------------------
					Number: {{.Number}}
					User:   {{.User.Login}}
					Title:  {{.Title | printf "%.64s"}}
					Age:    {{.CreatedAt | daysAgo}} days
					{{end}}`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

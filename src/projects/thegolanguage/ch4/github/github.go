package github

import "time"
/*
许多web服务都提供JSON接口，通过HTTP接口发送JSON格式请求并返回JSON格式的信息。
为了说明这一点，我们通过Github的issue查询服务来演示类似的用法。
首先，我们要定义合适的类型和常量：
*/
const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items          []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
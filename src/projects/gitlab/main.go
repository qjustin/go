package main


import (
"fmt"
"io/ioutil"
"net/http"
"reflect"
"strings"
)



type Person struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

func (c *BranchController) GetBranch() {
	// http://192.168.110.53/help/api/README.md
	url := "http://192.168.110.53/api/v3/projects?per_page=5000&search=trade"
	req, _ := http.NewRequest("GET", url, nil)
	// 输入自己的gitlab token
	req.Header.Add("PRIVATE-TOKEN", "iHyyvjVVKQM2fcyfu1A-")

	// q := req.URL.Query()
	// q.Add("search", "YJK-Java")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(reflect.TypeOf(body))
	str1 := string(body)
	// 这里本小白 不知道怎么处理json外面带的[] 有大佬记得留意下代码哦
	str2 := strings.Trim(str1, "[")
	str3 := strings.Trim(str2, "]")
	//fmt.Println(str3)
	fmt.Println(reflect.TypeOf(str3))

	// const json = `{"id":16,"description":"","default_branch":"master","tag_list":[],"public":false,"archived":false,"visibility_level":0,"ssh_url_to_repo":"git@gitlab.ihaozhuo.com:Java_Service/YJK-Java.git","http_url_to_repo":"http://gitlab.ihaozhuo.com/Java_Service/YJK-Java.git","web_url":"http://gitlab.ihaozhuo.com/Java_Service/YJK-Java","name":"YJK-Java","name_with_namespace":"Java_Service / YJK-Java","path":"YJK-Java","path_with_namespace":"Java_Service/YJK-Java","issues_enabled":true,"merge_requests_enabled":true,"wiki_enabled":true,"snippets_enabled":false,"created_at":"2015-11-04T01:34:40.000Z","last_activity_at":"2020-09-14T02:06:49.000Z","creator_id":4,"namespace":{"id":11,"name":"Java_Service","path":"Java_Service","owner_id":null,"created_at":"2015-10-26T08:48:47.000Z","updated_at":"2015-10-26T08:48:47.000Z","description":"","avatar":{"url":null}},"avatar_url":null}`
	// 	name := gojsonq.New().FromString(json).Find("id")

	name := gojsonq.New().FromString(str3).Find("id")
	pro_id := int64(name.(float64))
	println(pro_id)

	c.Ctx.WriteString("获取配置成功")
}
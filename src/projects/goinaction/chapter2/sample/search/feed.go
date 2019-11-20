package search

import (
	"encoding/json"
	"os"
)

const dataFile = "D://Dev//SourceCode//GITHUB//go//src//projects//goinaction//chapter2//sample//data//data.json"

// Feed 结构对外暴露，
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// func 方法名(参数列表) (返回值1， 返回值2) {}
func RetrieveFeeds() ([]*Feed, error) {
	// file指针，err
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// defer 确保函数返回时才执行 类似finally
	defer file.Close()

	// 切片的每一项时一个指向Feed类型值的指针
	var feeds []*Feed
	// decode 接收一个 interface{} 这个类型在go中很特殊，
	// 一般会配合reflect包里提供的反射一起使用
	err = json.NewDecoder(file).Decode(&feeds)

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}

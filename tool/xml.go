package tool

import (
	"asitop-go/model"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mitchellh/mapstructure"
	"howett.net/plist"
)

// 用来单体测试结构体
func ElpanXml(path string, v *map[string]interface{}) (err error) {
	// 使用os.Open打开文件，获取*os.File类型
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening XML file: %s\n", err)
		return
	}
	defer file.Close() // 确保在函数返回时关闭文件

	// 使用xml.NewDecoder和Decode方法解析文件内容到结构体
	dec := plist.NewDecoder(file)
	err = dec.Decode(v)
	if err != nil {
		fmt.Printf("Error parsing XML: %s\n", err)
	}
	return
}

// map => 结构体
func MapToPcInfo(v *map[string]interface{}, info *model.PcInfo) error {
	return mapstructure.Decode(v, info)
}

// map => json => struct
func MapJsonPcInfo(v *map[string]interface{}, info *model.PcInfo) error {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, info)
}

func MapToPcShowInfo(v *map[string]interface{}, info *model.PcShowInfo) error {
	return mapstructure.Decode(v, info)
}

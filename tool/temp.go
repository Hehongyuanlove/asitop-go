package tool

import (
	"asitop-go/config"
	"fmt"
	"os"
)

// 清理temp文件
func ClearTemp() (bool, error) {
	if _, err := os.Stat(config.TempFile); os.IsNotExist(err) {
		fmt.Println("文件不存在")
		return true, nil
	}
	if err := os.Remove(config.TempFile); err != nil {
		return false, err
	}
	return false, nil // 薛定谔的文件？
}

// 单纯的执行shell文件

// 解析 xml 文件

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

func main_paln() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Age", "Email", "Date"})

	// 模拟数据更新
	go func() {
		for {
			// 使用ANSI转义序列清空屏幕
			fmt.Print("\033[H\033[2J")

			// 获取当前时间
			currentTime := time.Now()

			// 如果你想要自定义时间格式，可以这样做：
			customFormat := "2006-01-02 15:04:05"
			customTimeStr := currentTime.Format(customFormat)

			// 假设这是从数据库或其他数据源获取的最新数据
			data := [][]string{
				{"1", "Michael", "30", "michael@example.com", customTimeStr},
				{"1", "Michael", "30", "michael@example.com", customTimeStr},
				{"1", "Michael", "30", "michael@example.com", customTimeStr},
				{"1", "Michael", "30", "michael@example.com", customTimeStr},
				// 这里可以添加更多数据
			}

			// 清空表格
			table.ClearRows()

			// 添加新数据
			for _, d := range data {
				table.Append(d)
			}

			// 重新渲染表格
			table.Render()

			// 等待一段时间再更新
			time.Sleep(time.Second)
		}
	}()

	// 等待用户输入以退出程序
	fmt.Println("Press enter to exit...")
	fmt.Scanln()

	// 退出定时器和程序
}

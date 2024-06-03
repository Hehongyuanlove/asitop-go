package service

import (
	"asitop-go/config"
	"asitop-go/model"
	"asitop-go/tool"
	"fmt"
	"os/exec"
	"time"

	"github.com/mitchellh/mapstructure"
)

type pcShowInfoService struct {
	PcShowInfo model.PcShowInfo
	FlashShell bool
}

var PcShowInfoService = new(pcShowInfoService)

func (p *pcShowInfoService) Explan() *pcShowInfoService {
	v := map[string]interface{}{}
	tool.ElpanXml(config.TempFile, &v)

	err := mapstructure.Decode(&v, &p.PcShowInfo)
	if err != nil {
		fmt.Println(err)
	}
	return p
}

func (p *pcShowInfoService) Get() model.PcShowInfo {
	return p.PcShowInfo
}

// 显示在 shell 上
func (p *pcShowInfoService) Show() {
	for {
		if !p.FlashShell {
			continue
		}
		p.Explan()
		// 使用ANSI转义序列清空屏幕
		fmt.Print("\033[H\033[2J")

		// 获取当前时间
		currentTime := time.Now()

		// 如果你想要自定义时间格式，可以这样做：
		customFormat := "2006-01-02 15:04:05"
		customTimeStr := currentTime.Format(customFormat)

		fmt.Println("#######################################")
		fmt.Println("#### ", p.PcShowInfo.HwModel, "......", p.PcShowInfo.KernOSVersion, " ####")
		fmt.Println("####  GPU      ", formatGhz(p.PcShowInfo.GPU.FreqHz), formatPercent(p.PcShowInfo.GPU.IdleRatio), "####")

		clusters := p.PcShowInfo.Processor.Clusters
		for i := range clusters {
			fmt.Println("#### ", clusters[i].Name, formatGhz(clusters[i].FreqHz), formatPercent(clusters[i].IdleRatio), "####")
		}

		fmt.Println("##### ", p.PcShowInfo.TimeStamp, "##")
		fmt.Println("##### ", customTimeStr, "############")
		fmt.Println("#######################################")

		// 等待一段时间再更新
		time.Sleep(time.Second)
	}
}

func (p *pcShowInfoService) ShellStr() string {

	// "sudo nice -n",
	//     str(nice),
	//     "powermetrics",
	//     "--samplers cpu_power,gpu_power,thermal",
	//     output_file_flag,
	//     "/tmp/asitop_powermetrics"+timecode,
	//     "-f plist",
	//     "-i",
	//     str(interval)
	return fmt.Sprintf("sudo nice -n %d powermetrics -s cpu_power,gpu_power,thermal -f plist  -o %s -i %d", 10, config.TempFile, 1000)
}

func (p *pcShowInfoService) Shell() {
	str := fmt.Sprintf("sudo powermetrics -s cpu_power,gpu_power -f plist -n 1 -o %s ", config.TempFile)

	for {
		cmd := exec.Command("sh", "-c", str)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running command:", err)
		}
		p.FlashShell = true
		time.Sleep(time.Second)
	}
}

func formatGhz(f int64) string {
	return fmt.Sprintf(" %.2f Ghz", float64(f/1e9))
}

func formatPercent(f float64) string {
	return fmt.Sprintf("  %.2f %%", 1-f)
}

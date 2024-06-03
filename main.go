package main

import "asitop-go/service"

func main() {
	service.PcShowInfoService.Explan()
	go service.PcShowInfoService.Shell()
	service.PcShowInfoService.Show()

}

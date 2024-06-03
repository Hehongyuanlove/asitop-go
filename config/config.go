package config

// temp 文件位置
var TempFile string = "./temp"

// 命令 执行层级 10 最好因为只是收集类 没必要占很多资源
var RunLevel int = 10

// 采集间隔 (ms)
var SamplingFrequency int = 1000

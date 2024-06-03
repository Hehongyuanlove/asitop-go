package main

func plan() {
	// E-CPU % Hz processor.clusters[0].[name, freq_hz  idle_ratio]
	// P-CPU % Hz processor.clusters[1].[name, freq_hz  idle_ratio]
	// GPU % Hz  gpu.freq_hz   1-gpu.idle_ratio
	// ANE % Hz

	// Memory %  swap status
	//

	// 总功率     processor.combined_power
	// CPU功率   processor.cpu_power
	// GPU功率   processor.gpu_power

	// 一个协程维持shell刷新 temp 文件
	// 一个协程读取temp文件 并刷新到模型文件
	// 一个携程监听信号
}

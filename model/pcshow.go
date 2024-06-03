package model

import (
	"time"
)

type PcShowInfo struct {
	IsDelta       bool          `mapstructure:"is_delta" json:"is_delta"`
	ElapsedNS     int64         `mapstructure:"elapsed_ns" json:"elapsed_ns"`
	HwModel       string        `mapstructure:"hw_model" json:"hw_model"`
	KernOSVersion string        `mapstructure:"kern_osversion" json:"kern_osversion"`
	KernBootargs  string        `mapstructure:"kern_bootargs" json:"kern_bootargs"`
	KernBoottime  int64         `mapstructure:"kern_boottime" json:"kern_boottime"`
	TimeStamp     time.Time     `mapstructure:"timestamp" json:"timestamp"`
	GPU           GPUShow       `mapstructure:"gpu" json:"gpu"`
	Processor     ProcessorShow `mapstructure:"processor" json:"processor"`
}

type GPUShow struct {
	FreqHz    int64   `mapstructure:"freq_hz" json:"freq_hz"`       //GPU 的工作频率
	IdleRatio float64 `mapstructure:"idle_ratio" json:"idle_ratio"` // GPU 空闲状态占比
}

type ProcessorShow struct {
	CpuEnergy int64     `mapstructure:"cpu_energy" json:"cpu_energy"`
	Clusters  []Cluster `mapstructure:"clusters" json:"clusters"`
}

type ClusterShow struct {
	Name      string  `mapstructure:"name" json:"name"`
	FreqHz    int64   `mapstructure:"freq_hz" json:"freq_hz"`
	IdleRatio float64 `mapstructure:"idle_ratio" json:"idle_ratio"`
}

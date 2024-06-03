package model

import "time"

// string, bool, uint64, float64

type PcInfo struct {
	IsDelta       bool      `mapstructure:"is_delta" json:"is_delta"`
	ElapsedNS     int64     `mapstructure:"elapsed_ns" json:"elapsed_ns"`
	HwModel       string    `mapstructure:"hw_model" json:"hw_model"`
	KernOSVersion string    `mapstructure:"kern_osversion" json:"kern_osversion"`
	KernBootargs  string    `mapstructure:"kern_bootargs" json:"kern_bootargs"`
	KernBoottime  int64     `mapstructure:"kern_boottime" json:"kern_boottime"`
	TimeStamp     time.Time `mapstructure:"timestamp" json:"timestamp"`
	GPU           GPU       `mapstructure:"gpu" json:"gpu"`
	Desk          Desk      `mapstructure:"disk" json:"disk"`
	Network       Network   `mapstructure:"network" json:"network"`
	Processor     Processor `mapstructure:"processor" json:"processor"`
}

type Desk struct {
	RbytesDiff int64   `mapstructure:"rbytes_diff" json:"rbytes_diff"`
	RbytesPers float64 `mapstructure:"rbytes_per_s" json:"rbytes_per_s"`
	RopsDiff   int64   `mapstructure:"rops_diff" json:"rops_diff"`
	RopsPers   float64 `mapstructure:"rops_per_s" json:"rops_per_s"`
	WbytesDiff int64   `mapstructure:"wbytes_diff" json:"wbytes_diff"`
	WbytesPers float64 `mapstructure:"wbytes_per_s" json:"wbytes_per_s"`
	WopsDiff   int64   `mapstructure:"wops_diff" json:"wops_diff"`
	WopsPers   float64 `mapstructure:"wops_per_s" json:"wops_per_s"`
}

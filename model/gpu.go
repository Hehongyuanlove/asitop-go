package model

type GPU struct {
	FreqHz            int64              `mapstructure:"freq_hz" json:"freq_hz"`                       //GPU 的工作频率
	GpuEnergy         int64              `mapstructure:"gpu_energy" json:"gpu_energy"`                 //GPU 消耗的能量
	IdleNs            int64              `mapstructure:"idle_ns" json:"idle_ns"`                       // GPU 处于空闲状态的时间
	IdleRatio         float64            `mapstructure:"idle_ratio" json:"idle_ratio"`                 // GPU 空闲状态占比
	DvfmStates        []DvfmState        `mapstructure:"dvfm_states" json:"dvfm_states"`               // GPU 几种工作频率
	SwRequestedStates []SwRequestedState `mapstructure:"sw_requested_state" json:"sw_requested_state"` // GPU 几种软件请求状态
	SwStates          []SwStateBase      `mapstructure:"sw_state" json:"sw_state"`                     // GPU 几种软件状态
}

// GPG 工作频率状态
// m1 有六种 396 528 720 924 1128 1278Hz
type DvfmState struct {
	Freq      int64   `mapstructure:"freq" json:"freq"`
	UsedNs    int64   `mapstructure:"used_ns" json:"used_ns"`
	UsedRatio float64 `mapstructure:"used_ratio" json:"used_ratio"`
}

// 软件请求状态 ？
// m1 六种 P1 => P6
type SwRequestedState struct {
	SwReqState string  `mapstructure:"sw_req_state" json:"sw_req_state"`
	UsedNs     int64   `mapstructure:"used_ns" json:"used_ns"`
	UsedRatio  float64 `mapstructure:"used_ratio" json:"used_ratio"`
}

// 软件状态 ？
// m1 六种 SW_P1 => SW_P6
type SwStateBase struct {
	SwState   string  `mapstructure:"sw_state" json:"sw_state"`
	UsedNs    int64   `mapstructure:"used_ns" json:"used_ns"`
	UsedRatio float64 `mapstructure:"used_ratio" json:"used_ratio"`
}

package model

type Processor struct {
	AneEnery             int64     `mapstructure:"ane_enery" json:"ane_enery"`
	CpuEnergy            int64     `mapstructure:"cpu_energy" json:"cpu_energy"`
	GpuEnergy            int64     `mapstructure:"gpu_energy" json:"gpu_energy"`
	AnePower             float64   `mapstructure:"ane_power" json:"ane_power"`           // ANE 功率
	CpuPower             float64   `mapstructure:"cpu_power" json:"cpu_power"`           // CPU 功率
	GpuPower             float64   `mapstructure:"gpu_power" json:"gpu_power"`           // GPU 功率
	CombinedPower        float64   `mapstructure:"combined_power" json:"combined_power"` // 总功率
	CpuPowerZonesEngaged float64   `mapstructure:"cpu_power_zones_engaged" json:"cpu_power_zones_engaged"`
	Clusters             []Cluster `mapstructure:"clusters" json:"clusters"`
}

// E-Cluster or P-Cluster
type Cluster struct {
	Name             string  `mapstructure:"name" json:"name"`
	FreqHz           int64   `mapstructure:"freq_hz" json:"freq_hz"`
	HwResidCounters  bool    `mapstructure:"hw_resid_counters" json:"hw_resid_counters"`
	IdleNs           int64   `mapstructure:"idle_ns" json:"idle_ns"`
	IdleRatio        float64 `mapstructure:"idle_ratio" json:"idle_ratio"`
	RecommendedCores int64   `mapstructure:"recommended_cores" json:"recommended_cores"`
	RequestedMhz     int64   `mapstructure:"requested_mhz" json:"requested_mhz"`
}

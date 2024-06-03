package tool

import (
	"asitop-go/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// xml 判断 GPU	 解析是否正确
func Test_Explan_Model_GPU(t *testing.T) {
	v := map[string]interface{}{}
	info := model.PcInfo{}
	err := ElpanXml("../simple/all.xml", &v)
	assert.Equal(t, err, nil, "ElpanXmlAny")

	err = MapToPcInfo(&v, &info)
	assert.Equal(t, err, nil, "MapToPcInfo")

	// 判断 GPU
	gpu := info.GPU
	assert.Equal(t, gpu.FreqHz, int64(396), "gpu.FreqHz")
	assert.Equal(t, gpu.GpuEnergy, int64(122), "gpu.GpuEnergy")
	assert.Equal(t, gpu.IdleNs, int64(4391487375), "gpu.IdleNs")
	assert.Equal(t, gpu.IdleRatio, float64(0.871448), "gpu.IdleRatio")
	assert.Equal(t, gpu.DvfmStates[5].Freq, int64(1278), "gpu.DvfmStates[5].Freq")
	assert.Equal(t, gpu.SwRequestedStates[5].SwReqState, "P6", "gpu.SwRequestedStates[5].SwReqState")
	assert.Equal(t, gpu.SwStates[5].SwState, "SW_P6", "gpu.SwStates[5].SwState")
}

func Test_Processor(t *testing.T) {
	v := map[string]interface{}{}
	info := model.PcInfo{}
	err := ElpanXml("../simple/all.xml", &v)
	assert.Equal(t, err, nil, "ElpanXmlAny")

	err = MapToPcInfo(&v, &info)
	assert.Equal(t, err, nil, "MapToPcInfo")

	pro := info.Processor
	assert.Equal(t, pro.CombinedPower, float64(260.644), "pro.CombinedPower")
	assert.Equal(t, pro.CpuPower, float64(236.352), "pro.CpuPower")
	assert.Equal(t, pro.GpuPower, float64(24.2923), "pro.GpuPower")

	assert.Equal(t, pro.Clusters[0].FreqHz, int64(1214100000), "pro.Clusters[0].FreqHz")
	assert.Equal(t, pro.Clusters[0].IdleNs, int64(2881572791), "pro.Clusters[0].IdleNs")
	assert.Equal(t, pro.Clusters[0].IdleRatio, float64(0.571996), "pro.Clusters[0].IdleRatio")

}

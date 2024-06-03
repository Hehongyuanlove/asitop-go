package tool

import (
	"asitop-go/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// xml 头部解析
func Test_ExplanXml(t *testing.T) {
	v := map[string]interface{}{}
	err := ElpanXml("../simple/network.xml", &v)
	assert.Equal(t, err, nil, "ElpanXmlAny")
	// t.Log(v)
	assert.Equal(t, v["elapsed_ns"], uint64(5001102250), "elapsed_ns")
	assert.Equal(t, v["kern_osversion"], "23D56", "kern_osversion")
	assert.Equal(t, v["kern_boottime"], uint64(1715615090), "kern_boottime")
	if network, ok := v["network"].(map[string]interface{}); ok {
		assert.Equal(t, network["opackets"], uint64(70), "network.opackets")
	} else {
		t.Error(v["network"])
	}

}

// map => struct
func Test_MapToPcInfo(t *testing.T) {
	v := map[string]interface{}{}
	info := model.PcInfo{}
	err := ElpanXml("../simple/network.xml", &v)
	assert.Equal(t, err, nil, "ElpanXmlAny")

	err = MapToPcInfo(&v, &info)
	assert.Equal(t, err, nil, "MapToPcInfo")

	assert.Equal(t, v["elapsed_ns"], uint64(5001102250), "elapsed_ns")
	assert.Equal(t, info.ElapsedNS, int64(5001102250), "elapsed_ns")
}

// map => struct
func Test_MapJsonPcInfo(t *testing.T) {
	v := map[string]interface{}{}
	info := model.PcInfo{}
	err := ElpanXml("../simple/network.xml", &v)
	assert.Equal(t, err, nil, "ElpanXmlAny")

	err = MapJsonPcInfo(&v, &info)
	assert.Equal(t, err, nil, "MapJsonPcInfo")

	assert.Equal(t, v["elapsed_ns"], uint64(5001102250), "elapsed_ns")
	assert.Equal(t, info.ElapsedNS, int64(5001102250), "elapsed_ns")
}

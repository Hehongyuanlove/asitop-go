package tool

import (
	"asitop-go/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// xml 判断 Desk 解析是否正确
func Test_Explan_Model_Desk(t *testing.T) {
	v := map[string]interface{}{}
	info := model.PcInfo{}
	err := ElpanXml("../simple/all.xml", &v)
	assert.Equal(t, err, nil, "ElpanXmlAny")

	err = MapToPcInfo(&v, &info)
	assert.Equal(t, err, nil, "MapToPcInfo")

	// 判断 Desk
	desk := info.Desk
	assert.Equal(t, desk.RbytesDiff, int64(327680), "desk.RbytesDiff")
	assert.Equal(t, desk.RbytesPers, float64(65246.6), "desk.RbytesPers")
	assert.Equal(t, desk.RopsDiff, int64(16), "desk.RopsDiff")
	assert.Equal(t, desk.RopsPers, float64(3.18587), "desk.RopsPers")
	assert.Equal(t, desk.WbytesDiff, int64(278528), "desk.WbytesDiff")
	assert.Equal(t, desk.WbytesPers, float64(55459.6), "desk.WbytesPers")
	assert.Equal(t, desk.WopsDiff, int64(49), "desk.WopsDiff")
	assert.Equal(t, desk.WopsPers, float64(9.75673), "desk.WopsPers")
}

// xml 判断 Network 解析是否正确
func Test_Explan_Model_Network(t *testing.T) {
	v := map[string]interface{}{}
	info := model.PcInfo{}
	err := ElpanXml("../simple/all.xml", &v)
	assert.Equal(t, err, nil, "ElpanXmlAny")

	err = MapToPcInfo(&v, &info)
	assert.Equal(t, err, nil, "MapToPcInfo")

	// 判断 Network
	network := info.Network
	assert.Equal(t, network.IbyteRate, float64(1253.24), "network.IbyteRate")
	assert.Equal(t, network.Ibytes, int64(6294), "network.Ibytes")
	assert.Equal(t, network.IpacketRate, float64(24.2923), "network.IpacketRate")
	assert.Equal(t, network.Ipackets, int64(122), "network.Ipackets")
	assert.Equal(t, network.ObyteRate, float64(1170.01), "network.ObyteRate")
	assert.Equal(t, network.Obytes, int64(5876), "network.Obytes")
	assert.Equal(t, network.OpacketRate, float64(22.6993), "network.OpacketRate")
	assert.Equal(t, network.Opackets, int64(114), "network.Opackets")
}

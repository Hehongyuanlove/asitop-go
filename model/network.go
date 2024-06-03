package model

type Network struct {
	IbyteRate   float64 `mapstructure:"ibyte_rate" json:"ibyte_rate"`
	Ibytes      int64   `mapstructure:"ibytes" json:"ibytes"`
	IpacketRate float64 `mapstructure:"ipacket_rate" json:"ipacket_rate"`
	Ipackets    int64   `mapstructure:"ipackets" json:"ipackets"`
	ObyteRate   float64 `mapstructure:"obyte_rate" json:"obyte_rate"`
	Obytes      int64   `mapstructure:"obytes" json:"obytes"`
	OpacketRate float64 `mapstructure:"opacket_rate" json:"opacket_rate"`
	Opackets    int64   `mapstructure:"opackets" json:"opackets"`
}

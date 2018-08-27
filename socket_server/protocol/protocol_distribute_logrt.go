package protocol

import (
	"github.com/giskook/ring2/pb/common"
)

type DistributeLogRtPkg struct {
	Imei      string
	Time      string
	Result    string
	RandomNum string
	Settings  map[string]string
}

func (d *DistributeLogRtPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_REPORT_LOGIN, d.Imei)
	cmd += d.RandomNum
	cmd += PROTOCOL_SEP
	if d.Result == "1" {
		cmd += "OK"
	} else {
		cmd += "ER"
	}
	cmd += PROTOCOL_SEP
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeLogRt(d *Carrier.Distribute, random_num string) (string, *DistributeLogRtPkg) {
	dlogrt := &DistributeLogRtPkg{
		Imei:     d.Logrt.Imei,
		Time:     d.Logrt.Time,
		Result:   d.Logrt.Result,
		Settings: make(map[string]string),
	}
	for k, v := range d.Logrt.Settings {
		dlogrt.Settings[k] = v
	}

	return d.Logrt.Imei, dlogrt
}

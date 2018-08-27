package protocol

import (
	"github.com/giskook/ring2/pb/common"
	"strconv"
)

type DistributeFresetPkg struct {
	Imei   string
	Serial string
	Freset string // unit: minute
}

func (d *DistributeFresetPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_FRESET, d.Imei)
	cmd += PROTOCOL_SEP
	fre, _ := strconv.Atoi(d.Freset)
	cmd += strconv.FormatInt(int64(fre)*60, 10)
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeFreset(d *Carrier.Distribute) (string, *DistributeFresetPkg) {
	return d.Freset.Imei, &DistributeFresetPkg{
		Imei:   d.Freset.Imei,
		Serial: d.Freset.Serial,
		Freset: d.Freset.Frequency,
	}
}

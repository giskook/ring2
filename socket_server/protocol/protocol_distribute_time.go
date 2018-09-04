package protocol

import (
	"time"
)

type DistributeTimePkg struct {
	Imei string
}

func (d *DistributeTimePkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_RESP_TIME, d.Imei)
	cmd += "TIME"
	cmd += PROTOCOL_SEP
	cmd += time.Now().Format("060102150405")

	return write_tail(cmd)
}

func ParseTime(imei string) (string, *DistributeTimePkg) {
	return imei, &DistributeTimePkg{
		Imei: imei,
	}
}

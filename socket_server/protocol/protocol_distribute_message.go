package protocol

import (
	"encoding/base64"
	iconv "github.com/djimenez/iconv-go"
	"github.com/giskook/ring2/pb/common"
)

type DistributeMessagePkg struct {
	Imei    string
	Serial  string
	Message string
}

func (d *DistributeMessagePkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_MESSAGE, d.Imei)
	msg, _ := iconv.ConvertString(d.Message, "UTF-8", "UTF-16LE")
	cmd += base64.StdEncoding.EncodeToString([]byte(msg))
	cmd += PROTOCOL_SEP
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func ParseDistributeMessage(d *Carrier.Distribute) (string, *DistributeMessagePkg) {
	return d.Message.Imei, &DistributeMessagePkg{
		Imei:    d.Message.Imei,
		Serial:  d.Message.Serial,
		Message: d.Message.Message,
	}
}

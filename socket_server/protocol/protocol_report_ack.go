package protocol

import (
	"github.com/giskook/ring2/base"
	"github.com/giskook/ring2/pb/common"
	"github.com/golang/protobuf/proto"
	"strings"
)

type ReportAckPkg struct {
	Header  *base.Header
	Imei    string
	Serial  string
	CmdType string
	Result  string
}

func (p *ReportAckPkg) Serialize() []byte {
	report := &Carrier.Report{
		Header: &Carrier.Header{
			Appid: p.Header.AppID,
			From:  p.Header.From,
			To:    p.Header.To,
		},
		Protocol: Carrier.Report_ACK,
		Ack: &Carrier.ReportAck{
			Imei:    p.Imei,
			Serial:  p.Serial,
			Cmdtype: p.CmdType,
			Result:  p.Result,
		},
	}

	data, _ := proto.Marshal(report)

	return data
}

func ParseReportAck(p []string, h *base.Header) *ReportAckPkg {
	result := "0"
	if strings.Contains(p[3], "ER") {
		result = "1"
	}

	return &ReportAckPkg{
		Header:  h,
		Imei:    p[2],
		Serial:  "0",
		CmdType: "PFRESET",
		Result:  result,
	}
}

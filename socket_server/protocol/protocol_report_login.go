package protocol

import (
	"github.com/giskook/ring2/base"
	"github.com/giskook/ring2/pb/common"
	"github.com/golang/protobuf/proto"
	"time"
)

type ReportLoginPkg struct {
	Header     *base.Header
	RandomNum  string
	Imei       string
	Imsi       string
	PhoneNum   string
	I          string
	J          string
	Auth       string
	Time       string
	DeviceType string
	Protocol   string
}

func (p *ReportLoginPkg) Serialize() []byte {
	report := &Carrier.Report{
		Header: &Carrier.Header{
			Appid: p.Header.AppID,
			From:  p.Header.From,
			To:    p.Header.To,
		},
		Protocol: Carrier.Report_LOGIN,
		Login: &Carrier.ReportLogin{
			Imei:       p.Imei,
			Imsi:       p.Imsi,
			DeviceType: p.DeviceType,
			Protocol:   p.Protocol,
			Time:       p.Time,
		},
	}

	data, _ := proto.Marshal(report)

	return data
}

func ParseReportLogin(p []string, h *base.Header) *ReportLoginPkg {
	return &ReportLoginPkg{
		Header:     h,
		Imei:       p[2],
		RandomNum:  p[3],
		Imsi:       p[4],
		DeviceType: "DK-WB100",
		Protocol:   "105",
		PhoneNum:   p[5],
		I:          p[6],
		J:          p[7],
		Auth:       p[8],
		Time:       time.Now().Format("060102-150405"),
	}
}

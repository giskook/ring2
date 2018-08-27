package protocol

import (
	"github.com/giskook/ring2/base"
	"github.com/giskook/ring2/pb/common"
	"github.com/giskook/ring2/pb/lbs_parser"
	"github.com/golang/protobuf/proto"
	"strings"
	"time"
)

type ReportLbsPkg struct {
	Header        *base.Header
	Imei          string
	Serial        string
	TerminalStaus string // "00" 正常 ”x1" 拆卸 "1x" 低电
	Batt          string
	Time          string // 格式为YYMMDD-HHMMSS
	PosReason     string // 0 终端主动上报 1 终端被动上报 2 终端位置补报 3 超圈检测
	PosType       string // 0 GPS 1 cell 2 wifi
	GpsInfo       *Gps
	WifiInfo      []*Wifi
	CellInfo      []*Cell
}

func (p *ReportLbsPkg) SerializeExtra() []byte {
	extra := &Carrier.LocationExtra{
		Imei:           p.Imei,
		Serial:         p.Serial,
		TerminalStatus: p.TerminalStaus,
		Time:           p.Time,
		Batt:           p.Batt,
		PosReason:      p.PosReason,
		PosType:        p.PosType,
	}

	data, _ := proto.Marshal(extra)

	return data
}

func (p *ReportLbsPkg) SerializeLbs() []byte {
	report := &Lbs.Report{
		Header: &Lbs.Header{
			Appid: p.Header.AppID,
			From:  p.Header.From,
			To:    p.Header.To,
		},
		Extra: p.SerializeExtra(),
	}

	if p.PosType == LOCATION_TYPE_WIFI {
		report.Type = Lbs.Report_WIFI
		for _, v := range p.WifiInfo {
			report.WifiCell = append(report.WifiCell, &Lbs.WifiCell{
				Mac:     v.Mac,
				Singnal: v.Singnal,
			})
		}
	} else if p.PosType == LOCATION_TYPE_CELL {
		report.Type = Lbs.Report_STATION
		for _, v := range p.CellInfo {
			report.StationCell = append(report.StationCell, &Lbs.StationCell{
				Lac: v.Lac,
				Cid: v.Cid,
				Dbm: v.Dbm,
			})
		}
	}

	data, _ := proto.Marshal(report)

	return data
}

func (p *ReportLbsPkg) Serialize() []byte {
	report := &Carrier.Report{
		Header: &Carrier.Header{
			Appid: p.Header.AppID,
			From:  p.Header.From,
			To:    p.Header.To,
		},
		Protocol: Carrier.Report_LOCATION,
		Location: &Carrier.ReportLocation{
			Extra:     p.SerializeExtra(),
			Longitude: p.GpsInfo.Longitude,
			Latitude:  p.GpsInfo.Latitude,
			Speed:     p.GpsInfo.Speed,
		},
	}

	data, _ := proto.Marshal(report)

	return data
}

func __terminal_status() string {
	return "00"
}

func __time() string {
	return time.Now().Format("060102-150405")
}

func __pos_type() string {
	return "2"
}

func ParseReportLbs(p []string, h *base.Header, from4gps string, from4lbs string) *ReportLbsPkg {
	r := &ReportLbsPkg{
		Header:        h,
		Imei:          p[2],
		Serial:        "0",
		TerminalStaus: __terminal_status(),
		Batt:          "80",
		Time:          __time(),
		PosReason:     "0",
		PosType:       __pos_type(),
	}
	r.Header.From = from4lbs
	values := strings.Split(p[4], "/")
	if len(values)%2 != 0 {
		return nil
	}
	for i := 0; i < len(values); i += 2 {
		mac := values[0+i]
		singnal := values[1+i]
		r.WifiInfo = append(r.WifiInfo, &Wifi{
			Mac:     add_colon(mac),
			Singnal: singnal,
		})
	}

	return r
}

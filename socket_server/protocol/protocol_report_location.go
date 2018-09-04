package protocol

import (
	"github.com/giskook/ring2/base"
	"github.com/giskook/ring2/pb/common"
	"github.com/giskook/ring2/pb/lbs_parser"
	"github.com/golang/protobuf/proto"
	"strconv"
	"strings"
)

const (
	LOCATION_TYPE_GPS  string = "0"
	LOCATION_TYPE_CELL string = "1"
	LOCATION_TYPE_WIFI string = "2"

	LOCATION_POS_REASON_REPORT_OLD_DATA string = "2"

	LOCATION_GPS_SEP       string = ","
	LOCATION_CELL_MAIN_SEP string = "^"
	LOCATION_CELL_SUB_SEP  string = ","
	LOCATION_WIFI_MAIN_SEP string = ";"
	LOCATION_WIFI_SUB_SEP  string = ","
)

type Gps struct {
	Longitude string
	Latitude  string
	Speed     string
}

type Wifi struct {
	Mac     string
	Singnal string
}

type Cell struct {
	Lac string
	Cid string
	Dbm string
}

func add_colon(mac string) string {
	return mac[0:2] + ":" + mac[2:4] + ":" + mac[4:6] + ":" + mac[6:8] + ":" + mac[8:10] + ":" + mac[10:12]
}

type ReportLocationPkg struct {
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

func (p *ReportLocationPkg) SerializeExtra() []byte {
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

func (p *ReportLocationPkg) SerializeLbs() []byte {
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

func (p *ReportLocationPkg) Serialize() []byte {
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

func _terminal_status(ts string) string {
	terminal_status := 0
	if strings.Contains(ts, "A11") {
		terminal_status |= 0x01
	}
	if strings.Contains(ts, "A17") {
		terminal_status |= 0x10
	}
	ret := "00"
	switch terminal_status {
	case 0x01:
		ret = "01"
	case 0x11:
		ret = "11"
	case 0x10:
		ret = "10"
	}

	return ret
}

func _time(date, time string) string {
	return date[4:6] + date[2:4] + date[0:2] + "-" + time
}

func _pos_type(pt, rmc, wifi string) string {
	if rmc != "" && (pt == "A" || pt == "S") {
		if string(rmc[0]) == "A" {
			return "0"
		} else {
			if pt == "A" {
				return "2"
			}
		}
	}

	if wifi != "" && (pt == "A" || pt == "W") {
		return "2"
	}

	return ""
}

func _parse_long_lat(rmc string) (string, string) {
	values := strings.Split(rmc, "/") // A V already judged in _pos_type func
	lat := values[2]                  // ddmm.xxxx
	long := values[4]                 // dddmm.xxxx
	latd, _ := strconv.Atoi(lat[0:2])
	latm, _ := strconv.ParseFloat(lat[2:], 64)
	latm2d := latm / 60
	latdm := float64(latd) + latm2d

	longd, _ := strconv.Atoi(long[0:3])
	longm, _ := strconv.ParseFloat(long[3:], 64)
	longm2d := longm / 60
	longdm := float64(longd) + longm2d

	return strconv.FormatFloat(latdm, 'f', -1, 64), strconv.FormatFloat(longdm, 'f', -1, 64)
}

func ParseReportLocation(p []string, h *base.Header, from4gps string, from4lbs string) *ReportLocationPkg {
	pos_type := _pos_type(p[3], p[6], p[9])
	if pos_type == "" {
		return nil
	}
	r := &ReportLocationPkg{
		Header:        h,
		Imei:          p[2],
		Serial:        "0",
		TerminalStaus: _terminal_status(p[10]),
		Batt:          "80",
		Time:          _time(p[4], p[5]),
		PosReason:     "0",
		PosType:       pos_type,
	}
	r.Header.From = from4lbs
	if r.PosType == LOCATION_TYPE_GPS {
		lat, long := _parse_long_lat(p[6])
		r.Header.From = from4gps
		r.GpsInfo = &Gps{
			Longitude: long,
			Latitude:  lat,
			Speed:     "0",
		}
	} else if r.PosType == LOCATION_TYPE_WIFI {
		values := strings.Split(p[9], "/")
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
	}

	return r
}

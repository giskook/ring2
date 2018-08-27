package protocol

import (
	"github.com/giskook/ring2/base"
	"github.com/giskook/ring2/pb/common"
	"github.com/giskook/ring2/pb/lbs_parser"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

type DistributeLocationParsedPkg struct {
	Header *base.Header

	Imei         string
	Time         string
	PosReason    string
	ParsedResult string

	Longitude string
	Latitude  string
	Extra     []byte
}

func (d *DistributeLocationParsedPkg) Serialize() []byte {
	cmd := write_header(PROTOCOL_DISTRIBUTE_RESP_LBS, d.Imei)
	t, err := time.Parse("060102-150405", d.Time)
	if err != nil {
		log.Println("parse time err")
	}

	location := t.Format("060102")
	location += "/"
	location += t.Format("150405")
	location += "/"
	location += base.Latd2dm(d.Latitude)
	location += "S"
	location += base.Longd2dm(d.Longitude)
	location += "E"
	cmd += PROTOCOL_SEP
	cmd += PROTOCOL_END_FLAG

	return []byte(cmd)
}

func (p *DistributeLocationParsedPkg) SerializeToUpper() []byte {
	report := &Carrier.Report{
		Header: &Carrier.Header{
			Appid: p.Header.AppID,
			From:  p.Header.From,
			To:    p.Header.To,
		},
		Protocol: Carrier.Report_LOCATION,
		Location: &Carrier.ReportLocation{
			Extra:     p.Extra,
			Longitude: p.Longitude,
			Latitude:  p.Latitude,
		},
	}
	data, _ := proto.Marshal(report)

	return data
}

func ParsedDistributeLocationParsed(d *Lbs.Distribute, header *base.Header) (string, *DistributeLocationParsedPkg) {
	extra := &Carrier.LocationExtra{}
	proto.Unmarshal(d.Extra, extra)
	return extra.Imei, &DistributeLocationParsedPkg{
		Header:       header,
		Imei:         extra.Imei,
		Time:         extra.Time,
		PosReason:    extra.PosReason,
		ParsedResult: d.ParseResult,
		Longitude:    d.Longitude,
		Latitude:     d.Latitude,
		Extra:        d.Extra,
	}
}

package socket_server

import (
	"github.com/giskook/ring2/base"
	"github.com/giskook/ring2/socket_server/protocol"
	"log"
	"strconv"
)

func (ss *SocketServer) eh_report_login(p []string, c *Connection) {
	header := &base.Header{
		AppID: ss.conf.AppID,
		From:  ss.conf.UUID,
		To:    ss.conf.Nsq.TopicPManage,
	}
	login := protocol.ParseReportLogin(p, header)
	id, _ := strconv.ParseUint(login.Imei, 10, 64)
	c.imei = login.Imei
	c.ID = id
	c.random_num = login.RandomNum
	ss.cm.Put(id, c)
	ss.SocketIn <- &base.SocketData{
		Header: header,
		Data:   login.Serialize(),
	}
}

func (ss *SocketServer) eh_report_location(p []string) {
	// because of the 'TPOSUP' protocol trigger two different way
	// so if 'TPOSUP' is GPS go to the upper server gpsly
	// if 'TPOSUP' is wifi or cell infos the protocol goto location parser server
	// but in the protobuf the header's finnal dst will always be upper server
	header := &base.Header{
		AppID: ss.conf.AppID,
		From:  ss.conf.UUID,
		To:    ss.conf.Nsq.TopicPLocation,
	}

	location := protocol.ParseReportLocation(p, header, ss.conf.UUID, ss.conf.Nsq.TopicLbsConsumer.Topic)
	if location == nil {
		log.Println("<INF> location error")
		return
	}
	report_loaction_gps := func() {
		ss.SocketIn <- &base.SocketData{
			Header: header,
			Data:   location.Serialize(),
		}
	}
	report_location_lbs := func() {
		header.From = ss.conf.Nsq.TopicLbsConsumer.Topic
		header.To = ss.conf.Nsq.TopicPLocationParser

		ss.SocketIn <- &base.SocketData{
			Header: header,
			Data:   location.SerializeLbs(),
		}
	}
	if location.PosReason == protocol.LOCATION_POS_REASON_REPORT_OLD_DATA {
		if location.PosType != protocol.LOCATION_TYPE_GPS {
			report_location_lbs()
		} else {
			report_loaction_gps()
		}
	} else {
		if location.PosType != protocol.LOCATION_TYPE_GPS {
			report_location_lbs()
		} else {
			report_loaction_gps()
		}
	}

}

func (ss *SocketServer) eh_report_time(p []string) {
	ss.Send(protocol.ParseTime(p[2]))
}

func (ss *SocketServer) eh_report_lbs(p []string) {
	header := &base.Header{
		AppID: ss.conf.AppID,
		From:  ss.conf.UUID,
		To:    ss.conf.Nsq.TopicPLocation,
	}

	location := protocol.ParseReportLbs(p, header, ss.conf.UUID, ss.conf.Nsq.TopicLbsConsumer.Topic)
	if location == nil {
		log.Println("<INF> lbs error")
		return
	}
	report_location_lbs := func() {
		header.From = ss.conf.Nsq.TopicLbsConsumer.Topic
		header.To = ss.conf.Nsq.TopicPLocationParser

		ss.SocketIn <- &base.SocketData{
			Header: header,
			Data:   location.SerializeLbs(),
		}
	}
	report_location_lbs()
}

func (ss *SocketServer) eh_report_ack(p []string) {
	header := &base.Header{
		AppID: ss.conf.AppID,
		From:  ss.conf.UUID,
		To:    ss.conf.Nsq.TopicPControl,
	}
	ack := protocol.ParseReportAck(p, header)

	ss.SocketIn <- &base.SocketData{
		Header: header,
		Data:   ack.Serialize(),
	}
}

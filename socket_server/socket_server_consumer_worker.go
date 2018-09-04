package socket_server

import (
	"github.com/giskook/ring2/base"
	"github.com/giskook/ring2/pb/common"
	"github.com/giskook/ring2/pb/lbs_parser"
	"github.com/giskook/ring2/socket_server/protocol"
	"github.com/golang/protobuf/proto"
	"log"
	//"strconv"
)

func (ss *SocketServer) consumer_lbs(l []byte) {
	lbs := &Lbs.Distribute{}
	err := proto.Unmarshal(l, lbs)
	log.Printf("<IN NSQ LBS> %s \n", l)
	if err != nil {
		log.Printf("<ERR> %s unmarshal error\n", l)
	} else {
		header := &base.Header{
			AppID: ss.conf.AppID,
			From:  ss.conf.UUID,
			To:    ss.conf.Nsq.TopicPLocation,
		}
		imei, d := protocol.ParsedDistributeLocationParsed(lbs, header)
		// no need to send to upper any more
		//		if d.ParsedResult == "0" {
		//			ss.SocketIn <- &base.SocketData{
		//				Header: header,
		//				Data:   d.SerializeToUpper(),
		//			}
		//		}
		ss.Send(imei, d)
	}
}

func (ss *SocketServer) consumer_worker() {
	ss.wait_exit.Add(1)
	go func() {
		for {
			select {
			case <-ss.exit:
				ss.wait_exit.Done()
				return
			case l := <-ss.SocketLbsOut:
				ss.consumer_lbs(l)
			case p := <-ss.SocketOut:
				distribute := &Carrier.Distribute{}
				err := proto.Unmarshal(p, distribute)
				log.Printf("<IN NSQ> %s \n", p)
				if err != nil {
					log.Println("<ERR> %s unmarshal error\n", p)
				} else {
					var err error
					switch distribute.Protocol {
					case Carrier.Distribute_LOGRT:
						imei, p, _ := protocol.ParseDistributeLogRt(distribute, "")
						//interal, _ := strconv.Atoi(i)

						c := ss.SetStatus(imei, p.Result)
						p.RandomNum = c.random_num
						//go c.reqp(3)
						err := c.Send(p)
						c.Send(&protocol.DistributeFresetPkg{
							Imei:   c.imei,
							Freset: "5",
						})
						//c.Send(&protocol.DistributeReqpPkg{
						//						Imei: c.imei,
						//				})

						if err != nil {
							log.Printf("<ERR> %s %s\n", imei, err.Error())
						}
						//err = ss.Send(protocol.ParseDistributeLogRt(distribute))
					case Carrier.Distribute_REQP:
						err = ss.Send(protocol.ParseDistributeReqp(distribute))
					case Carrier.Distribute_MESSAGE:

						imei, d := protocol.ParseDistributeMessage(distribute)
						err = ss.Send(imei, d)
						header := &base.Header{
							AppID: ss.conf.AppID,
							From:  ss.conf.UUID,
							To:    ss.conf.Nsq.TopicPControl,
						}
						receipt := protocol.ParseReportReceipt(d.Imei, d.Serial, header)
						ss.SocketIn <- &base.SocketData{
							Header: header,
							Data:   receipt.Serialize(),
						}
					case Carrier.Distribute_FRESET:
						err = ss.Send(protocol.ParseDistributeFreset(distribute))
					}

					if err != nil {
						log.Printf("%x %s\n", p, err.Error())
					}
				}

			}
		}
	}()
}

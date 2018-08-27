package main

import (
	"github.com/bitly/go-nsq"
	"github.com/giskook/ring2/pb"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

type nsq_server struct {
	c *nsq.Consumer
	p *nsq.Producer
}

func (n *nsq_server) feedback_login(r *Carrier.Report) {
	from := r.Header.From
	r.Header.From = r.Header.To
	r.Header.To = from
	fr := &Carrier.Distribute{
		Header:   r.Header,
		Protocol: Carrier.Distribute_LOGRT,
		Logrt: &Carrier.DistributeLogRt{
			Imei:   r.Login.Imei,
			Time:   time.Now().Format("060102-150405"),
			Result: "1",
		},
	}
	data, _ := proto.Marshal(fr)

	n.p.PublishAsync(from, data, nil)
}

func (n *nsq_server) HandleMessage(message *nsq.Message) error {
	data := message.Body
	r := &Carrier.Report{}
	err := proto.Unmarshal(data, r)
	if err != nil {
		log.Println("unmarshal error")
	} else {
		switch r.Protocol {
		case Carrier.Report_LOGIN:
			n.feedback_login(r)
		case Carrier.Report_LOCATION:
			log.Println(r.Location.Longitude)
		}
	}

	return nil
}

func new_nsq() *nsq_server {
	return &nsq_server{}
}

func (n *nsq_server) start() {
	nsq_conf_p := nsq.NewConfig()
	p, e := nsq.NewProducer("192.168.2.67:4150", nsq_conf_p)
	if e != nil {
		log.Println(e.Error())
		return
	}
	n.p = p

	nsq_conf := nsq.NewConfig()

	c, err := nsq.NewConsumer("ring_dcs", "dcs_chan", nsq_conf)

	if err != nil {
		return
	}
	n.c = c

	c.AddHandler(n)
	err = c.ConnectToNSQD("192.168.2.67:4150")
	if err != nil {
		log.Println(err.Error())
	}

}

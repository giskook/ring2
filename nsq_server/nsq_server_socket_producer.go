package nsq_server

import (
	"github.com/bitly/go-nsq"
	"github.com/giskook/ring2/base"
	"github.com/giskook/ring2/conf"
	"log"
)

type nsq_server_socket_p struct {
	conf *conf.Nsq
	p    *nsq.Producer
}

func new_nsq_server_socket_p(conf *conf.Nsq) *nsq_server_socket_p {
	nsq_conf := nsq.NewConfig()

	p, err := nsq.NewProducer(conf.Addr, nsq_conf)
	if err != nil {
		return nil
	}

	return &nsq_server_socket_p{
		conf: conf,
		p:    p,
	}
}

func (n *nsq_server_socket_p) stop() {
	n.p.Stop()
}

func (n *nsq_server_socket_p) send(topic string, data []byte) {
	log.Printf("[NSQ OUT] topic %s value %s\n", topic, string(data))
	n.p.PublishAsync(topic, data, nil)
}

func (n *NsqServer) create_producer_socket() error {
	for i := 0; i < n.conf.ProducerCount; i++ {
		n.producers = append(n.producers, new_nsq_server_socket_p(n.conf))
	}

	for _, v := range n.producers {
		if v == nil {
			return base.ErrNsqProducerSocketCreateFail
		}
	}

	return nil
}

func (n *NsqServer) producer_worker(c chan *base.SocketData, producer_inx int) {
	n.wait_exit.Add(1)
	go func() {
		p := n.producers[producer_inx]
		for {
			select {
			case <-n.exit:
				n.wait_exit.Done()
				return
			case socket_data := <-c:
				p.send(socket_data.Header.To, socket_data.Data)
			}
		}
	}()
}

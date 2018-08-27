package nsq_server

import (
	"github.com/bitly/go-nsq"
	"github.com/giskook/ring2/base"
	"github.com/giskook/ring2/conf"
)

type nsq_server_socket_c struct {
	conf    *conf.NsqConsumerTopic
	c       *nsq.Consumer
	handler nsq.HandlerFunc
}

func new_nsq_server_socket_c(conf *conf.NsqConsumerTopic, handler nsq.HandlerFunc) *nsq_server_socket_c {
	nsq_conf := nsq.NewConfig()

	c, err := nsq.NewConsumer(conf.Topic, conf.Channel, nsq_conf)
	if err != nil {
		return nil
	}

	return &nsq_server_socket_c{
		conf:    conf,
		c:       c,
		handler: handler,
	}
}

func (n *nsq_server_socket_c) connect_to_server(addr string) error {
	n.c.AddHandler(n.handler)
	err := n.c.ConnectToNSQD(addr)
	if err != nil {
		return err
	}

	return nil
}

func (n *nsq_server_socket_c) stop() {
	n.c.Stop()
}

func (n *NsqServer) create_consumer_socket(addr string) error {
	for i := 0; i < n.conf.TopicConsumer.Number; i++ {
		c := new_nsq_server_socket_c(n.conf.TopicConsumer, n.HandleMessage)

		n._consumers = append(n._consumers, c)

	}

	for i := 0; i < n.conf.TopicLbsConsumer.Number; i++ {
		c := new_nsq_server_socket_c(n.conf.TopicLbsConsumer, n.HandleLbsMessage)

		n._consumers = append(n._consumers, c)
	}
	for _, c := range n._consumers {
		if c != nil {
			err := c.connect_to_server(addr)
			if err != nil {
				return err
			}
		} else {
			return base.ErrNsqConsumerSocketCreateFail
		}
	}

	return nil
}

package nsq_server

import (
	"github.com/bitly/go-nsq"
)

func (n *NsqServer) HandleMessage(message *nsq.Message) error {
	n.NsqDataIn <- message.Body
	return nil
}

func (n *NsqServer) HandleLbsMessage(message *nsq.Message) error {
	n.NsqLbsIn <- message.Body
	return nil
}

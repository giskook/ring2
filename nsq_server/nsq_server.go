package nsq_server

import (
	"github.com/giskook/ring2/base"
	"github.com/giskook/ring2/conf"
	"sync"
)

type NsqServer struct {
	conf       *conf.Nsq
	_consumers []*nsq_server_socket_c
	producers  []*nsq_server_socket_p
	NsqDataIn  chan []byte
	NsqLbsIn   chan []byte
	NsqDataOut chan *base.SocketData
	exit       chan struct{}
	wait_exit  *sync.WaitGroup
}

func NewNsqServer(conf *conf.Nsq) *NsqServer {
	return &NsqServer{
		conf:       conf,
		NsqDataIn:  make(chan []byte),
		NsqLbsIn:   make(chan []byte),
		NsqDataOut: make(chan *base.SocketData),
		exit:       make(chan struct{}),
		_consumers: make([]*nsq_server_socket_c, 0),
		producers:  make([]*nsq_server_socket_p, 0),
		wait_exit:  new(sync.WaitGroup),
	}
}

func (n *NsqServer) Start() error {
	err := n.create_consumer_socket(n.conf.Addr)
	if err != nil {
		return err
	}
	err = n.create_producer_socket()
	if err != nil {
		return err
	}

	for i := 0; i < n.conf.ProducerCount; i++ {
		n.producer_worker(n.NsqDataOut, i)
	}

	return nil
}

func (n *NsqServer) Stop() {
	close(n.exit)
	n.wait_exit.Wait()
	close(n.NsqDataIn)
	close(n.NsqLbsIn)
	close(n.NsqDataOut)
}

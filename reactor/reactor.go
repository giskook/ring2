package reactor

import (
	"github.com/giskook/ring2/base"
	"github.com/giskook/ring2/conf"
	"github.com/giskook/ring2/nsq_server"
	"github.com/giskook/ring2/socket_server"
	"log"
)

type Reactor struct {
	conf *conf.Conf
	exit chan struct{}

	socket_server *socket_server.SocketServer
	nsq_server    *nsq_server.NsqServer
}

func NewReactor(conf *conf.Conf) *Reactor {
	return &Reactor{
		conf:          conf,
		exit:          make(chan struct{}),
		socket_server: socket_server.NewSocketServer(conf),
		nsq_server:    nsq_server.NewNsqServer(conf.Nsq),
	}
}

func (r *Reactor) Start() error {
	err := r.socket_server.Start()
	if err != nil {
		return err
	}
	err = r.nsq_server.Start()
	if err != nil {
		return err
	}

	r.shunt()

	return nil
}

func (r *Reactor) Stop() {
	r.socket_server.Stop()
	log.Printf("<INFO> %s\n", base.SOCKET_SERVER_STOPPED)
	r.nsq_server.Stop()
	log.Printf("<INFO> %s\n", base.NSQ_SERVER_STOPPED)
	close(r.exit)
}

package reactor

import ()

func (rc *Reactor) shunt() {
	defer func() {
	}()

	for {
		select {
		case <-rc.exit:
			return
		case socket_in := <-rc.socket_server.SocketIn:
			rc.nsq_server.NsqDataOut <- socket_in
		case nsq_data_in := <-rc.nsq_server.NsqDataIn:
			rc.socket_server.SocketOut <- nsq_data_in
		case nsq_lbs_in := <-rc.nsq_server.NsqLbsIn:
			rc.socket_server.SocketLbsOut <- nsq_lbs_in
		}

	}
}

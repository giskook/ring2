package conf

import (
	"encoding/json"
	"os"
)

type NsqConsumerTopic struct {
	Topic   string
	Channel string
	Number  int
}

type Nsq struct {
	Addr                 string
	ProducerCount        int
	TopicPManage         string
	TopicPControl        string
	TopicPLocation       string
	TopicPLocationParser string
	TopicConsumer        *NsqConsumerTopic
	TopicLbsConsumer     *NsqConsumerTopic
}

type TcpServer struct {
	BindPort          string
	ReadLimit         int
	WriteLimit        int
	ConnTimeout       int
	ConnCheckInterval int
	WorkerNum         int
}

type Conf struct {
	AppID     string
	UUID      string
	Nsq       *Nsq
	TcpServer *TcpServer
}

func ReadConfig(confpath string) (*Conf, error) {
	file, _ := os.Open(confpath)
	decoder := json.NewDecoder(file)
	config := Conf{}
	err := decoder.Decode(&config)

	return &config, err
}

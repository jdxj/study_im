package client

import (
	"log"

	"github.com/name5566/leaf/network"
)

type agent struct {
	conn        *network.TCPConn
	client      *Client
	sendChan    chan interface{}
	receiveChan chan interface{}
}

func (a *agent) Run() {
	go a.ReadMsg()

	for {
		msg := <-a.sendChan
		if msg == nil {
			log.Println("agent stop to read")
			return
		}

		data, err := a.client.Processor.Marshal(msg)
		if err != nil {
			log.Printf("Marshal: %s\n", err)
			continue
		}

		err = a.conn.WriteMsg(data...)
		if err != nil {
			log.Printf("WriteMsg: %s\n", err)
		}
	}
}

func (a *agent) ReadMsg() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Printf("ReadMsg: %s\n", err)
			return
		}

		msg, err := a.client.Processor.Unmarshal(data)
		if err != nil {
			log.Printf("Unmarshal: %s\n", err)
			continue
		}

		a.receiveChan <- msg
	}
}

func (a *agent) OnClose() {

}

package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/panjf2000/gnet"

	"github.com/jdxj/study_im/codec/frame"
	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/logger"
)

func NewGate(host string, port int) *Gate {
	gate := &Gate{
		host: host,
		port: port,
	}
	return gate
}

type Gate struct {
	*gnet.EventServer
	host string
	port int

	// todo: id 生成器
	idGen  int
	idConn *sync.Map
}

func (gate *Gate) Serve() error {
	log.Printf("server started\n")
	addr := fmt.Sprintf("%s:%d", gate.host, gate.port)
	return gnet.Serve(gate, addr,
		gnet.WithMulticore(true),
		gnet.WithCodec(frame.NewLengthFieldBasedFrameCodec()),
	)
}

func (gate *Gate) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	_, msg, err := protobuf.Unmarshal(frame)
	if err != nil {
		logger.Errorf("Unmarshal: %s", err)
		return nil, 0
	}

	agent := &Agent{
		nodeID: 0, // todo: 初始化
		userID: gate.idGen,
		conn:   c,
	}
	out = handle(agent, msg)
	return
}

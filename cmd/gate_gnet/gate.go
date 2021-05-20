package main

import (
	"fmt"
	"sync"

	"github.com/bwmarrin/snowflake"
	"github.com/panjf2000/gnet"

	"github.com/jdxj/study_im/codec/frame"
	"github.com/jdxj/study_im/codec/protobuf"
	"github.com/jdxj/study_im/logger"
)

func NewGate(host string, port, nodeID int) (*Gate, error) {
	gate := &Gate{
		host:   host,
		port:   port,
		nodeID: uint32(nodeID),
	}

	gate.am = &AgentManager{
		mutex:  &sync.RWMutex{},
		agents: make(map[int64]gnet.Conn),
	}
	gate.sm = &SeqManager{}

	var err error
	gate.idGenerator, err = snowflake.NewNode(int64(nodeID))
	return gate, err
}

type Gate struct {
	*gnet.EventServer
	host string
	port int

	nodeID      uint32
	idGenerator *snowflake.Node

	am *AgentManager
	sm *SeqManager // todo: seq 由发送队列管理
}

func (gate *Gate) Serve() error {
	logger.Infof("server started")
	addr := fmt.Sprintf("%s:%d", gate.host, gate.port)
	return gnet.Serve(gate, addr,
		gnet.WithMulticore(true),
		gnet.WithCodec(frame.NewLengthFieldBasedFrameCodec()),
	)
}

func (gate *Gate) React(frame []byte, conn gnet.Conn) (out []byte, action gnet.Action) {
	rawMsg, err := protobuf.Unmarshal(frame)
	if err != nil {
		logger.Errorf("Unmarshal: %s", err)
		return nil, 0
	}

	out = gate.handle(conn, rawMsg)
	return
}

func (gate *Gate) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	agentID := gate.nextID()
	c.SetContext(agentID)
	gate.am.AddAgent(agentID, c)
	return
}

func (gate *Gate) OnClosed(conn gnet.Conn, err error) (action gnet.Action) {
	agentID, ok := conn.Context().(int64)
	if !ok {
		return
	}
	gate.am.DelAgent(agentID)
	logger.Debugf("del agent: %d", agentID)
	return
}

func (gate *Gate) nextID() int64 {
	return gate.idGenerator.Generate().Int64()
}

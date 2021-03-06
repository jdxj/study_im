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

	gate.cm = &ClientManager{
		clients: make(map[uint32]*Client, 100000),
	}
	//gate.gm = &GroupManager{
	//	groups: make(map[uint32]*Group, 1000),
	//}
	gate.rm = &RelationManager{
		connections: make(map[int64]gnet.Conn),
	}

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

	seqMutex sync.Mutex
	seq      uint32

	// todo: 心跳
	cm *ClientManager
	//gm *GroupManager
	rm *RelationManager
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
	connID := gate.nextConnID()
	c.SetContext(connID)
	gate.rm.AddConn(connID, c)
	return
}

func (gate *Gate) OnClosed(conn gnet.Conn, err error) (action gnet.Action) {
	connID := conn.Context().(int64)
	gate.rm.DelConn(connID)
	return
}

func (gate *Gate) nextConnID() int64 {
	return gate.idGenerator.Generate().Int64()
}

func (gate *Gate) nextSeq() uint32 {
	gate.seqMutex.Lock()
	curSeq := gate.seq
	gate.seq++
	gate.seqMutex.Unlock()
	return curSeq
}
